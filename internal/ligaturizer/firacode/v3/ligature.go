package firacode

import (
	"go.nhat.io/ligaturizer/internal/ligaturizer"
)

// Ligature names.
const (
	NoLigatureName                                       = ligaturizer.LigatureName("")
	LigatureNameAmpersandAmpersand                       = ligaturizer.LigatureName("ampersand_ampersand.liga")
	LigatureNameCircumflexEqual                          = ligaturizer.LigatureName("asciicircum_equal.liga")
	LigatureNameTildeTildeGreater                        = ligaturizer.LigatureName("asciitilde_asciitilde_greater.liga")
	LigatureNameTildeTilde                               = ligaturizer.LigatureName("asciitilde_asciitilde.liga")
	LigatureNameTildeAt                                  = ligaturizer.LigatureName("asciitilde_at.liga")
	LigatureNameTildeEqual                               = ligaturizer.LigatureName("asciitilde_equal.liga")
	LigatureNameTildeGreater                             = ligaturizer.LigatureName("asciitilde_greater.liga")
	LigatureNameTildeHyphen                              = ligaturizer.LigatureName("asciitilde_hyphen.liga")
	LigatureNameAsteriskAsteriskAsterisk                 = ligaturizer.LigatureName("asterisk_asterisk_asterisk.liga")
	LigatureNameAsteriskAsterisk                         = ligaturizer.LigatureName("asterisk_asterisk.liga")
	LigatureNameAsteriskGreater                          = ligaturizer.LigatureName("asterisk_greater.liga")
	LigatureNameAsteriskSlash                            = ligaturizer.LigatureName("asterisk_slash.liga")
	LigatureNameBackslashSlash                           = ligaturizer.LigatureName("backslash_slash.liga")
	LigatureNameBarBarBarGreater                         = ligaturizer.LigatureName("bar_bar_bar_greater.liga")
	LigatureNameBarBarEqual                              = ligaturizer.LigatureName("bar_bar_equal.liga")
	LigatureNameBarBarGreater                            = ligaturizer.LigatureName("bar_bar_greater.liga")
	LigatureNameBarBarHyphen                             = ligaturizer.LigatureName("bar_bar_hyphen.liga")
	LigatureNameBarBar                                   = ligaturizer.LigatureName("bar_bar.liga")
	LigatureNameBarBraceRight                            = ligaturizer.LigatureName("bar_braceright.liga")
	LigatureNameBarBracketRight                          = ligaturizer.LigatureName("bar_bracketright.liga")
	LigatureNameBarEqualGreater                          = ligaturizer.LigatureName("bar_equal_greater.liga")
	LigatureNameBarEqual                                 = ligaturizer.LigatureName("bar_equal.liga")
	LigatureNameBarGreater                               = ligaturizer.LigatureName("bar_greater.liga")
	LigatureNameBarHyphenGreater                         = ligaturizer.LigatureName("bar_hyphen_greater.liga")
	LigatureNameBarHyphen                                = ligaturizer.LigatureName("bar_hyphen.liga")
	LigatureNameBraceLeftBar                             = ligaturizer.LigatureName("braceleft_bar.liga")
	LigatureNameBracketLeftBar                           = ligaturizer.LigatureName("bracketleft_bar.liga")
	LigatureNameBracketRightNumberSign                   = ligaturizer.LigatureName("bracketright_numbersign.liga")
	LigatureNameColonColonColon                          = ligaturizer.LigatureName("colon_colon_colon.liga")
	LigatureNameColonColonEqual                          = ligaturizer.LigatureName("colon_colon_equal.liga")
	LigatureNameColonColon                               = ligaturizer.LigatureName("colon_colon.liga")
	LigatureNameColonEqual                               = ligaturizer.LigatureName("colon_equal.liga")
	LigatureNameColonGreater                             = ligaturizer.LigatureName("colon_greater.liga")
	LigatureNameColonLess                                = ligaturizer.LigatureName("colon_less.liga")
	LigatureNameDollarGreater                            = ligaturizer.LigatureName("dollar_greater.liga")
	LigatureNameEqualColonEqual                          = ligaturizer.LigatureName("equal_colon_equal.liga")
	LigatureNameEqualEqualEqual                          = ligaturizer.LigatureName("equal_equal_equal.liga")
	LigatureNameEqualEqualGreater                        = ligaturizer.LigatureName("equal_equal_greater.liga")
	LigatureNameEqualEqual                               = ligaturizer.LigatureName("equal_equal.liga")
	LigatureNameEqualExclamationEqual                    = ligaturizer.LigatureName("equal_exclam_equal.liga")
	LigatureNameEqualGreaterGreater                      = ligaturizer.LigatureName("equal_greater_greater.liga")
	LigatureNameEqualGreater                             = ligaturizer.LigatureName("equal_greater.liga")
	LigatureNameEqualLessLess                            = ligaturizer.LigatureName("equal_less_less.liga")
	LigatureNameEqualSlashEqual                          = ligaturizer.LigatureName("equal_slash_equal.liga")
	LigatureNameExclamationEqualEqual                    = ligaturizer.LigatureName("exclam_equal_equal.liga")
	LigatureNameExclamationEqual                         = ligaturizer.LigatureName("exclam_equal.liga")
	LigatureNameExclamationExclamationPeriod             = ligaturizer.LigatureName("exclam_exclam_period.liga")
	LigatureNameExclamationExclamation                   = ligaturizer.LigatureName("exclam_exclam.liga")
	LigatureNameLowerFLowerI                             = ligaturizer.LigatureName("f_i.liga")
	LigatureNameLowerFLowerJ                             = ligaturizer.LigatureName("f_j.liga")
	LigatureNameUpperFLowerL                             = ligaturizer.LigatureName("F_l.liga")
	LigatureNameLowerFLowerL                             = ligaturizer.LigatureName("f_l.liga")
	LigatureNameLowerFLowerT                             = ligaturizer.LigatureName("f_t.liga")
	LigatureNameUpperTLowerL                             = ligaturizer.LigatureName("T_l.liga")
	LigatureNameGreaterColon                             = ligaturizer.LigatureName("greater_colon.liga")
	LigatureNameGreaterEqualGreater                      = ligaturizer.LigatureName("greater_equal_greater.liga")
	LigatureNameGreaterEqual                             = ligaturizer.LigatureName("greater_equal.liga")
	LigatureNameGreaterGreaterEqual                      = ligaturizer.LigatureName("greater_greater_equal.liga")
	LigatureNameGreaterGreaterGreater                    = ligaturizer.LigatureName("greater_greater_greater.liga")
	LigatureNameGreaterGreaterHyphen                     = ligaturizer.LigatureName("greater_greater_hyphen.liga")
	LigatureNameGreaterGreater                           = ligaturizer.LigatureName("greater_greater.liga")
	LigatureNameGreaterHyphenGreater                     = ligaturizer.LigatureName("greater_hyphen_greater.liga")
	LigatureNameGreaterHyphen                            = ligaturizer.LigatureName("greater_hyphen.liga")
	LigatureNameHyphenTilde                              = ligaturizer.LigatureName("hyphen_asciitilde.liga")
	LigatureNameHyphenBar                                = ligaturizer.LigatureName("hyphen_bar.liga")
	LigatureNameHyphenGreaterGreater                     = ligaturizer.LigatureName("hyphen_greater_greater.liga")
	LigatureNameHyphenGreater                            = ligaturizer.LigatureName("hyphen_greater.liga")
	LigatureNameHyphenHyphenGreater                      = ligaturizer.LigatureName("hyphen_hyphen_greater.liga")
	LigatureNameHyphenHyphenHyphen                       = ligaturizer.LigatureName("hyphen_hyphen_hyphen.liga")
	LigatureNameHyphenHyphen                             = ligaturizer.LigatureName("hyphen_hyphen.liga")
	LigatureNameHyphenLessLess                           = ligaturizer.LigatureName("hyphen_less_less.liga")
	LigatureNameHyphenLess                               = ligaturizer.LigatureName("hyphen_less.liga")
	LigatureNameLessTildeTilde                           = ligaturizer.LigatureName("less_asciitilde_asciitilde.liga")
	LigatureNameLessTildeGreater                         = ligaturizer.LigatureName("less_asciitilde_greater.liga")
	LigatureNameLessTilde                                = ligaturizer.LigatureName("less_asciitilde.liga")
	LigatureNameLessAsteriskGreater                      = ligaturizer.LigatureName("less_asterisk_greater.liga")
	LigatureNameLessAsterisk                             = ligaturizer.LigatureName("less_asterisk.liga")
	LigatureNameLessBarBarBar                            = ligaturizer.LigatureName("less_bar_bar_bar.liga")
	LigatureNameLessBarBar                               = ligaturizer.LigatureName("less_bar_bar.liga")
	LigatureNameLessBarGreater                           = ligaturizer.LigatureName("less_bar_greater.liga")
	LigatureNameLessBar                                  = ligaturizer.LigatureName("less_bar.liga")
	LigatureNameLessColon                                = ligaturizer.LigatureName("less_colon.liga")
	LigatureNameLessDollarGreater                        = ligaturizer.LigatureName("less_dollar_greater.liga")
	LigatureNameLessDollar                               = ligaturizer.LigatureName("less_dollar.liga")
	LigatureNameLessEqualBar                             = ligaturizer.LigatureName("less_equal_bar.liga")
	LigatureNameLessEqualEqualGreater                    = ligaturizer.LigatureName("less_equal_equal_greater.liga")
	LigatureNameLessEqualEqual                           = ligaturizer.LigatureName("less_equal_equal.liga")
	LigatureNameLessEqualGreater                         = ligaturizer.LigatureName("less_equal_greater.liga")
	LigatureNameLessEqualLess                            = ligaturizer.LigatureName("less_equal_less.liga")
	LigatureNameLessEqual                                = ligaturizer.LigatureName("less_equal.liga")
	LigatureNameLessExclamationHyphenHyphen              = ligaturizer.LigatureName("less_exclam_hyphen_hyphen.liga")
	LigatureNameLessGreater                              = ligaturizer.LigatureName("less_greater.liga")
	LigatureNameLessHyphenBar                            = ligaturizer.LigatureName("less_hyphen_bar.liga")
	LigatureNameLessHyphenGreater                        = ligaturizer.LigatureName("less_hyphen_greater.liga")
	LigatureNameLessHyphenHyphen                         = ligaturizer.LigatureName("less_hyphen_hyphen.liga")
	LigatureNameLessHyphenLess                           = ligaturizer.LigatureName("less_hyphen_less.liga")
	LigatureNameLessHyphen                               = ligaturizer.LigatureName("less_hyphen.liga")
	LigatureNameLessLessEqual                            = ligaturizer.LigatureName("less_less_equal.liga")
	LigatureNameLessLessHyphenGreaterGreater             = ligaturizer.LigatureName("less_less_hyphen_greater_greater.liga")
	LigatureNameLessLessHyphen                           = ligaturizer.LigatureName("less_less_hyphen.liga")
	LigatureNameLessLessLess                             = ligaturizer.LigatureName("less_less_less.liga")
	LigatureNameLessLess                                 = ligaturizer.LigatureName("less_less.liga")
	LigatureNameLessPlusGreater                          = ligaturizer.LigatureName("less_plus_greater.liga")
	LigatureNameLessPlus                                 = ligaturizer.LigatureName("less_plus.liga")
	LigatureNameLessSlashGreater                         = ligaturizer.LigatureName("less_slash_greater.liga")
	LigatureNameLessSlash                                = ligaturizer.LigatureName("less_slash.liga")
	LigatureNameNumberSignBraceLeft                      = ligaturizer.LigatureName("numbersign_braceleft.liga")
	LigatureNameNumberSignBracketLeft                    = ligaturizer.LigatureName("numbersign_bracketleft.liga")
	LigatureNameNumberSignColon                          = ligaturizer.LigatureName("numbersign_colon.liga")
	LigatureNameNumberSignEqual                          = ligaturizer.LigatureName("numbersign_equal.liga")
	LigatureNameNumberSignExclamation                    = ligaturizer.LigatureName("numbersign_exclam.liga")
	LigatureNameNumberSignNumberSignNumberSignNumberSign = ligaturizer.LigatureName("numbersign_numbersign_numbersign_numbersign.liga")
	LigatureNameNumberSignNumberSignNumberSign           = ligaturizer.LigatureName("numbersign_numbersign_numbersign.liga")
	LigatureNameNumberSignNumberSign                     = ligaturizer.LigatureName("numbersign_numbersign.liga")
	LigatureNameNumberSignParenthesisLeft                = ligaturizer.LigatureName("numbersign_parenleft.liga")
	LigatureNameNumberSignQuestion                       = ligaturizer.LigatureName("numbersign_question.liga")
	LigatureNameNumberSignUnderscoreParenthesisLeft      = ligaturizer.LigatureName("numbersign_underscore_parenleft.liga")
	LigatureNameNumberSignUnderscore                     = ligaturizer.LigatureName("numbersign_underscore.liga")
	LigatureNamePercentPercent                           = ligaturizer.LigatureName("percent_percent.liga")
	LigatureNamePeriodEqual                              = ligaturizer.LigatureName("period_equal.liga")
	LigatureNamePeriodHyphen                             = ligaturizer.LigatureName("period_hyphen.liga")
	LigatureNamePeriodPeriodEqual                        = ligaturizer.LigatureName("period_period_equal.liga")
	LigatureNamePeriodPeriodLess                         = ligaturizer.LigatureName("period_period_less.liga")
	LigatureNamePeriodPeriodPeriod                       = ligaturizer.LigatureName("period_period_period.liga")
	LigatureNamePeriodPeriod                             = ligaturizer.LigatureName("period_period.liga")
	LigatureNamePeriodQuestion                           = ligaturizer.LigatureName("period_question.liga")
	LigatureNamePlusGreater                              = ligaturizer.LigatureName("plus_greater.liga")
	LigatureNamePlusPlusPlus                             = ligaturizer.LigatureName("plus_plus_plus.liga")
	LigatureNamePlusPlus                                 = ligaturizer.LigatureName("plus_plus.liga")
	LigatureNameQuestionColon                            = ligaturizer.LigatureName("question_colon.liga")
	LigatureNameQuestionEqual                            = ligaturizer.LigatureName("question_equal.liga")
	LigatureNameQuestionPeriod                           = ligaturizer.LigatureName("question_period.liga")
	LigatureNameQuestionQuestion                         = ligaturizer.LigatureName("question_question.liga")
	LigatureNameSemicolonSemicolon                       = ligaturizer.LigatureName("semicolon_semicolon.liga")
	LigatureNameSlashAsterisk                            = ligaturizer.LigatureName("slash_asterisk.liga")
	LigatureNameSlashBackslash                           = ligaturizer.LigatureName("slash_backslash.liga")
	LigatureNameSlashEqualEqual                          = ligaturizer.LigatureName("slash_equal_equal.liga")
	LigatureNameSlashEqual                               = ligaturizer.LigatureName("slash_equal.liga")
	LigatureNameSlashGreater                             = ligaturizer.LigatureName("slash_greater.liga")
	LigatureNameSlashSlashSlash                          = ligaturizer.LigatureName("slash_slash_slash.liga")
	LigatureNameSlashSlash                               = ligaturizer.LigatureName("slash_slash.liga")
	LigatureNameUnderscoreBarUnderscore                  = ligaturizer.LigatureName("underscore_bar_underscore.liga")
	LigatureNameUnderscoreUnderscore                     = ligaturizer.LigatureName("underscore_underscore.liga")
	LigatureNameLowerWWW                                 = ligaturizer.LigatureName("w_w_w.liga")
)

