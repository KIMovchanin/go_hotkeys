package config

import (
	"path/filepath"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	// t.TempDir создаёт временную папку для теста.
	// После завершения теста Go её удалит
	dir := t.TempDir()

	// filepath.Join безопасно собирает путь под текущую ОС.
	path := filepath.Join(dir, "config.json")

	// Данные которые мы хотим сохранить в JSON.
	want := []Bind{
		{
			Name:   "YouTube",
			Hotkey: "Ctrl+Alt+Y",
			Target: "https://www.youtube.com",
		},
		{
			Name:   "Notepad",
			Hotkey: "Ctrl+Alt+N",
			Target: "notepad.exe",
		},
	}

	// передаю в Save собарнный путь из временной папки и названия файла
	// и затем то, что я хочу сохранить (оно будет превращено в byte и затем записано на диск).
	err := Save(path, want)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Читаем ранее созданный файл
	got, err := Load(path)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d binds, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("bind %d: expected %+v, got %+v", i, want[i], got[i])
		}
	}

}
