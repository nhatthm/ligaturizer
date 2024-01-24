package firacode

import "go.nhat.io/ligaturizer/internal/ligaturizer"

// Config is the configuration for Fira Code ligaturizer.
type Config struct {
	// Copy glyphs for (some) individual characters from the ligature font as well. This will result in punctuation that matches the ligatures more closely,
	// but may not fit in as well with the rest of the font.
	CopyCharacterGlyphs bool `json:"copy-character-glyphs" toml:"copy-character-glyphs" yaml:"copy-character-glyphs"`
	// When copying character glyphs, if they differ in width from the width of the input font by at least this much, scale them horizontally to match the input
	// font even if this noticeably changes their aspect ratio. The default (0.1) means to scale if they are at least 10%% wider or narrower. A value of 0 will
	// scale all copied character glyphs; a value of 2 effectively disables character glyph scaling.
	ScaleCharacterGlyphsThreshold float64 `json:"scale-character-glyphs-threshold" toml:"scale-character-glyphs-threshold" yaml:"scale-character-glyphs-threshold"`
	// Ligatures is a list of ligatures that the ligaturizer will attempt to copy from Fira Code to the output font. Ligatures that aren't present in Fira Code
	// will be skipped.
	Ligatures []LigatureMapping `json:"ligatures" toml:"ligatures" yaml:"ligatures"`
}

// LigatureMapping is a mapping between a ligature name and the characters that make up the ligature.
type LigatureMapping struct {
	Chars    ligaturizer.Chars        `json:"chars" toml:"chars" yaml:"chars"`
	Ligature ligaturizer.LigatureName `json:"ligature" toml:"ligature" yaml:"ligature"`
}

const defaultScaleCharacterGlyphsThreshold = 0.1

