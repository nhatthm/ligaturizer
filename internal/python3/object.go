package python3

import "go.nhat.io/cpy3"

// PyObject is a wrapper around the C type python3.PyObject.
type PyObject = cpy3.PyObject

// Object is a wrapper around the C type python3.PyObject.
type Object struct {
	obj *PyObject
}

// DecRef decreases the reference count of the object.
func (o *Object) DecRef() {
	if o == nil {
		return
	}

	o.obj.DecRef()
}

// PyObject returns the underlying PyObject.
func (o *Object) PyObject() *PyObject {
	return o.obj
}

// Type returns the type of the object.
func (o *Object) Type() *Object {
	return NewObject(o.obj.Type())
}

// Length returns the length of the object.
func (o *Object) Length() int {
	return o.obj.Length()
}

// CallMethodArgs calls a method of the object.
func (o *Object) CallMethodArgs(name string, args ...any) *Object {
	oArgs := make([]*PyObject, len(args))
	for i, arg := range args {
		oArgs[i] = ToPyObject(arg)
	}

	return NewObject(o.obj.CallMethodArgs(name, oArgs...))
}

// GetItem returns the item of the object.
func (o *Object) GetItem(key any) *Object {
	return NewObject(o.obj.GetItem(ToPyObject(key)))
}

// HasItem returns true if the object has the item.
func (o *Object) HasItem(value any) bool {
	return cpy3.PySequence_Contains(o.obj, ToPyObject(value)) == 1
}

// GetAttr returns the attribute value of the object.
func (o *Object) GetAttr(name string) *Object {
	return NewObject(o.obj.GetAttrString(name))
}

// SetAttr sets the attribute value of the object.
func (o *Object) SetAttr(name string, value any) {
	o.obj.SetAttrString(name, ToPyObject(value))
}

// String returns the string representation of the object.
func (o *Object) String() string {
	return AsString(o)
}

// NewObject creates a new Object.
func NewObject(obj *PyObject) *Object {
	if obj == nil {
		return nil
	}

	return &Object{obj: obj}
}

// ToPyObject converts a value to a PyObject.
func ToPyObject(v any) *PyObject {
	switch v := v.(type) {
	case *PyObject:
		return v

	case string:
		return PyString(v)

	case rune:
		return PyInt(int(v))

	case float64:
		return PyFloat64(v)

	case int:
		return PyInt(v)

	case interface{ PyObject() *PyObject }:
		return v.PyObject()

	default:
		panic("unknown type")
	}
}
