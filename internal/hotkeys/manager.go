package hotkeys

import (
	"golang.design/x/hotkey"
)

// Manager хранит зарегистрированные горячие клавиши.
type Manager struct {
	keys []*hotkey.Hotkey
}

func NewManager() *Manager {
	return &Manager{keys: make([]*hotkey.Hotkey, 0)}
}

func (m *Manager) Register(mods []hotkey.Modifier, key hotkey.Key, action func()) error {
	hk := hotkey.New(mods, key)
	if err := hk.Register(); err != nil {
		return err
	}

	m.keys = append(m.keys, hk)

	go func() { // горутина для обработки нажатий клавиш но как работает не совсем понимаю
		for range hk.Keydown() {
			action()
		}
	}()

	return nil
}

func (m *Manager) UnregisterAll() {
	for _, hk := range m.keys {
		_ = hk.Unregister()
	}
}