var defaultLigatures = []LigatureMapping{
	{
		// These are all the punctuation characters used in Fira Code ligatures.
		// Use the Config.CopyCharacterGlyphs option to copy these into the output font along with the ligatures themselves.
		Chars: ligaturizer.Chars{
			ligaturizer.CharAmpersand, ligaturizer.CharCircumflex, ligaturizer.CharTilde, ligaturizer.CharAsterisk, ligaturizer.CharBackslash,
			ligaturizer.CharBar, ligaturizer.CharColon, ligaturizer.CharEqual, ligaturizer.CharExclamation, ligaturizer.CharGreater, ligaturizer.CharHyphen,
			ligaturizer.CharLess, ligaturizer.CharNumberSign, ligaturizer.CharPercent, ligaturizer.CharPeriod, ligaturizer.CharPlus, ligaturizer.CharQuestion,
			ligaturizer.CharSemicolon, ligaturizer.CharSlash, ligaturizer.CharUnderscore,

			// These characters are also used by the ligatures, but are likely to look more out of place when spliced into another font.
			// ligaturizer.CharAt, ligaturizer.CharBraceLeft, ligaturizer.CharBraceRight, ligaturizer.CharBracketLeft, ligaturizer.CharBracketRight,
			// ligaturizer.CharDollar, ligaturizer.CharParenthesisLeft, ligaturizer.CharParenthesisRight, ligaturizer.CharLowerW,
		},
	},

	// These are traditional (i.e. present in most variable-width fonts) aesthetic ligatures. They are commented out here so that they don't
	// overwrite similar ligatures present in the destination font.
	// {
	// 	Chars:    []ligaturizer.Char{ligaturizer.CharUpperF, ligaturizer.CharLowerL},
	// 	Ligature: LigatureNameUpperFLowerL,
	// },
	// {
	// 	Chars: []ligaturizer.Char{ligaturizer.CharUpperT, ligaturizer.CharLowerL},
	// 	Ligature: LigatureNameUpperTLowerL,
	// },
	// {
	// 	Chars: []ligaturizer.Char{ligaturizer.CharLowerF, ligaturizer.CharLowerI},
	// 	Ligature: LigatureNameLowerFLowerI,
	// },
	// {
	// 	Chars: []ligaturizer.Char{ligaturizer.CharLowerF, ligaturizer.CharLowerJ},
	// 	Ligature: LigatureNameLowerFLowerJ,
	// },
	// {
	// 	Chars: []ligaturizer.Char{ligaturizer.CharLowerF, ligaturizer.CharLowerL},
	// 	Ligature: LigatureNameLowerFLowerL,
	// },
	// {
	// 	Chars: []ligaturizer.Char{ligaturizer.CharLowerF, ligaturizer.CharLowerT},
	// 	Ligature: LigatureNameLowerFLowerT,
	// },

	// Programming ligatures begin here.
	{
		// &&
		Chars:    []ligaturizer.Char{ligaturizer.CharAmpersand, ligaturizer.CharAmpersand},
		Ligature: LigatureNameAmpersandAmpersand,
	},
	{
		// ^=
		Chars:    []ligaturizer.Char{ligaturizer.CharCircumflex, ligaturizer.CharEqual},
		Ligature: LigatureNameCircumflexEqual,
	},
	{
		// ~~
		Chars:    []ligaturizer.Char{ligaturizer.CharTilde, ligaturizer.CharTilde},
		Ligature: LigatureNameTildeTilde,
	},
	{
		// ~~>
		Chars:    []ligaturizer.Char{ligaturizer.CharTilde, ligaturizer.CharTilde, ligaturizer.CharGreater},
		Ligature: LigatureNameTildeTildeGreater,
	},
	{
		// ~@
		Chars:    []ligaturizer.Char{ligaturizer.CharTilde, ligaturizer.CharAt},
		Ligature: LigatureNameTildeAt,
	},
	{
		// ~=
		Chars:    []ligaturizer.Char{ligaturizer.CharTilde, ligaturizer.CharEqual},
		Ligature: LigatureNameTildeEqual,
	},
	{
		// ~>
		Chars:    []ligaturizer.Char{ligaturizer.CharTilde, ligaturizer.CharGreater},
		Ligature: LigatureNameTildeGreater,
	},
	{
		// ~-
		Chars:    []ligaturizer.Char{ligaturizer.CharTilde, ligaturizer.CharHyphen},
		Ligature: LigatureNameTildeHyphen,
	},
	{
		// **
		Chars:    []ligaturizer.Char{ligaturizer.CharAsterisk, ligaturizer.CharAsterisk},
		Ligature: LigatureNameAsteriskAsterisk,
	},
	{
		// ***
		Chars:    []ligaturizer.Char{ligaturizer.CharAsterisk, ligaturizer.CharAsterisk, ligaturizer.CharAsterisk},
		Ligature: LigatureNameAsteriskAsteriskAsterisk,
	},
	{
		// *>
		Chars:    []ligaturizer.Char{ligaturizer.CharAsterisk, ligaturizer.CharGreater},
		Ligature: LigatureNameAsteriskGreater,
	},
	{
		// */
		Chars:    []ligaturizer.Char{ligaturizer.CharAsterisk, ligaturizer.CharSlash},
		Ligature: LigatureNameAsteriskSlash,
	},
	{
		// \/
		Chars:    []ligaturizer.Char{ligaturizer.CharBackslash, ligaturizer.CharSlash},
		Ligature: LigatureNameBackslashSlash,
	},
	{
		// ||
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharBar},
		Ligature: LigatureNameBarBar,
	},
	{
		// |||>
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharBar, ligaturizer.CharBar, ligaturizer.CharGreater},
		Ligature: LigatureNameBarBarBarGreater,
	},
	{
		// ||=
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharBar, ligaturizer.CharEqual},
		Ligature: LigatureNameBarBarEqual,
	},
	{
		// ||>
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharBar, ligaturizer.CharGreater},
		Ligature: LigatureNameBarBarGreater,
	},
	{
		// ||-
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharBar, ligaturizer.CharHyphen},
		Ligature: LigatureNameBarBarHyphen,
	},
	{
		// |}
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharBraceRight},
		Ligature: LigatureNameBarBraceRight,
	},
	{
		// |]
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharBracketRight},
		Ligature: LigatureNameBarBracketRight,
	},
	{
		// |=
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharEqual},
		Ligature: LigatureNameBarEqual,
	},
	{
		// |=>
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharEqual, ligaturizer.CharGreater},
		Ligature: LigatureNameBarEqualGreater,
	},
	{
		// |>
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharGreater},
		Ligature: LigatureNameBarGreater,
	},
	{
		// |-
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharHyphen},
		Ligature: LigatureNameBarHyphen,
	},
	{
		// |->
		Chars:    []ligaturizer.Char{ligaturizer.CharBar, ligaturizer.CharHyphen, ligaturizer.CharGreater},
		Ligature: LigatureNameBarHyphenGreater,
	},
	{
		// {|
		Chars:    []ligaturizer.Char{ligaturizer.CharBraceLeft, ligaturizer.CharBar},
		Ligature: LigatureNameBraceLeftBar,
	},
	{
		// [|
		Chars:    []ligaturizer.Char{ligaturizer.CharBracketLeft, ligaturizer.CharBar},
		Ligature: LigatureNameBracketLeftBar,
	},
	{
		// ]#
		Chars:    []ligaturizer.Char{ligaturizer.CharBracketRight, ligaturizer.CharNumberSign},
		Ligature: LigatureNameBracketRightNumberSign,
	},
	{
		// ::
		Chars:    []ligaturizer.Char{ligaturizer.CharColon, ligaturizer.CharColon},
		Ligature: LigatureNameColonColon,
	},
	{
		// :::
		Chars:    []ligaturizer.Char{ligaturizer.CharColon, ligaturizer.CharColon, ligaturizer.CharColon},
		Ligature: LigatureNameColonColonColon,
	},
	{
		// ::=
		Chars:    []ligaturizer.Char{ligaturizer.CharColon, ligaturizer.CharColon, ligaturizer.CharEqual},
		Ligature: LigatureNameColonColonEqual,
	},
	{
		// :=
		Chars:    []ligaturizer.Char{ligaturizer.CharColon, ligaturizer.CharEqual},
		Ligature: LigatureNameColonEqual,
	},
	{
		// :>
		Chars:    []ligaturizer.Char{ligaturizer.CharColon, ligaturizer.CharGreater},
		Ligature: LigatureNameColonGreater,
	},
	{
		// :<
		Chars:    []ligaturizer.Char{ligaturizer.CharColon, ligaturizer.CharLess},
		Ligature: LigatureNameColonLess,
	},
	{
		// $>
		Chars:    []ligaturizer.Char{ligaturizer.CharDollar, ligaturizer.CharGreater},
		Ligature: LigatureNameDollarGreater,
	},
	{
		// =:=
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharColon, ligaturizer.CharEqual},
		Ligature: LigatureNameEqualColonEqual,
	},
	{
		// ==
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharEqual},
		Ligature: LigatureNameEqualEqual,
	},
	{
		// ===
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharEqual, ligaturizer.CharEqual},
		Ligature: LigatureNameEqualEqualEqual,
	},
	{
		// ==>
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharEqual, ligaturizer.CharGreater},
		Ligature: LigatureNameEqualEqualGreater,
	},
	{
		// =!=
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharExclamation, ligaturizer.CharEqual},
		Ligature: LigatureNameEqualExclamationEqual,
	},
	{
		// =>
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharGreater},
		Ligature: LigatureNameEqualGreater,
	},
	{
		// =>>
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharGreater, ligaturizer.CharGreater},
		Ligature: LigatureNameEqualGreaterGreater,
	},
	{
		// =<<
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharLess, ligaturizer.CharLess},
		Ligature: LigatureNameEqualLessLess,
	},
	{
		// =/=
		Chars:    []ligaturizer.Char{ligaturizer.CharEqual, ligaturizer.CharSlash, ligaturizer.CharEqual},
		Ligature: LigatureNameEqualSlashEqual,
	},
	{
		// !=
		Chars:    []ligaturizer.Char{ligaturizer.CharExclamation, ligaturizer.CharEqual},
		Ligature: LigatureNameExclamationEqual,
	},
	{
		// !==
		Chars:    []ligaturizer.Char{ligaturizer.CharExclamation, ligaturizer.CharEqual, ligaturizer.CharEqual},
		Ligature: LigatureNameExclamationEqualEqual,
	},
	{
		// !!
		Chars:    []ligaturizer.Char{ligaturizer.CharExclamation, ligaturizer.CharExclamation},
		Ligature: LigatureNameExclamationExclamation,
	},
	{
		// !!.
		Chars:    []ligaturizer.Char{ligaturizer.CharExclamation, ligaturizer.CharExclamation, ligaturizer.CharPeriod},
		Ligature: LigatureNameExclamationExclamationPeriod,
	},
	{
		// >:
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharColon},
		Ligature: LigatureNameGreaterColon,
	},
	{
		// >=
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharEqual},
		Ligature: LigatureNameGreaterEqual,
	},
	{
		// >=>
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharEqual, ligaturizer.CharGreater},
		Ligature: LigatureNameGreaterEqualGreater,
	},
	{
		// >>
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharGreater},
		Ligature: LigatureNameGreaterGreater,
	},
	{
		// >>=
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharGreater, ligaturizer.CharEqual},
		Ligature: LigatureNameGreaterGreaterEqual,
	},
	{
		// >>>
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharGreater, ligaturizer.CharGreater},
		Ligature: LigatureNameGreaterGreaterGreater,
	},
	{
		// >>-
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharGreater, ligaturizer.CharHyphen},
		Ligature: LigatureNameGreaterGreaterHyphen,
	},
	{
		// >-
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharHyphen},
		Ligature: LigatureNameGreaterHyphen,
	},
	{
		// >->
		Chars:    []ligaturizer.Char{ligaturizer.CharGreater, ligaturizer.CharHyphen, ligaturizer.CharGreater},
		Ligature: LigatureNameGreaterHyphenGreater,
	},
	{
		// -~
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharTilde},
		Ligature: LigatureNameHyphenTilde,
	},
	{
		// -|
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharBar},
		Ligature: LigatureNameHyphenBar,
	},
	{
		// ->
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharGreater},
		Ligature: LigatureNameHyphenGreater,
	},
	{
		// ->>
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharGreater, ligaturizer.CharGreater},
		Ligature: LigatureNameHyphenGreaterGreater,
	},
	{
		// --
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharHyphen},
		Ligature: LigatureNameHyphenHyphen,
	},
	{
		// -->
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharHyphen, ligaturizer.CharGreater},
		Ligature: LigatureNameHyphenHyphenGreater,
	},
	{
		// ---
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharHyphen, ligaturizer.CharHyphen},
		Ligature: LigatureNameHyphenHyphenHyphen,
	},
	{
		// -<
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharLess},
		Ligature: LigatureNameHyphenLess,
	},
	{
		// -<<
		Chars:    []ligaturizer.Char{ligaturizer.CharHyphen, ligaturizer.CharLess, ligaturizer.CharLess},
		Ligature: LigatureNameHyphenLessLess,
	},
	{
		// <~
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharTilde},
		Ligature: LigatureNameLessTilde,
	},
	{
		// <~~
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharTilde, ligaturizer.CharTilde},
		Ligature: LigatureNameLessTildeTilde,
	},
	{
		// <~>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharTilde, ligaturizer.CharGreater},
		Ligature: LigatureNameLessTildeGreater,
	},
	{
		// <*
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharAsterisk},
		Ligature: LigatureNameLessAsterisk,
	},
	{
		// <*>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharAsterisk, ligaturizer.CharGreater},
		Ligature: LigatureNameLessAsteriskGreater,
	},
	{
		// <|
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharBar},
		Ligature: LigatureNameLessBar,
	},
	{
		// <||
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharBar, ligaturizer.CharBar},
		Ligature: LigatureNameLessBarBar,
	},
	{
		// <|||
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharBar, ligaturizer.CharBar, ligaturizer.CharBar},
		Ligature: LigatureNameLessBarBarBar,
	},
	{
		// <|>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharBar, ligaturizer.CharGreater},
		Ligature: LigatureNameLessBarGreater,
	},
	{
		// <:
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharColon},
		Ligature: LigatureNameLessColon,
	},
	{
		// <$
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharDollar},
		Ligature: LigatureNameLessDollar,
	},
	{
		// <$>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharDollar, ligaturizer.CharGreater},
		Ligature: LigatureNameLessDollarGreater,
	},
	{
		// <=
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharEqual},
		Ligature: LigatureNameLessEqual,
	},
	{
		// <=|
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharEqual, ligaturizer.CharBar},
		Ligature: LigatureNameLessEqualBar,
	},
	{
		// <==
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharEqual, ligaturizer.CharEqual},
		Ligature: LigatureNameLessEqualEqual,
	},
	{
		// <==>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharEqual, ligaturizer.CharEqual, ligaturizer.CharGreater},
		Ligature: LigatureNameLessEqualEqualGreater,
	},
	{
		// <=>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharEqual, ligaturizer.CharGreater},
		Ligature: LigatureNameLessEqualGreater,
	},
	{
		// <=<
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharEqual, ligaturizer.CharLess},
		Ligature: LigatureNameLessEqualLess,
	},
	{
		// <!--
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharExclamation, ligaturizer.CharHyphen, ligaturizer.CharHyphen},
		Ligature: LigatureNameLessExclamationHyphenHyphen,
	},
	{
		// <>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharGreater},
		Ligature: LigatureNameLessGreater,
	},
	{
		// <-
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharHyphen},
		Ligature: LigatureNameLessHyphen,
	},
	{
		// <-|
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharHyphen, ligaturizer.CharBar},
		Ligature: LigatureNameLessHyphenBar,
	},
	{
		// <->
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharHyphen, ligaturizer.CharGreater},
		Ligature: LigatureNameLessHyphenGreater,
	},
	{
		// <--
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharHyphen, ligaturizer.CharHyphen},
		Ligature: LigatureNameLessHyphenHyphen,
	},
	{
		// <-<
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharHyphen, ligaturizer.CharLess},
		Ligature: LigatureNameLessHyphenLess,
	},
	{
		// <<
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharLess},
		Ligature: LigatureNameLessLess,
	},
	{
		// <<=
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharLess, ligaturizer.CharEqual},
		Ligature: LigatureNameLessLessEqual,
	},
	{
		// <<-
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharLess, ligaturizer.CharHyphen},
		Ligature: LigatureNameLessLessHyphen,
	},
	{
		// <<->>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharLess, ligaturizer.CharHyphen, ligaturizer.CharGreater, ligaturizer.CharGreater},
		Ligature: LigatureNameLessLessHyphenGreaterGreater,
	},
	{
		// <<<
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharLess, ligaturizer.CharLess},
		Ligature: LigatureNameLessLessLess,
	},
	{
		// <+
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharPlus},
		Ligature: LigatureNameLessPlus,
	},
	{
		// <+>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharPlus, ligaturizer.CharGreater},
		Ligature: LigatureNameLessPlusGreater,
	},
	{
		// </
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharSlash},
		Ligature: LigatureNameLessSlash,
	},
	{
		// </>
		Chars:    []ligaturizer.Char{ligaturizer.CharLess, ligaturizer.CharSlash, ligaturizer.CharGreater},
		Ligature: LigatureNameLessSlashGreater,
	},
	{
		// #{
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharBraceLeft},
		Ligature: LigatureNameNumberSignBraceLeft,
	},
	{
		// #[
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharBracketLeft},
		Ligature: LigatureNameNumberSignBracketLeft,
	},
	{
		// #:
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharColon},
		Ligature: LigatureNameNumberSignColon,
	},
	{
		// #=
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharEqual},
		Ligature: LigatureNameNumberSignEqual,
	},
	{
		// #!
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharExclamation},
		Ligature: LigatureNameNumberSignExclamation,
	},
	{
		// ##
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharNumberSign},
		Ligature: LigatureNameNumberSignNumberSign,
	},
	{
		// ###
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharNumberSign, ligaturizer.CharNumberSign},
		Ligature: LigatureNameNumberSignNumberSignNumberSign,
	},
	{
		// ####
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharNumberSign, ligaturizer.CharNumberSign, ligaturizer.CharNumberSign},
		Ligature: LigatureNameNumberSignNumberSignNumberSignNumberSign,
	},
	{
		// #(
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharParenthesisLeft},
		Ligature: LigatureNameNumberSignParenthesisLeft,
	},
	{
		// #?
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharQuestion},
		Ligature: LigatureNameNumberSignQuestion,
	},
	{
		// #_
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharUnderscore},
		Ligature: LigatureNameNumberSignUnderscore,
	},
	{
		// #_(
		Chars:    []ligaturizer.Char{ligaturizer.CharNumberSign, ligaturizer.CharUnderscore, ligaturizer.CharParenthesisLeft},
		Ligature: LigatureNameNumberSignUnderscoreParenthesisLeft,
	},
	{
		// %%
		Chars:    []ligaturizer.Char{ligaturizer.CharPercent, ligaturizer.CharPercent},
		Ligature: LigatureNamePercentPercent,
	},
	{
		// .=
		Chars:    []ligaturizer.Char{ligaturizer.CharPeriod, ligaturizer.CharEqual},
		Ligature: LigatureNamePeriodEqual,
	},
	{
		// .-
		Chars:    []ligaturizer.Char{ligaturizer.CharPeriod, ligaturizer.CharHyphen},
		Ligature: LigatureNamePeriodHyphen,
	},
	{
		// ..
		Chars:    []ligaturizer.Char{ligaturizer.CharPeriod, ligaturizer.CharPeriod},
		Ligature: LigatureNamePeriodPeriod,
	},
	{
		// ..=
		Chars:    []ligaturizer.Char{ligaturizer.CharPeriod, ligaturizer.CharPeriod, ligaturizer.CharEqual},
		Ligature: LigatureNamePeriodPeriodEqual,
	},
	{
		// ..<
		Chars:    []ligaturizer.Char{ligaturizer.CharPeriod, ligaturizer.CharPeriod, ligaturizer.CharLess},
		Ligature: LigatureNamePeriodPeriodLess,
	},
	{
		// ...
		Chars:    []ligaturizer.Char{ligaturizer.CharPeriod, ligaturizer.CharPeriod, ligaturizer.CharPeriod},
		Ligature: LigatureNamePeriodPeriodPeriod,
	},
	{
		// .?
		Chars:    []ligaturizer.Char{ligaturizer.CharPeriod, ligaturizer.CharQuestion},
		Ligature: LigatureNamePeriodQuestion,
	},
	{
		// +>
		Chars:    []ligaturizer.Char{ligaturizer.CharPlus, ligaturizer.CharGreater},
		Ligature: LigatureNamePlusGreater,
	},
	{
		// ++
		Chars:    []ligaturizer.Char{ligaturizer.CharPlus, ligaturizer.CharPlus},
		Ligature: LigatureNamePlusPlus,
	},
	{
		// +++
		Chars:    []ligaturizer.Char{ligaturizer.CharPlus, ligaturizer.CharPlus, ligaturizer.CharPlus},
		Ligature: LigatureNamePlusPlusPlus,
	},
	{
		// ?:
		Chars:    []ligaturizer.Char{ligaturizer.CharQuestion, ligaturizer.CharColon},
		Ligature: LigatureNameQuestionColon,
	},
	{
		// ?=
		Chars:    []ligaturizer.Char{ligaturizer.CharQuestion, ligaturizer.CharEqual},
		Ligature: LigatureNameQuestionEqual,
	},
	{
		// ?.
		Chars:    []ligaturizer.Char{ligaturizer.CharQuestion, ligaturizer.CharPeriod},
		Ligature: LigatureNameQuestionPeriod,
	},
	{
		// ??
		Chars:    []ligaturizer.Char{ligaturizer.CharQuestion, ligaturizer.CharQuestion},
		Ligature: LigatureNameQuestionQuestion,
	},
	{
		// ;;
		Chars:    []ligaturizer.Char{ligaturizer.CharSemicolon, ligaturizer.CharSemicolon},
		Ligature: LigatureNameSemicolonSemicolon,
	},
	{
		// /*
		Chars:    []ligaturizer.Char{ligaturizer.CharSlash, ligaturizer.CharAsterisk},
		Ligature: LigatureNameSlashAsterisk,
	},
	{
		// /\
		Chars:    []ligaturizer.Char{ligaturizer.CharSlash, ligaturizer.CharBackslash},
		Ligature: LigatureNameSlashBackslash,
	},
	{
		// /=
		Chars:    []ligaturizer.Char{ligaturizer.CharSlash, ligaturizer.CharEqual},
		Ligature: LigatureNameSlashEqual,
	},
	{
		// /==
		Chars:    []ligaturizer.Char{ligaturizer.CharSlash, ligaturizer.CharEqual, ligaturizer.CharEqual},
		Ligature: LigatureNameSlashEqualEqual,
	},
	{
		// />
		Chars:    []ligaturizer.Char{ligaturizer.CharSlash, ligaturizer.CharGreater},
		Ligature: LigatureNameSlashGreater,
	},
	{
		// //
		Chars:    []ligaturizer.Char{ligaturizer.CharSlash, ligaturizer.CharSlash},
		Ligature: LigatureNameSlashSlash,
	},
	{
		// ///
		Chars:    []ligaturizer.Char{ligaturizer.CharSlash, ligaturizer.CharSlash, ligaturizer.CharSlash},
		Ligature: LigatureNameSlashSlashSlash,
	},
	{
		// _|_
		Chars:    []ligaturizer.Char{ligaturizer.CharUnderscore, ligaturizer.CharBar, ligaturizer.CharUnderscore},
		Ligature: LigatureNameUnderscoreBarUnderscore,
	},
	{
		// __
		Chars:    []ligaturizer.Char{ligaturizer.CharUnderscore, ligaturizer.CharUnderscore},
		Ligature: LigatureNameUnderscoreUnderscore,
	},
	{
		// www
		Chars:    []ligaturizer.Char{ligaturizer.CharLowerW, ligaturizer.CharLowerW, ligaturizer.CharLowerW},
		Ligature: LigatureNameLowerWWW,
	},
}
