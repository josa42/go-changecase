package changecase

import (
	"regexp"
	"strings"
)

var (
	exprWordFirst       = regexp.MustCompile(`^.| .`)
	exprCammelCase      = regexp.MustCompile(`([a-z])([A-Z])`)
	exprTraingNoWord    = regexp.MustCompile(`(^[^A-Za-z]|[^A-Za-z]$)`)
	exprNoWord          = regexp.MustCompile(`[-_.()\[\]\s]+`)
	exprChar            = regexp.MustCompile(`.`)
	exprWordFirstMiddle = regexp.MustCompile(` .`)
)

// ToLower converts case to lower (abcdef).
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper converts case to upper (ABCDEF).
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToUpperFirst converts first character to upper case.
func ToUpperFirst(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

// ToLowerFirst converts case first character to upper case.
func ToLowerFirst(s string) string {
	if len(s) > 0 {
		return strings.ToLower(s[:1]) + s[1:]
	}
	return s
}

// IsUpper determinea whether a string is all upper case or not
func IsUpper(s string) bool {
	return ToUpper(s) == s
}

// IsLower determinea whether a string is all lower case or not
func IsLower(s string) bool {
	return ToLower(s) == s
}

// ToTitle converts	case to title case (Abc Def).
func ToTitle(s string) string {
	s = ToNo(s)

	s = exprWordFirst.ReplaceAllStringFunc(s, func(s string) string {
		return ToUpper(s)
	})

	return s
}

// ToNo converts case to no case (abc def).
func ToNo(s string) string {
	// https://github.com/blakeembrey/no-case/blob/master/no-case.js
	s = exprCammelCase.ReplaceAllString(s, `$1 $2`)
	s = exprTraingNoWord.ReplaceAllString(s, "")
	s = exprNoWord.ReplaceAllString(s, " ")

	return ToLower(s)
}

// ToHeader converts case to header case (Abc-Def).
func ToHeader(s string) string {
	s = ToNo(s)
	s = exprWordFirst.ReplaceAllStringFunc(s, func(s string) string {
		return ToUpper(s)
	})
	s = strings.Replace(s, " ", "-", -1)
	return s
}

// ToSentence converts case to sentence case (Abc def).
func ToSentence(s string) string {
	s = ToNo(s)
	s = ToUpperFirst(s)
	return s
}

// ToCamel converts case to camel case (abcDef).
func ToCamel(s string) string {
	s = ToNo(s)
	s = exprWordFirstMiddle.ReplaceAllStringFunc(s, func(g string) string {
		return ToUpper(g[1:])
	})
	return s
}

// ToPascal converts case to snake case (AbcDef).
func ToPascal(s string) string {
	s = ToCamel(s)
	s = ToUpperFirst(s)
	return s
}

// ToSnake converts case to snake case (abc_def),
func ToSnake(s string) string {
	s = ToNo(s)
	s = strings.Replace(s, " ", "_", -1)
	return s
}

// ToParam converts case to param (akka kebab or hyphen) case (abc-def)
func ToParam(s string) string {
	s = ToNo(s)
	s = strings.Replace(s, " ", "-", -1)
	return s
}

// ToConstant converts case to constant (ABC_DEF).
func ToConstant(s string) string {
	s = ToSnake(s)
	s = ToUpper(s)
	return s
}

// ToDot converts case to dot (abc.def)
func ToDot(s string) string {
	s = ToNo(s)
	s = strings.Replace(s, " ", ".", -1)
	return s
}

// ToPath converts case to path (abc/def).
func ToPath(s string) string {
	s = ToNo(s)
	s = strings.Replace(s, " ", "/", -1)
	return s
}

// Swap case (aBcDeF -> AbCdEf)
func Swap(s string) string {
	return exprChar.ReplaceAllStringFunc(s, func(c string) string {
		if IsUpper(c) {
			return ToLower(c)
		}
		return ToUpper(c)
	})
}
