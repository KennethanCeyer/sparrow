package tit

import (
	"fmt"
	"github.com/KennethanCeyer/tit/resolver"
	"github.com/KennethanCeyer/tit/resolver/handler"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
)

const AppName = "tit"

func getResolver(resolverType resolver.Type) (rv *resolver.Resolver, err error) {
	var rvHandler resolver.Resolver
	switch resolver.Type(strings.ToLower(string(resolverType))) {
	case resolver.YAMLType:
		rvHandler = handler.YAMLResolver{}
	case resolver.TOMLType:
		rvHandler = handler.TOMLResolver{}
	case resolver.JSONType:
		rvHandler = handler.JSONResolver{}
	default:
		return nil, fmt.Errorf("`%v` is not supported type of resolver", resolverType)
	}

	return &rvHandler, nil
}

func ReadFile(file string, model interface{}) (err error) {
	if !path.IsAbs(file) {
		_, callerPath, _, _ := runtime.Caller(1)
		file = path.Join(path.Dir(callerPath), file)
	}

	resolverType, err := getTypeFromExt(file)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	rvPointer, err := getResolver(resolverType)
	if err != nil {
		return err
	}

	rv := *rvPointer
	resolvedData, err := rv.Resolve(data)
	if err != nil {
		return err
	}

	propagateToModel(resolvedData, model)
	return
}
