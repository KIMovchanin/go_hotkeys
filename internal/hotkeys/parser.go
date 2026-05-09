package hotkeys

import (
	"fmt"
	"strings"

	"golang.design/x/hotkey"
)

// будет связывать текст из config.json с типами библиотеки hotkey
var modifierByName = map[string]hotkey.Modifier{
	"ctrl":    hotkey.ModCtrl,
	"control": hotkey.ModCtrl,
	"alt":     hotkey.ModAlt,
	"shift":   hotkey.ModShift,
}

// будет связывать текст, но уже именно Key, а не Modifier
var keyByName = map[string]hotkey.Key{
	"y": hotkey.KeyY,
	"n": hotkey.KeyN,
	"g": hotkey.KeyG,
}

// ParseHotkey будет превращать строку "Ctrl+Alt+Y" в mods+key для Register
func ParseHotKey(text string) ([]hotkey.Modifier, hotkey.Key, error) {
	parts := strings.Split(text, "+")
	if len(parts) < 2 {
		return nil, 0, fmt.Errorf("hotkey must contain modifiers and key: %q", text)
	}

	mods := make([]hotkey.Modifier, 0, len(parts)-1)

	for i, part := range parts {
		token := strings.ToLower(strings.TrimSpace(part))
		if token == "" {
			return nil, 0, fmt.Errorf("empty hotkey, part in %q", text)
		}

		isLastPart := i == len(parts)-1
		if isLastPart {
			key, ok := keyByName[token]
			if !ok {
				return nil, 0, fmt.Errorf("unknown key %q in hotkey %q", token, text)
			}

			return mods, key, nil
		}

		mod, ok := modifierByName[token]
		if !ok {
			return nil, 0, fmt.Errorf("unknown modifier %q in token %q", token, text)
		}

		mods = append(mods, mod)
	}

	return nul, 0, fmt.Errorf("invalid hotkey: %q", text)
}
