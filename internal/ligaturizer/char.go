package ligaturizer

import (
	"errors"
	"fmt"
)

// ErrCharNotSupported is returned when char is not supported.
var ErrCharNotSupported = errors.New("char not supported")

// Char constants.
var (
	CharAmpersand        = Char{Name: "ampersand", Rune: '&'}
	CharCircumflex       = Char{Name: "asciicircum", Rune: '^'}
	CharTilde            = Char{Name: "asciitilde", Rune: '~'}
	CharAsterisk         = Char{Name: "asterisk", Rune: '*'}
	CharBackslash        = Char{Name: "backslash", Rune: '\\'}
	CharBar              = Char{Name: "bar", Rune: '|'}
	CharColon            = Char{Name: "colon", Rune: ':'}
	CharEqual            = Char{Name: "equal", Rune: '='}
	CharExclamation      = Char{Name: "exclam", Rune: '!'}
	CharGreater          = Char{Name: "greater", Rune: '>'}
	CharHyphen           = Char{Name: "hyphen", Rune: '-'}
	CharLess             = Char{Name: "less", Rune: '<'}
	CharNumberSign       = Char{Name: "numbersign", Rune: '#'}
	CharPercent          = Char{Name: "percent", Rune: '%'}
	CharPeriod           = Char{Name: "period", Rune: '.'}
	CharPlus             = Char{Name: "plus", Rune: '+'}
	CharQuestion         = Char{Name: "question", Rune: '?'}
	CharSemicolon        = Char{Name: "semicolon", Rune: ';'}
	CharSlash            = Char{Name: "slash", Rune: '/'}
	CharUnderscore       = Char{Name: "underscore", Rune: '_'}
	CharAt               = Char{Name: "at", Rune: '@'}
	CharBraceLeft        = Char{Name: "braceleft", Rune: '{'}
	CharBraceRight       = Char{Name: "braceright", Rune: '}'}
	CharBracketLeft      = Char{Name: "bracketleft", Rune: '['}
	CharBracketRight     = Char{Name: "bracketright", Rune: ']'}
	CharDollar           = Char{Name: "dollar", Rune: '$'}
	CharParenthesisLeft  = Char{Name: "parenleft", Rune: '('}
	CharParenthesisRight = Char{Name: "parenright", Rune: ')'}
	CharUpperF           = Char{Name: "F", Rune: 'F'}
	CharUpperT           = Char{Name: "T", Rune: 'T'}
	CharLowerF           = Char{Name: "f", Rune: 'f'}
	CharLowerI           = Char{Name: "i", Rune: 'i'}
	CharLowerJ           = Char{Name: "j", Rune: 'j'}
	CharLowerL           = Char{Name: "l", Rune: 'l'}
	CharLowerT           = Char{Name: "t", Rune: 't'}
	CharLowerW           = Char{Name: "w", Rune: 'w'}
)

// Chars is a list of Char.
type Chars []Char

// Names returns a list of names.
func (cs Chars) Names() []string {
	names := make([]string, len(cs))

	for i, c := range cs {
		names[i] = c.Name
	}

	return names
}

// Char is a character.
type Char struct {
	Name string
	Rune rune
}

// UnmarshalText unmarshal text to Char.
func (c *Char) UnmarshalText(text []byte) error { //nolint: unparam
	r, ok := charAliases[string(text)]
	if !ok {
		return fmt.Errorf("%w: %q", ErrCharNotSupported, text)
	}

	*c = r

	return nil
}

// MarshalText marshal Char to text.
func (c Char) MarshalText() ([]byte, error) { //nolint: unparam
	return []byte(c.Name), nil
}

// String returns a string representation of the Char.
func (c Char) String() string {
	return string(c.Rune)
}

var charAliases = map[string]Char{
	"ampersand":        CharAmpersand,
	"asciicircum":      CharCircumflex,
	"circumflex":       CharCircumflex,
	"asciitilde":       CharTilde,
	"tilde":            CharTilde,
	"asterisk":         CharAsterisk,
	"backslash":        CharBackslash,
	"bar":              CharBar,
	"colon":            CharColon,
	"equal":            CharEqual,
	"exclam":           CharExclamation,
	"exclamation":      CharExclamation,
	"greater":          CharGreater,
	"hyphen":           CharHyphen,
	"less":             CharLess,
	"numbersign":       CharNumberSign,
	"percent":          CharPercent,
	"period":           CharPeriod,
	"plus":             CharPlus,
	"question":         CharQuestion,
	"semicolon":        CharSemicolon,
	"slash":            CharSlash,
	"underscore":       CharUnderscore,
	"at":               CharAt,
	"braceleft":        CharBraceLeft,
	"braceright":       CharBraceRight,
	"bracketleft":      CharBracketLeft,
	"bracketright":     CharBracketRight,
	"dollar":           CharDollar,
	"parenleft":        CharParenthesisLeft,
	"parenthesisleft":  CharParenthesisLeft,
	"parenright":       CharParenthesisRight,
	"parenthesisright": CharParenthesisRight,

	CharAmpersand.String():        CharAmpersand,
	CharCircumflex.String():       CharCircumflex,
	CharTilde.String():            CharTilde,
	CharAsterisk.String():         CharAsterisk,
	CharBackslash.String():        CharBackslash,
	CharBar.String():              CharBar,
	CharColon.String():            CharColon,
	CharEqual.String():            CharEqual,
	CharExclamation.String():      CharExclamation,
	CharGreater.String():          CharGreater,
	CharHyphen.String():           CharHyphen,
	CharLess.String():             CharLess,
	CharNumberSign.String():       CharNumberSign,
	CharPercent.String():          CharPercent,
	CharPeriod.String():           CharPeriod,
	CharPlus.String():             CharPlus,
	CharQuestion.String():         CharQuestion,
	CharSemicolon.String():        CharSemicolon,
	CharSlash.String():            CharSlash,
	CharUnderscore.String():       CharUnderscore,
	CharAt.String():               CharAt,
	CharBraceLeft.String():        CharBraceLeft,
	CharBraceRight.String():       CharBraceRight,
	CharBracketLeft.String():      CharBracketLeft,
	CharBracketRight.String():     CharBracketRight,
	CharDollar.String():           CharDollar,
	CharParenthesisLeft.String():  CharParenthesisLeft,
	CharParenthesisRight.String(): CharParenthesisRight,
	CharUpperF.String():           CharUpperF,
	CharUpperT.String():           CharUpperT,
	CharLowerF.String():           CharLowerF,
	CharLowerI.String():           CharLowerI,
	CharLowerJ.String():           CharLowerJ,
	CharLowerL.String():           CharLowerL,
	CharLowerT.String():           CharLowerT,
	CharLowerW.String():           CharLowerW,
}
