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
	"win":     hotkey.ModWin,
	"windows": hotkey.ModWin,
}

// будет связывать текст, но уже именно Key, а не Modifier
var keyByName = map[string]hotkey.Key{
	"a": hotkey.KeyA,
	"b": hotkey.KeyB,
	"c": hotkey.KeyC,
	"d": hotkey.KeyD,
	"e": hotkey.KeyE,
	"f": hotkey.KeyF,
	"g": hotkey.KeyG,
	"h": hotkey.KeyH,
	"i": hotkey.KeyI,
	"j": hotkey.KeyJ,
	"k": hotkey.KeyK,
	"l": hotkey.KeyL,
	"m": hotkey.KeyM,
	"n": hotkey.KeyN,
	"o": hotkey.KeyO,
	"p": hotkey.KeyP,
	"q": hotkey.KeyQ,
	"r": hotkey.KeyR,
	"s": hotkey.KeyS,
	"t": hotkey.KeyT,
	"u": hotkey.KeyU,
	"v": hotkey.KeyV,
	"w": hotkey.KeyW,
	"x": hotkey.KeyX,
	"y": hotkey.KeyY,
	"z": hotkey.KeyZ,

	"space":  hotkey.KeySpace,
	"esc":    hotkey.KeyEscape,
	"escape": hotkey.KeyEscape,
	"tab":    hotkey.KeyTab,

	"0": hotkey.Key0,
	"1": hotkey.Key1,
	"2": hotkey.Key2,
	"3": hotkey.Key3,
	"4": hotkey.Key4,
	"5": hotkey.Key5,
	"6": hotkey.Key6,
	"7": hotkey.Key7,
	"8": hotkey.Key8,
	"9": hotkey.Key9,
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

	// создаём проверку на дубликаты для защиты от странных данных.
	seenMods := make(map[hotkey.Modifier]bool)

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

		if seenMods[mod] {
			return nil, 0, fmt.Errorf("dublicate modifier %q in hotkey %q", token, text)
		}

		seenMods[mod] = true
		mods = append(mods, mod)
	}

	return nil, 0, fmt.Errorf("invalid hotkey: %q", text)
}
