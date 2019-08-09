package utils

import (
	"unicode"
)

type caseType uint8

const (
	pascalCase caseType = 0
	camelCase caseType = 1
)

func asCase(s string, t caseType) string {
	b := make([]byte, len(s))
	cursor := 0

	headPrefixHandler := func(x rune) rune { return x }
	switch t {
	case pascalCase:
		headPrefixHandler = unicode.ToUpper
	case camelCase:
		headPrefixHandler = unicode.ToLower
	}

	for i:=0; i<len(s); i, cursor = i+1, cursor+1 {
		if i == 0 {
			b[i] = byte(headPrefixHandler(rune(s[i])))
			continue
		}

		if s[i] == ' ' || s[i] == '-' || s[i] == '_' {
			b[cursor] = byte(unicode.ToUpper(rune(s[i + 1])))
			i++
			continue
		}
		b[cursor] = s[i]
	}
	return string(b[:cursor])
}

func ToPascalCase(s string) string {
	return asCase(s, pascalCase)
}

func ToCamelCase(s string) string {
	return asCase(s, camelCase)
}
