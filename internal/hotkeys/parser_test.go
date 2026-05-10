package hotkeys

import (
	"testing"

	"golang.design/x/hotkey"
)

func TestParseHotKey(t *testing.T) {
	// Это table-driven test: список сценариев.
	// Вместо шести отдельных функций мы описываем шесть случаев в таблице.
	tests := []struct {
		name     string
		input    string
		wantMods []hotkey.Modifier
		wantKey  hotkey.Key
		wantErr  bool
	}{ //перечисляю названия, ввод и ожидаемые значения (want...)
		{
			name:     "ctrl alt y",
			input:    "Ctrl+Alt+Y",
			wantMods: []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt},
			wantKey:  hotkey.KeyY,
		},
		{
			name:     "spaces are allowed",
			input:    "Ctrl + Alt + N",
			wantMods: []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt},
			wantKey:  hotkey.KeyN,
		},
		{
			name:     "lowercase is allowed",
			input:    "ctrl+alt+g",
			wantMods: []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt},
			wantKey:  hotkey.KeyG,
		},
		{
			name:    "unknown key",
			input:   "Ctrl+Alt+Unknown",
			wantErr: true,
		},
		{
			name:    "empty part",
			input:   "Ctrl++Y",
			wantErr: true,
		},
		{
			name:    "missing modifier",
			input:   "Y",
			wantErr: true,
		},
		{
			name:     "new letter",
			input:    "ctrl+alt+R",
			wantMods: []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt},
			wantKey:  hotkey.KeyR,
			wantErr:  false, // не нужно, тк false и так значение по умолчанию.
		},
		{
			name:    "duplicate modifier",
			input:   "ctrl+ctrl+Y",
			wantErr: true,
		},
		{
			name:    "two keys",
			input:   "ctrl+esc+Y",
			wantErr: true,
		},
		{
			name:     "new key",
			input:    "win+1",
			wantMods: []hotkey.Modifier{hotkey.ModWin},
			wantKey:  hotkey.Key1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // берём имя теста
			// и запускаем для каждого функцию.
			gotMods, gotKey, err := ParseHotKey(tt.input) // берём ввод.

			if tt.wantErr { // проверяем должна ли быть ошибка.
				if err == nil { // если ошибки нет
					t.Fatalf("expected error, got nil") // выводим err.
				}
				return // останавливаем работу.
			}

			if err != nil { // если получили ошибку, когда её не ждали
				t.Fatalf("expected no error, got %v", err) // выводим err.
			}

			if gotKey != tt.wantKey { // если полученный ключ не равен ожидаемому
				t.Fatalf("expected key %v, got %v", tt.wantKey, gotKey) // err.
			}

			if len(gotMods) != len(tt.wantMods) { // если длина полученных модификаторов не равна ожидаемой
				t.Fatalf("expected %d modifiers, got %d", len(tt.wantMods), len(gotMods)) // err.
			}

			for i := range tt.wantMods { // если модификаторы не равны ожидаемым
				if gotMods[i] != tt.wantMods[i] { // проверяем каждый.
					t.Fatalf("modifier %d: expected %v, got %v", i, tt.wantMods[i], gotMods[i])
				}
			}
		})
	}
}
