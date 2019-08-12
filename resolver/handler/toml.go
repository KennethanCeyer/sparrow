package handler

import (
	"github.com/BurntSushi/toml"
	"github.com/KennethanCeyer/tit/utils"
)

type TOMLResolver struct {
}

func (rv TOMLResolver) Resolve(data []byte) (resolvedData map[interface{}]interface{}, err error) {
	var dataPack interface{}
	err = toml.Unmarshal(data, &dataPack)
	if err != nil {
		return nil, err
	}
	return utils.ConvertMapForm(dataPack.(map[string]interface{})), nil
}
