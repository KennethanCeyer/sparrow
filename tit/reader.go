package tit

import (
	"path/filepath"
	"strings"
	"tit/resolver"
)

func cleansingAlias(name string) (alias resolver.Type) {
	lowerName := strings.ToLower(name)
	switch lowerName {
	case "yml":
		return resolver.YAMLType
	}
	return resolver.Type(lowerName)
}

func getTypeFromExt(file string) (rvType resolver.Type, err error) {
	ext := strings.Trim(filepath.Ext(file), ".")
	resolveType := cleansingAlias(ext)
	return resolveType, nil
}

