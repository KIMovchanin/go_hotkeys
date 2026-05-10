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

// функция записи новых биндов в json
func Save(path string, binds []Bind) error {
	// json.MarshalIndent превращает Go-структуры в красивые JSON-байты.
	// MarshalIndent удобнее обычного Marshal, потому что файл остаётся читаемым.
	data, err := json.MarshalIndent(binds, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
