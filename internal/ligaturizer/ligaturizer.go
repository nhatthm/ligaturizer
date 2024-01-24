package ligaturizer

import (
	"context"
	"fmt"
	"strings"

	"github.com/bool64/ctxd"

	"go.nhat.io/ligaturizer/internal/fontforge"
)

// Ligaturizer is a font ligaturizer.
type Ligaturizer interface {
	Ligaturize(ctx context.Context, src *fontforge.Font, dst *fontforge.Font) error
}

// Ligaturize ligaturizes a font.
func Ligaturize(
	ctx context.Context,
	l Ligaturizer,
	ligFont, inputFont *fontforge.Font,
	fontName string,
) error {
	if err := l.Ligaturize(ctx, ligFont, inputFont); err != nil {
		return ctxd.WrapError(ctx, err, "failed to ligaturize")
	}

	updateFontMetadata(inputFont, fontName)
	updateCopyright(ligFont, inputFont)

	return nil
}

func updateFontMetadata(font *fontforge.Font, newFamilyName string) {
	curFamilyName := font.FamilyName()
	suffix := ""

	if parts := strings.Split(curFamilyName, "-"); len(parts) >= 2 {
		suffix = parts[1]
	}

	font.SetFamilyName(newFamilyName)

	if suffix != "" {
		font.SetFullName(fmt.Sprintf("%s %s", newFamilyName, suffix))
		font.SetFontName(fmt.Sprintf("%s-%s", strings.ReplaceAll(newFamilyName, " ", ""), suffix))
	} else {
		font.SetFullName(newFamilyName)
		font.SetFontName(strings.ReplaceAll(newFamilyName, " ", ""))
	}

	font.SetSFNTNames(
		"UniqueID", fmt.Sprintf("%s; Ligaturized", font.FullName()),
		"Preferred Family", newFamilyName,
		"Compatible Full", newFamilyName,
		"Family", newFamilyName,
		"WWS Family", newFamilyName,
	)
}

func updateCopyright(src, dst *fontforge.Font) {
	dst.SetCopyright(formatCopyright(
		src.Copyright(),
		dst.Copyright(),
	))
	dst.SetSFNTNames("Copyright", formatCopyright(
		src.SFNTNames().Find("Copyright"),
		dst.SFNTNames().Find("Copyright"),
	))
}

func formatCopyright(srcCopyright, dstCopyright string) string {
	copyright := fmt.Sprintf("%s\n%s", dstCopyright, srcCopyright)
	copyright = strings.Trim(copyright, "\n")

	return copyright
}