var availableLigatureNames = map[ligaturizer.LigatureName]struct{}{
	NoLigatureName:                                       {},
	LigatureNameAmpersandAmpersand:                       {},
	LigatureNameCircumflexEqual:                          {},
	LigatureNameTildeTildeGreater:                        {},
	LigatureNameTildeTilde:                               {},
	LigatureNameTildeAt:                                  {},
	LigatureNameTildeEqual:                               {},
	LigatureNameTildeGreater:                             {},
	LigatureNameTildeHyphen:                              {},
	LigatureNameAsteriskAsteriskAsterisk:                 {},
	LigatureNameAsteriskAsterisk:                         {},
	LigatureNameAsteriskGreater:                          {},
	LigatureNameAsteriskSlash:                            {},
	LigatureNameBackslashSlash:                           {},
	LigatureNameBarBarBarGreater:                         {},
	LigatureNameBarBarEqual:                              {},
	LigatureNameBarBarGreater:                            {},
	LigatureNameBarBarHyphen:                             {},
	LigatureNameBarBar:                                   {},
	LigatureNameBarBraceRight:                            {},
	LigatureNameBarBracketRight:                          {},
	LigatureNameBarEqualGreater:                          {},
	LigatureNameBarEqual:                                 {},
	LigatureNameBarGreater:                               {},
	LigatureNameBarHyphenGreater:                         {},
	LigatureNameBarHyphen:                                {},
	LigatureNameBraceLeftBar:                             {},
	LigatureNameBracketLeftBar:                           {},
	LigatureNameBracketRightNumberSign:                   {},
	LigatureNameColonColonColon:                          {},
	LigatureNameColonColonEqual:                          {},
	LigatureNameColonColon:                               {},
	LigatureNameColonEqual:                               {},
	LigatureNameColonGreater:                             {},
	LigatureNameColonLess:                                {},
	LigatureNameDollarGreater:                            {},
	LigatureNameEqualColonEqual:                          {},
	LigatureNameEqualEqualEqual:                          {},
	LigatureNameEqualEqualGreater:                        {},
	LigatureNameEqualEqual:                               {},
	LigatureNameEqualExclamationEqual:                    {},
	LigatureNameEqualGreaterGreater:                      {},
	LigatureNameEqualGreater:                             {},
	LigatureNameEqualLessLess:                            {},
	LigatureNameEqualSlashEqual:                          {},
	LigatureNameExclamationEqualEqual:                    {},
	LigatureNameExclamationEqual:                         {},
	LigatureNameExclamationExclamationPeriod:             {},
	LigatureNameExclamationExclamation:                   {},
	LigatureNameLowerFLowerI:                             {},
	LigatureNameLowerFLowerJ:                             {},
	LigatureNameUpperFLowerL:                             {},
	LigatureNameLowerFLowerL:                             {},
	LigatureNameLowerFLowerT:                             {},
	LigatureNameUpperTLowerL:                             {},
	LigatureNameGreaterColon:                             {},
	LigatureNameGreaterEqualGreater:                      {},
	LigatureNameGreaterEqual:                             {},
	LigatureNameGreaterGreaterEqual:                      {},
	LigatureNameGreaterGreaterGreater:                    {},
	LigatureNameGreaterGreaterHyphen:                     {},
	LigatureNameGreaterGreater:                           {},
	LigatureNameGreaterHyphenGreater:                     {},
	LigatureNameGreaterHyphen:                            {},
	LigatureNameHyphenTilde:                              {},
	LigatureNameHyphenBar:                                {},
	LigatureNameHyphenGreaterGreater:                     {},
	LigatureNameHyphenGreater:                            {},
	LigatureNameHyphenHyphenGreater:                      {},
	LigatureNameHyphenHyphenHyphen:                       {},
	LigatureNameHyphenHyphen:                             {},
	LigatureNameHyphenLessLess:                           {},
	LigatureNameHyphenLess:                               {},
	LigatureNameLessTildeTilde:                           {},
	LigatureNameLessTildeGreater:                         {},
	LigatureNameLessTilde:                                {},
	LigatureNameLessAsteriskGreater:                      {},
	LigatureNameLessAsterisk:                             {},
	LigatureNameLessBarBarBar:                            {},
	LigatureNameLessBarBar:                               {},
	LigatureNameLessBarGreater:                           {},
	LigatureNameLessBar:                                  {},
	LigatureNameLessColon:                                {},
	LigatureNameLessDollarGreater:                        {},
	LigatureNameLessDollar:                               {},
	LigatureNameLessEqualBar:                             {},
	LigatureNameLessEqualEqualGreater:                    {},
	LigatureNameLessEqualEqual:                           {},
	LigatureNameLessEqualGreater:                         {},
	LigatureNameLessEqualLess:                            {},
	LigatureNameLessEqual:                                {},
	LigatureNameLessExclamationHyphenHyphen:              {},
	LigatureNameLessGreater:                              {},
	LigatureNameLessHyphenBar:                            {},
	LigatureNameLessHyphenGreater:                        {},
	LigatureNameLessHyphenHyphen:                         {},
	LigatureNameLessHyphenLess:                           {},
	LigatureNameLessHyphen:                               {},
	LigatureNameLessLessEqual:                            {},
	LigatureNameLessLessHyphenGreaterGreater:             {},
	LigatureNameLessLessHyphen:                           {},
	LigatureNameLessLessLess:                             {},
	LigatureNameLessLess:                                 {},
	LigatureNameLessPlusGreater:                          {},
	LigatureNameLessPlus:                                 {},
	LigatureNameLessSlashGreater:                         {},
	LigatureNameLessSlash:                                {},
	LigatureNameNumberSignBraceLeft:                      {},
	LigatureNameNumberSignBracketLeft:                    {},
	LigatureNameNumberSignColon:                          {},
	LigatureNameNumberSignEqual:                          {},
	LigatureNameNumberSignExclamation:                    {},
	LigatureNameNumberSignNumberSignNumberSignNumberSign: {},
	LigatureNameNumberSignNumberSignNumberSign:           {},
	LigatureNameNumberSignNumberSign:                     {},
	LigatureNameNumberSignParenthesisLeft:                {},
	LigatureNameNumberSignQuestion:                       {},
	LigatureNameNumberSignUnderscoreParenthesisLeft:      {},
	LigatureNameNumberSignUnderscore:                     {},
	LigatureNamePercentPercent:                           {},
	LigatureNamePeriodEqual:                              {},
	LigatureNamePeriodHyphen:                             {},
	LigatureNamePeriodPeriodEqual:                        {},
	LigatureNamePeriodPeriodLess:                         {},
	LigatureNamePeriodPeriodPeriod:                       {},
	LigatureNamePeriodPeriod:                             {},
	LigatureNamePeriodQuestion:                           {},
	LigatureNamePlusGreater:                              {},
	LigatureNamePlusPlusPlus:                             {},
	LigatureNamePlusPlus:                                 {},
	LigatureNameQuestionColon:                            {},
	LigatureNameQuestionEqual:                            {},
	LigatureNameQuestionPeriod:                           {},
	LigatureNameQuestionQuestion:                         {},
	LigatureNameSemicolonSemicolon:                       {},
	LigatureNameSlashAsterisk:                            {},
	LigatureNameSlashBackslash:                           {},
	LigatureNameSlashEqualEqual:                          {},
	LigatureNameSlashEqual:                               {},
	LigatureNameSlashGreater:                             {},
	LigatureNameSlashSlashSlash:                          {},
	LigatureNameSlashSlash:                               {},
	LigatureNameUnderscoreBarUnderscore:                  {},
	LigatureNameUnderscoreUnderscore:                     {},
	LigatureNameLowerWWW:                                 {},
}

// isLigatureSupported checks if ligature is supported.
func isLigatureSupported(name ligaturizer.LigatureName) bool {
	_, ok := availableLigatureNames[name]

	return ok
}
