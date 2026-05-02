package hotkeys

import (
	"golang.design/x/hotkey"
)

type Manager struct {
	keys []*hotkey.Hotkey
}

func NewManager() *Manager {
	return &Manager{keys: make([]*hotkey.Hotkey, 0)}
}