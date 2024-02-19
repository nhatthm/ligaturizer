package firacode

import (
	"context"
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/bool64/ctxd"
	"go.nhat.io/fontforge"
	"go.nhat.io/psmat"

	"go.nhat.io/ligaturizer/internal/ligaturizer"
)

var (
	// ErrNoLigatures indicates that there is no ligatures in the config.
	ErrNoLigatures = errors.New("no ligatures")
	// ErrLigatureMappingHasNoChar indicates that a ligature mapping has no character.
	ErrLigatureMappingHasNoChar = errors.New("ligature mapping has no char")
	// ErrLigatureUnsupported indicates that a ligature is not supported.
	ErrLigatureUnsupported = errors.New("ligature unsupported")
)

var _ ligaturizer.Ligaturizer = (*Ligaturizer)(nil)

// Ligaturizer is a font ligaturizer.
type Ligaturizer struct {
	logger ctxd.Logger

	cfg Config
}

// correctCharacterWidth corrects the horizontal advance of characters to match the em width of the output font, and depending on the width of the glyph, the
// em width of the output font, and the value of the Config.ScaleCharacterGlyphsThreshold config optionally horizontally scale it.
//
//	Glyphs that are not horizontally scaled, but which still need horizontal advance correction, will be centered instead.
//
// Note: Width-correct copied individual characters (not ligatures!).
func (l *Ligaturizer) correctCharacterWidth(glyph *fontforge.Glyph, emWidth int) {
	if glyph.Width() == emWidth {
		return
	}

	widthDelta := math.Abs(float64(glyph.Width()-emWidth)) / float64(emWidth)
	if widthDelta >= l.cfg.ScaleCharacterGlyphsThreshold {
		// Character is too wide/narrow compared to output font; scale it.
		scale := psmat.Scale(float64(emWidth)/float64(glyph.Width()), 1.0)

		glyph.Transform(scale)
	} else {
		// Do not scale; just center copied characters in their hbox.
		// Fix horizontal advance first, to recalculate the bearings.
		glyph.SetWidth(emWidth)
		// Correct bearings to center the glyph.
		glyph.SetLeftSideBearing((glyph.LeftSideBearing() + glyph.RightSideBearing()) / 2)
		glyph.SetRightSideBearing(glyph.LeftSideBearing())
	}

	// Final adjustment of horizontal advance to correct for rounding errors when scaling/centering -- otherwise small errors can result in visible misalignment
	// near the end of long lines.
	glyph.SetWidth(emWidth)
}

// correctLigatureWidth corrects the horizontal advance and scale of a ligature.
func (l *Ligaturizer) correctLigatureWidth(glyph *fontforge.Glyph, emWidth int) {
	if glyph.Width() == emWidth {
		return
	}

	//nolint: godox
	// TODO: some kind of threshold here, similar to the character glyph scale threshold? The largest ligature uses 0.956 of its hbox, so if the target font is
	// within 4% of the source font size, we don't need to resize -- but we may want to adjust the bearings. And we can't just center it, because ligatures are
	// characterized by very large negative left bearings -- they advance 1em, but draw from (-(n-1))em to +1em.
	scale := float64(emWidth) / float64(glyph.Width())

	glyph.Transform(psmat.Scale(scale, 1.0))
	glyph.SetWidth(emWidth)
}

func (l *Ligaturizer) copyCharacterGlyphs(ctx context.Context, src *fontforge.Font, dst *fontforge.Font, chars []ligaturizer.Char, emWidth int) error {
	if !l.cfg.CopyCharacterGlyphs {
		return nil
	}

	for _, c := range chars {
		ctx := ctxd.AddFields(ctx, "glyph", c.Name)

		if err := fontforge.CopyGlyph(src, c.Name); err != nil {
			return ctxd.WrapError(ctx, err, "failed to copy glyph")
		}

		if err := fontforge.PasteGlyph(dst, c.Name); err != nil {
			return ctxd.WrapError(ctx, err, "failed to paste glyph")
		}

		l.correctCharacterWidth(dst.Glyph(c.Rune), emWidth)
	}

	return nil
}

