package config

import (
	"encoding/json"
	"os"
)

type Bind struct {
	Name   string `json:"name"`
	Hotkey string `json:"hotkey"`
	Target string `json:"target"`
}

func Load(path string) ([]Bind, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var binds []Bind

	err = json.Unmarshal(data, &binds)
	if err != nil {
		return nil, err
	}

	return binds, nil
}
