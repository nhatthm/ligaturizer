package ligaturizer

// LigatureName is a font ligature name.
type LigatureName string

// IsZero returns true if ligature name is empty.
func (l LigatureName) IsZero() bool {
	return l == ""
}

// UnmarshalText unmarshal text to Char.
func (l *LigatureName) UnmarshalText(text []byte) error { //nolint: unparam
	*l = LigatureName(text)

	return nil
}

// MarshalText marshal Char to text.
func (l LigatureName) MarshalText() ([]byte, error) { //nolint: unparam
	return []byte(l), nil
}

// String returns a string representation of the Char.
func (l LigatureName) String() string {
	return string(l)
}
