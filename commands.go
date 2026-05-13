package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/KIMovchanin/go_hotkeys/internal/config"
	"github.com/KIMovchanin/go_hotkeys/internal/hotkeys"
)

func addBind() {
	// Ожидаем команду такого вида:
	// go run . add "Name" "Ctrl+Alt+R" "target"
	if len(os.Args) != 5 {
		fmt.Println(`Usage: go run . add "Name" "Ctrl+Alt+R" "Target"`)
		return
	}

	// Создаём из полученных аргументов горячую клавишу
	newBind := config.Bind{
		Name:   os.Args[2],
		Hotkey: os.Args[3],
		Target: os.Args[4],
	}

	if _, _, err := hotkeys.ParseHotKey(newBind.Hotkey); err != nil {
		log.Fatal("Invalid hotkey:", err)
	}

	binds, err := config.Load(configPath)
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	for _, bind := range binds {
		if strings.EqualFold(bind.Name, newBind.Name) {
			log.Fatal("Bind with this name already exists:", newBind.Name)
		}

		sameHotkey, err := hotkeys.SameHotKey(bind.Hotkey, newBind.Hotkey)
		if err != nil {
			log.Fatal("Failed to compare hotkeys:", err)
		}

		if sameHotkey {
			log.Fatal("Bind with this hotkey already exists:", newBind.Hotkey)
		}
	}

	binds = append(binds, newBind)

	if err := config.Save(configPath, binds); err != nil {
		log.Fatal("Failed to save config:", err)
	}

	fmt.Println("Added: ", newBind.Hotkey, "->", newBind.Name)
}

func seeBinds() {
	if len(os.Args) > 2 {
		fmt.Println("You do not need to write more arguments then just 'see'")
	}
	binds, err := config.Load(configPath)
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	if len(binds) == 0 {
		fmt.Println("No binds configured")
		return
	}

	for _, bind := range binds {
		fmt.Println(bind.Hotkey, "->", bind.Name, "->", bind.Target)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(`  go run .`)
	fmt.Println(`  go run . add "Name" "Ctrl+Alt+R" "Target"`)
	fmt.Println(`  go run . see`)
	fmt.Println(`  go run . delete "Name"`)
	fmt.Println(`  go run . delete_all`)
}

func deleteBind() {
	if len(os.Args) < 3 {
		fmt.Println(`Usage: go run . delete "Name" ["Another name"]`)
		return
	}

	namesToDelete := os.Args[2:]

	binds, err := config.Load(configPath)
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	filtered := make([]config.Bind, 0, len(binds))

	var deletedCount int = 0
	detected := false

	for _, bind := range binds {
		detected = false
		for _, ntd := range namesToDelete {
			if strings.EqualFold(strings.TrimSpace(bind.Name), strings.TrimSpace(ntd)) {
				detected = true
				deletedCount++
				break
			}
		}
		if detected {
			continue
		}
		filtered = append(filtered, bind)
	}

	if deletedCount <= 0 {
		log.Fatal("Bind not found:", namesToDelete)
	}

	if err := config.Save(configPath, filtered); err != nil {
		log.Fatal("Failed to save config:", err)
	}

	fmt.Println("Deleted:", namesToDelete)
}

func deleteAll() {
	if err := config.Save(configPath, []config.Bind{}); err != nil {
		log.Fatal("Failed to save config:", err)
	}

	fmt.Println("All hotkeys was removed!")
}
