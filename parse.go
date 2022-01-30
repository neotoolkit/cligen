package cligen

import (
	"io/ioutil"

	"github.com/goccy/go-yaml"
)

func Parse(path string) (API, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return API{}, err
	}

	var api API

	if err := yaml.Unmarshal(file, &api); err != nil {
		return API{}, err
	}

	return api, nil
}
