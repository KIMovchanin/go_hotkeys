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
	parts := strings.Split(text, "+") // разбили строку по +.
	if len(parts) < 2 {
		return nil, 0, fmt.Errorf("hotkey must contain modifiers and key: %q", text)
	}

	// инициализируем слайс (срез) для модификаторов.
	// он в длину меньше на 1 чем длина всех частей тк не включает
	// в себя hotkey.Key (последнйи символ в горячей клавише).
	mods := make([]hotkey.Modifier, 0, len(parts)-1)

	for i, part := range parts {
		// типизируем все части под один формат.
		token := strings.ToLower(strings.TrimSpace(part))
		if token == "" {
			return nil, 0, fmt.Errorf("empty hotkey, part in %q", text)
		}

		// определяем последний ли это эллемент из parts
		// благодаря сравнению индекса с длиной parts
		// (-1 т.к. инедксы с 0, а длина с 1).
		isLastPart := i == len(parts)-1
		if isLastPart {
			// если такой ключ есть в мапе.
			key, ok := keyByName[token]
			if !ok {
				return nil, 0, fmt.Errorf("unknown key %q in hotkey %q", token, text)
			}

			return mods, key, nil
		}

		// если такой модификатор есть в мапе.
		mod, ok := modifierByName[token]
		if !ok {
			return nil, 0, fmt.Errorf("unknown modifier %q in token %q", token, text)
		}

		mods = append(mods, mod)
	}

	return nil, 0, fmt.Errorf("invalid hotkey: %q", text)
}