func (l *Ligaturizer) addContextualLookupSubtable(dst *fontforge.Font, lookupName string, lookupSubtableName string, ruleFormat string, ruleArgs ...string) {
	rules := make([]string, 0, len(ruleArgs))

	for i := 0; i < len(ruleArgs); i += 2 {
		rules = append(rules, fmt.Sprintf("{%s}", ruleArgs[i]), ruleArgs[i+1])
	}

	replacer := strings.NewReplacer(rules...)
	rule := replacer.Replace(ruleFormat)

	fontforge.AddContextualLookupSubtable(dst, lookupName, lookupSubtableName, "glyph", rule)
}

// nolint: cyclop,funlen
func (l *Ligaturizer) addLigature(
	ctx context.Context,
	src *fontforge.Font,
	dst *fontforge.Font,
	pos int,
	chars ligaturizer.Chars,
	ligature ligaturizer.LigatureName,
	emWidth int,
) (bool, error) {
	ctx = ctxd.AddFields(ctx, "ligature", ligature, "position", pos)

	// Test if the ligature exists in the source font.
	if err := fontforge.CopyGlyph(src, ligature.String()); err != nil {
		l.logger.Warn(ctx, "ligature does not exist in source font")

		return false, nil //nolint: nilerr
	}

	dstLigatureName := fmt.Sprintf("lig.%d", pos)

	if err := dst.CreateGlyph(dstLigatureName); err != nil {
		return false, ctxd.WrapError(ctx, err, "failed to create glyph in destination font", "glyph", dstLigatureName)
	}

	if err := fontforge.PasteGlyph(dst, dstLigatureName); err != nil {
		return false, ctxd.WrapError(ctx, err, "failed to paste ligature into destination font", "glyph", dstLigatureName)
	}

	l.correctLigatureWidth(dst.Glyph(dstLigatureName), emWidth)

	if err := fontforge.CopyGlyph(dst, "space"); err != nil {
		return false, ctxd.WrapError(ctx, err, "failed to copy space glyph")
	}

	getLookupName := func(i int) string { return fmt.Sprintf("lookup.%d.%d", pos, i) }
	getLookupSubName := func(i int) string { return fmt.Sprintf("lookup.sub.%d.%d", pos, i) }
	getCharName := func(i int) string { return fmt.Sprintf("CR.%d.%d", pos, i) }

	for i, c := range chars {
		fontforge.AddLookup(dst, getLookupName(i), fontforge.LookupTypeGSubSingle, fontforge.LookupFlags{}, fontforge.LookupFeatures{})
		fontforge.AddLookupSubtable(dst, getLookupName(i), getLookupSubName(i))

		if !dst.HasGlyph(c.Name) {
			// We assume here that this is because char is a single letter
			// (e.g. 'w') rather than a character name, and the font we're
			// editing doesn't have glyphnames for letters.
			dst.Glyph(c.Rune).SetGlyphName(c.Name)
		}

		glyphVariant := dstLigatureName

		if i < len(chars)-1 {
			charName := getCharName(i)
			glyphVariant = charName

			if err := dst.CreateGlyph(charName); err != nil {
				return false, ctxd.WrapError(ctx, err, "failed to create glyph in destination font", "glyph", charName)
			}

			if err := fontforge.PasteGlyph(dst, charName); err != nil {
				return false, ctxd.WrapError(ctx, err, "failed to paste glyph into destination font", "glyph", charName)
			}
		}

		fontforge.AddPositionSubstitutionVariant(dst.Glyph(c.Name), getLookupSubName(i), glyphVariant)
	}

	contextLookupName := fmt.Sprintf("calt.%d", pos)
	getContextLookupSubName := func(i int) string { return fmt.Sprintf("calt.%d.%d", pos, i) }

	fontforge.AddLookup(dst, contextLookupName, fontforge.LookupTypeGSubContextChain, fontforge.LookupFlags{}, fontforge.LookupFeatures{
		fontforge.NewLookupFeature("calt").
			WithScript("DFLT", "dflt").
			WithScript("arab", "dflt").
			WithScript("armn", "dflt").
			WithScript("cyrl", "SRB", "dflt").
			WithScript("geor", "dflt").
			WithScript("grek", "dflt").
			WithScript("lao", "dflt").
			WithScript("latn", "CAT", "ESP", "GAL", "ISM", "KSM", "LSM", "MOL", "NSM", "ROM", "SKS", "SSM", "dflt").
			WithScript("math", "dflt").
			WithScript("thai", "dflt"),
	})

	for i, c := range chars {
		prevs := make([]string, 0, i)
		nexts := chars.Names()[i+1:]

		for j := 0; j < i; j++ {
			prevs = append(prevs, getCharName(j))
		}

		l.addContextualLookupSubtable(dst, contextLookupName, getContextLookupSubName(i),
			"{prev} | {cur} @<{lookup}> | {next}",
			"prev", strings.Join(prevs, " "),
			"cur", c.Name,
			"lookup", getLookupName(i),
			"next", strings.Join(nexts, " "),
		)
	}

	// Add ignore rules.
	l.addContextualLookupSubtable(dst, contextLookupName, getContextLookupSubName(len(chars)),
		"| {first} | {rest} {last}",
		"first", chars[0].Name,
		"rest", strings.Join(chars.Names()[1:], " "),
		"last", chars[len(chars)-1].Name,
	)
	l.addContextualLookupSubtable(dst, contextLookupName, getContextLookupSubName(len(chars)+1),
		"{first} | {first} | {rest}",
		"first", chars[0].Name,
		"rest", strings.Join(chars.Names()[1:], " "),
	)

	return true, nil
}

