package handler

import "gopkg.in/yaml.v2"

type YAMLResolver struct {
}

func (rv YAMLResolver) Tag(model interface{}) (err error) {
	return
}

func (rv YAMLResolver) Resolve(data []byte) (resolvedData map[interface{}]interface{}, err error) {
	dataPack := make(map[interface{}]interface{})
	err = yaml.Unmarshal(data, &dataPack)
	if err != nil {
		return nil, err
	}
	return dataPack, nil
}
