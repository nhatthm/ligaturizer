package fontforge

import "go.nhat.io/python3"

// Glyph is a font glyph.
type Glyph struct {
	obj *python3.Object
}

// callMethodArgs calls a method with args.
func (g *Glyph) callMethodArgs(name string, args ...any) *python3.Object {
	return g.obj.CallMethodArgs(name, args...)
}

// Close closes the glyph.
func (g *Glyph) Close() error { //nolint: unparam
	g.obj.DecRef()

	return nil
}

// DecRef decreases the reference count of the object.
func (g *Glyph) DecRef() {
	g.obj.DecRef()
}

// AsObject returns the underlying Object.
func (g *Glyph) AsObject() *python3.Object {
	return g.obj
}

// GlyphName returns the name of the glyph.
func (g *Glyph) GlyphName() string {
	attr := g.obj.GetAttr("glyphname")
	defer attr.DecRef()

	return python3.AsString(attr)
}

// SetGlyphName sets the name of the glyph.
func (g *Glyph) SetGlyphName(name string) {
	g.obj.SetAttr("glyphname", name)
}

// Width returns the width of the glyph.
func (g *Glyph) Width() int {
	attr := g.obj.GetAttr("width")
	defer attr.DecRef()

	return python3.AsInt(attr)
}

// SetWidth sets the width of the glyph.
func (g *Glyph) SetWidth(width int) {
	g.obj.SetAttr("width", width)
}

// LeftSideBearing returns the left side bearing of the glyph.
func (g *Glyph) LeftSideBearing() float64 {
	attr := g.obj.GetAttr("left_side_bearing")
	defer attr.DecRef()

	return python3.AsFloat64(attr)
}

// SetLeftSideBearing sets the left side bearing of the glyph.
func (g *Glyph) SetLeftSideBearing(bearing float64) {
	g.obj.SetAttr("left_side_bearing", bearing)
}

// RightSideBearing returns the right side bearing of the glyph.
func (g *Glyph) RightSideBearing() float64 {
	attr := g.obj.GetAttr("right_side_bearing")
	defer attr.DecRef()

	return python3.AsFloat64(attr)
}

// SetRightSideBearing sets the right side bearing of the glyph.
func (g *Glyph) SetRightSideBearing(bearing float64) {
	g.obj.SetAttr("right_side_bearing", bearing)
}

// Transform applies a transformation matrix to the glyph.
func (g *Glyph) Transform(matrix []float64) {
	tuple := python3.NewTupleFromValues(matrix...)
	defer tuple.DecRef()

	g.obj.CallMethodArgs("transform", tuple)
}

// newGlyph creates a new Glyph.
func newGlyph(obj *python3.Object) *Glyph {
	return &Glyph{obj: obj}
}
