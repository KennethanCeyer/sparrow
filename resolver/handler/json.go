package handler

import (
	"encoding/json"
	"github.com/KennethanCeyer/tit/utils"
)

type JSONResolver struct {
}

func (rv JSONResolver) Resolve(data []byte) (resolvedData map[interface{}]interface{}, err error) {
	var dataPack interface{}
	err = json.Unmarshal(data, &dataPack)
	if err != nil {
		return nil, err
	}
	return utils.ConvertMapForm(dataPack.(map[string]interface{})), nil
}