// Ligaturize ligaturizes the font.
func (l *Ligaturizer) Ligaturize(ctx context.Context, src *fontforge.Font, dst *fontforge.Font) error {
	// Scale FiraCode to correct em height.
	src.SetEm(dst.Em())

	emWidth := dst.Glyph('m').Width()
	pos := 1

	for _, m := range l.cfg.Ligatures {
		if m.Ligature == NoLigatureName {
			if err := l.copyCharacterGlyphs(ctx, src, dst, m.Chars, emWidth); err != nil {
				return fmt.Errorf("failed to copy character glyphs: %w", err)
			}

			continue
		}

		if ok, err := l.addLigature(ctx, src, dst, pos, m.Chars, m.Ligature, emWidth); err != nil {
			return fmt.Errorf("failed to add ligature: %w (%q)", err, m.Ligature.String())
		} else if ok {
			pos++
		}
	}

	return nil
}

// NewLigaturizer initiates a new Ligaturizer.
func NewLigaturizer(cfg Config, logger ctxd.Logger) (*Ligaturizer, error) {
	populateConfig(&cfg)

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	sortLigatureMappings(cfg.Ligatures)

	l := &Ligaturizer{
		logger: logger,

		cfg: cfg,
	}

	return l, nil
}

func populateConfig(cfg *Config) {
	if len(cfg.Ligatures) == 0 {
		ligatures := make([]LigatureMapping, len(defaultLigatures))

		for i, m := range defaultLigatures {
			ligatures[i] = LigatureMapping{
				Chars:    append([]ligaturizer.Char(nil), m.Chars...),
				Ligature: m.Ligature,
			}
		}

		cfg.Ligatures = ligatures
	}

	if cfg.ScaleCharacterGlyphsThreshold <= 0 {
		cfg.ScaleCharacterGlyphsThreshold = defaultScaleCharacterGlyphsThreshold
	}
}

func validateConfig(cfg Config) error {
	// Validate ligatures.
	if len(cfg.Ligatures) == 0 {
		return ErrNoLigatures
	}

	for i, m := range cfg.Ligatures {
		if len(m.Chars) == 0 {
			return fmt.Errorf("%w (#%d)", ErrLigatureMappingHasNoChar, i)
		}

		if !isLigatureSupported(m.Ligature) {
			return fmt.Errorf("%w: %q", ErrLigatureUnsupported, m.Ligature)
		}
	}

	return nil
}

func sortLigatureMappings(mappings []LigatureMapping) {
	originalIndexes := make(map[ligaturizer.LigatureName]int, len(mappings))

	for i, m := range mappings {
		originalIndexes[m.Ligature] = i
	}

	sort.Slice(mappings, func(i, j int) bool {
		lenI := len(mappings[i].Chars)
		lenJ := len(mappings[j].Chars)

		if lenI == lenJ {
			return originalIndexes[mappings[i].Ligature] < originalIndexes[mappings[j].Ligature]
		}

		return lenI < lenJ
	})
}
