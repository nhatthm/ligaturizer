package fontforge

import (
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
	"go.nhat.io/python3"
)

// Font is a wrapper around a Python object.
type Font struct {
	obj *python3.Object
}

// callMethodArgs calls a method with args.
func (f *Font) callMethodArgs(name string, args ...any) *python3.Object { //nolint: unparam
	return f.obj.CallMethodArgs(name, args...)
}

// Close closes the font.
func (f *Font) Close() error {
	f.obj.DecRef()
	f.obj.CallMethodArgs("close")

	return python3.LastError() //nolint: wrapcheck
}

// DecRef decreases the reference count of the object.
func (f *Font) DecRef() {
	f.obj.DecRef()
}

// AsObject returns the underlying Object.
func (f *Font) AsObject() *python3.Object {
	return f.obj
}

// FontName returns the font name of the font.
func (f *Font) FontName() string {
	attr := f.obj.GetAttr("fontname")
	defer attr.DecRef()

	return python3.AsString(attr)
}

// SetFontName sets the font name of the font.
func (f *Font) SetFontName(name string) {
	f.obj.SetAttr("fontname", name)
}

// FullName returns the full name of the font.
func (f *Font) FullName() string {
	attr := f.obj.GetAttr("fullname")
	defer attr.DecRef()

	return python3.AsString(attr)
}

// SetFullName sets the full name of the font.
func (f *Font) SetFullName(name string) {
	f.obj.SetAttr("fullname", name)
}

// FamilyName returns the family name of the font.
func (f *Font) FamilyName() string {
	attr := f.obj.GetAttr("familyname")
	defer attr.DecRef()

	return python3.AsString(attr)
}

// SetFamilyName sets the family name of the font.
func (f *Font) SetFamilyName(name string) {
	f.obj.SetAttr("familyname", name)
}

// Em returns the em size of the font.
func (f *Font) Em() int {
	attr := f.obj.GetAttr("em")
	defer attr.DecRef()

	return python3.AsInt(attr)
}

// SetEm sets the em size of the font.
func (f *Font) SetEm(em int) {
	f.obj.SetAttr("em", em)
}

// UnderlinePosition returns the underline position of the font.
func (f *Font) UnderlinePosition() float64 {
	attr := f.obj.GetAttr("upos")
	defer attr.DecRef()

	return python3.AsFloat64(attr)
}

// SetUnderlinePosition sets the underline position of the font.
func (f *Font) SetUnderlinePosition(pos float64) {
	f.obj.SetAttr("upos", pos)
}

// UnderlineWith returns the underline with of the font.
func (f *Font) UnderlineWith() float64 {
	attr := f.obj.GetAttr("uwidth")
	defer attr.DecRef()

	return python3.AsFloat64(attr)
}

// SetUnderlineWith sets the underline with of the font.
func (f *Font) SetUnderlineWith(with float64) {
	f.obj.SetAttr("uwidth", with)
}

// Path returns the path of the font.
func (f *Font) Path() string {
	attr := f.obj.GetAttr("path")
	defer attr.DecRef()

	return python3.AsString(attr)
}

// Copyright returns the copyright.
func (f *Font) Copyright() string {
	attr := f.obj.GetAttr("copyright")
	defer attr.DecRef()

	return python3.AsString(attr)
}

// SetCopyright sets the copyright.
func (f *Font) SetCopyright(copyright string) {
	f.obj.SetAttr("copyright", copyright)
}

// SFNTNames returns the SFNT names.
func (f *Font) SFNTNames() SFNTNames {
	attr := f.obj.GetAttr("sfnt_names")
	defer attr.DecRef()

	return python3.MustUnmarshalAs[SFNTNames](attr)
}

// SetSFNTNames sets the SFNT names.
func (f *Font) SetSFNTNames(keyAndValues ...string) {
	if len(keyAndValues)%2 != 0 {
		panic("key and value must be provided in pair")
	}

	names := f.SFNTNames()
	values := make(map[string]string, len(keyAndValues)%2)

	for i := 0; i < len(keyAndValues); i += 2 {
		values[keyAndValues[i]] = keyAndValues[i+1]
	}

	for j := range names {
		if value, ok := values[names[j].Key]; ok {
			names[j].Value = value
		}
	}

	f.obj.SetAttr("sfnt_names", names)
}

// Version returns the version of the font.
func (f *Font) Version() *semver.Version {
	attr := f.obj.GetAttr("version")
	defer attr.DecRef()

	return parseVersion(python3.AsString(attr))
}

// SetVersion sets the version of the font.
func (f *Font) SetVersion(version semver.Version) {
	f.obj.SetAttr("version", version.String())
}

// HasGlyph returns true if the font has the glyph.
func (f *Font) HasGlyph(glyph any) bool {
	return f.obj.HasItem(glyph)
}

// Glyph returns the glyph of the font.
func (f *Font) Glyph(glyph any) *Glyph {
	o := f.obj.GetItem(glyph)
	if o == nil {
		return nil
	}

	return newGlyph(o)
}

// CreateGlyph creates a new glyph.
func (f *Font) CreateGlyph(glyph string) error {
	f.obj.CallMethodArgs("createChar", -1, glyph)

	return python3.LastError() //nolint: wrapcheck
}

// newFont creates a new Font.
func newFont(obj *python3.Object) *Font {
	return &Font{obj: obj}
}

// SFNTNames is a list of SFNTName.
type SFNTNames []SFNTName

// MarshalPyObject marshals a SFNTName to a python3.Object.
func (n SFNTNames) MarshalPyObject() *python3.Object {
	return python3.NewTupleFromValues(([]SFNTName)(n)...).AsObject()
}

// Find finds a SFNTName by key.
func (n SFNTNames) Find(key string) string {
	for _, name := range n {
		if name.Key == key {
			return name.Value
		}
	}

	return ""
}

// SFNTName is a SFNT name.
type SFNTName struct {
	Locale string
	Key    string
	Value  string
}

// MarshalPyObject marshals a SFNTName to a python3.Object.
func (n SFNTName) MarshalPyObject() *python3.Object {
	return python3.NewTupleFromValues(n.Locale, n.Key, n.Value).AsObject()
}

// UnmarshalPyObject unmarshals a python3.Object to a SFNTName.
func (n *SFNTName) UnmarshalPyObject(o *python3.Object) error { //nolint: unparam
	locale := o.GetItem(0)
	key := o.GetItem(1)
	value := o.GetItem(2)

	defer locale.DecRef()
	defer key.DecRef()
	defer value.DecRef()

	*n = SFNTName{
		Locale: locale.String(),
		Key:    key.String(),
		Value:  value.String(),
	}

	return nil
}

// buildPattern is a regex pattern to detect build id in semver, such as `1.2 build 110`.
var buildPattern = regexp.MustCompile(`\s+build\s+(\d+)$`)

func parseVersion(v string) *semver.Version {
	// Sanitize the version.
	v, _, _ = strings.Cut(v, ";")
	v = strings.TrimSpace(v)
	v = buildPattern.ReplaceAllString(v, "+$1")

	r, _ := semver.NewVersion(v) //nolint: errcheck

	return r
}
