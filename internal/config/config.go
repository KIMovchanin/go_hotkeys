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
	data, err := os.ReadFile(path) // вернёт слайс байтов
	if err != nil {
		return nil, err
	}

	var binds []Bind

	err = json.Unmarshal(data, &binds) // те байты превратит в строки
	if err != nil {
		return nil, err
	}

	return binds, nil
}
