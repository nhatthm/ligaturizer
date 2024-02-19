package fontforge

import (
	"fmt"
	"sync"

	"github.com/spf13/afero"
	"go.nhat.io/python3"
)

const moduleName = "fontforge"

var getModule = sync.OnceValue(func() *python3.Object {
	module, err := python3.ImportModule(moduleName)
	if err != nil {
		panic(err)
	}

	return module
})

// Open opens a font file.
func Open(path string) (*Font, error) {
	exists, _ := afero.Exists(afero.NewOsFs(), path) //nolint: errcheck
	if !exists {
		return nil, fmt.Errorf("%w: %q", afero.ErrFileNotFound, path)
	}

	f := getModule().CallMethodArgs("open", path)

	if err := python3.LastError(); err != nil {
		return nil, err //nolint: wrapcheck
	}

	return newFont(f), nil
}

// Generate generates a font file.
func Generate(font *Font, path string) error {
	// Work around a bug in Fontforge where the underline height is subtracted from the underline width when calling generate().
	font.SetUnderlinePosition(font.UnderlinePosition() + font.UnderlineWith())

	font.callMethodArgs("generate", path)

	return python3.LastError() //nolint: wrapcheck
}

// CopyGlyph copies a glyph into (FontForge's internal) clipboard.
func CopyGlyph(font *Font, glyph string) error {
	sel := font.obj.GetAttr("selection")
	defer sel.DecRef()

	sel.CallMethodArgs("none")

	if err := python3.LastError(); err != nil {
		return err //nolint: wrapcheck
	}

	sel.CallMethodArgs("select", glyph)

	if err := python3.LastError(); err != nil {
		return err //nolint: wrapcheck
	}

	font.callMethodArgs("copy")

	if err := python3.LastError(); err != nil {
		return err //nolint: wrapcheck
	}

	return nil
}

// PasteGlyph pastes the contents of (FontForge's internal) clipboard into the selected glyphs – and removes what was there before.
func PasteGlyph(font *Font, glyph string) error {
	sel := font.obj.GetAttr("selection")
	defer sel.DecRef()

	sel.CallMethodArgs("none")

	if err := python3.LastError(); err != nil {
		return err //nolint: wrapcheck
	}

	sel.CallMethodArgs("select", glyph)

	if err := python3.LastError(); err != nil {
		return err //nolint: wrapcheck
	}

	font.callMethodArgs("paste")

	if err := python3.LastError(); err != nil {
		return err //nolint: wrapcheck
	}

	return nil
}

// Lookup types.
const (
	LookupTypeGSubSingle       = LookupType("gsub_single")
	LookupTypeGSubContextChain = LookupType("gsub_contextchain")
)

// LookupType is a font lookup type.
type LookupType string

// MarshalPyObject returns the underlying PyObject.
func (t LookupType) MarshalPyObject() *python3.Object {
	return python3.NewString(string(t))
}

// Lookup flags.
const (
	LookupFlagRightToLeft     = LookupFlag("right_to_left")
	LookupFlagIgnoreBases     = LookupFlag("ignore_bases")
	LookupFlagIgnoreLigatures = LookupFlag("ignore_ligatures")
	LookupFlagIgnoreMarks     = LookupFlag("ignore_marks")
)

// LookupFlags are font lookup flags.
type LookupFlags []LookupFlag

// MarshalPyObject returns the underlying Object.
func (f LookupFlags) MarshalPyObject() *python3.Object {
	return python3.NewTupleFromValues(([]LookupFlag)(f)...).AsObject()
}

// LookupFlag is a font lookup flag.
type LookupFlag string

// MarshalPyObject returns the underlying Object.
func (f LookupFlag) MarshalPyObject() *python3.Object {
	return python3.NewString(string(f))
}

// LookupFeatures are font lookup features.
type LookupFeatures []LookupFeature

// MarshalPyObject returns the underlying Object.
func (s LookupFeatures) MarshalPyObject() *python3.Object {
	return python3.NewTupleFromValues(([]LookupFeature)(s)...).AsObject()
}

// LookupFeature is a font lookup feature.
type LookupFeature struct {
	TagName LookupFeatureTag
	Scripts []LookupFeatureScript
}

// MarshalPyObject returns the underlying Object.
func (f LookupFeature) MarshalPyObject() *python3.Object {
	return python3.NewTupleFromAny(
		f.TagName,
		python3.NewTupleFromValues(f.Scripts...),
	).AsObject()
}

// WithScript adds a script to the lookup feature.
func (f LookupFeature) WithScript(scriptTag LookupFeatureScriptTag, scriptLanguages ...LookupFeatureScriptLanguage) LookupFeature {
	f.Scripts = append(f.Scripts, LookupFeatureScript{
		Tag:       scriptTag,
		Languages: scriptLanguages,
	})

	return f
}

// NewLookupFeature creates a new lookup feature.
func NewLookupFeature(tag string) LookupFeature {
	return LookupFeature{
		TagName: LookupFeatureTag(tag),
	}
}

// LookupFeatureTag is a font lookup feature tag.
type LookupFeatureTag string

// MarshalPyObject returns the underlying Object.
func (t LookupFeatureTag) MarshalPyObject() *python3.Object {
	return python3.NewString(fmt.Sprintf("%-4s", t))
}

// LookupFeatureScript is a font lookup feature script.
type LookupFeatureScript struct {
	Tag       LookupFeatureScriptTag
	Languages []LookupFeatureScriptLanguage
}

// MarshalPyObject returns the underlying Object.
func (s LookupFeatureScript) MarshalPyObject() *python3.Object {
	return python3.NewTupleFromAny(
		s.Tag,
		python3.NewTupleFromValues(s.Languages...),
	).AsObject()
}

// LookupFeatureScriptTag is a font lookup feature script tag.
type LookupFeatureScriptTag string

// MarshalPyObject returns the underlying Object.
func (t LookupFeatureScriptTag) MarshalPyObject() *python3.Object {
	return python3.NewString(fmt.Sprintf("%-4s", t))
}

// LookupFeatureScriptLanguage is a font lookup feature script language.
type LookupFeatureScriptLanguage string

// MarshalPyObject returns the underlying Object.
func (l LookupFeatureScriptLanguage) MarshalPyObject() *python3.Object {
	return python3.NewString(fmt.Sprintf("%-4s", l))
}

// AddLookup creates a new lookup with the given name, type and flags. It will tag it with any indicated features.
func AddLookup(font *Font, lookupName string, lookupType LookupType, lookupFlags LookupFlags, lookupFeatures LookupFeatures) {
	font.callMethodArgs("addLookup", lookupName, lookupType, lookupFlags, lookupFeatures)
}

// AddLookupSubtable creates a new subtable within the specified lookup. The lookup name should be a string specifying an existing lookup.
// The subtable name should also be a string and should not match any currently existing subtable in the lookup.
func AddLookupSubtable(font *Font, lookupName string, lookupSubtableName string) {
	font.callMethodArgs("addLookupSubtable", lookupName, lookupSubtableName)
}

// AddContextualLookupSubtable creates a new subtable within the specified contextual lookup (contextual, contextual chaining, or reverse contextual chaining).
// The lookup name should be a string specifying an existing lookup. The subtable name should also be a string and should not match any currently existing
// subtable in the lookup.
func AddContextualLookupSubtable(font *Font, lookupName string, lookupSubtableName string, lookupType, lookupRule string) {
	font.callMethodArgs("addContextualSubtable", lookupName, lookupSubtableName, lookupType, lookupRule)
}

// AddPositionSubstitutionVariant adds position/substitution data to the glyph. The number and type of the arguments vary according to the type of the lookup
// containing the subtable.
func AddPositionSubstitutionVariant(glyph *Glyph, subtableName string, variant string) {
	glyph.callMethodArgs("addPosSub", subtableName, variant)
}
