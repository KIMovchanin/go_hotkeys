package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/KIMovchanin/go_hotkeys/internal/config"
	"github.com/KIMovchanin/go_hotkeys/internal/hotkeys"
	"github.com/KIMovchanin/go_hotkeys/internal/launcher"

	"golang.design/x/hotkey/mainthread"
)

func main() {
	// Если длина предаваемых через консоль аргументов больше 1
	// и этот аргумет "add", то запустить addBind().
	// if len(os.Args) > 1 {
	// 	switch os.Args[1] {
	// 	case "add":
	// 		addBind()
	// 	case "delete":
	// 		deleteBind()
	// 	case "del":
	// 		deleteBind()
	// 	case "delete_all":
	// 		deleteAllBinds()
	// 	case "del_all":
	// 		deleteAllBinds()
	// 	case "see":
	// 		seeBinds()
	// 	default:
	// 		printUsage()
	// 	}
	// 	return
	// }

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			addBind()
		case "see":
			seeBinds()
		default:
			printUsage()
		}
		return

	}

	// mainthread.Init нужен, чтобы библиотека hotkey нормально работала
	// с системным event loop. В документации он указан как способ запускать
	// self-contained приложения с hotkey.
	mainthread.Init(run)
}

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

	binds, err := config.Load("config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	for _, bind := range binds {
		if strings.EqualFold(bind.Name, newBind.Name) {
			log.Fatal("Bind with this name already exists:", newBind.Name)
		}

		if strings.EqualFold(bind.Hotkey, newBind.Hotkey) {
			log.Fatal("Bind with this hotkey already exists:", newBind.Hotkey)
		}
	}

	binds = append(binds, newBind)

	if err := config.Save("config.json", binds); err != nil {
		log.Fatal("Failed to save config:", err)
	}

	fmt.Println("Added: ", newBind.Hotkey, "->", newBind.Name)
}

func seeBinds() {
	if len(os.Args) > 2 {
		fmt.Println("You do not need to write more arguments then just 'see'")
	}
	binds, err := config.Load("config.json")
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
}

func run() {
	fmt.Println("Hotkeys app started.")

	manager := hotkeys.NewManager()
	defer manager.UnregisterAll() // активируется при завершении функции run, что удаляет все хоткеи

	configBinds, err := config.Load("config.json")
	if err != nil {
		log.Fatal("Error of read json:", err)
	}

	for _, bind := range configBinds {
		currentBind := bind
		mods, key, err := hotkeys.ParseHotKey(currentBind.Hotkey)
		if err != nil {
			log.Fatal("Error of parsing:", err)
		}

		err = manager.Register(mods, key, func() {
			err := launcher.Launch(currentBind.Target)
			if err != nil {
				log.Println("Error of launch:", err, currentBind.Name)
			}
		})
		if err != nil {
			log.Fatal("Error of register:", err, currentBind.Name)
		}

		fmt.Println("Registered:", currentBind.Hotkey, "->", currentBind.Name)
	}

	fmt.Println("Press Ctrl + C to exit")

	waitForExit()
}

func waitForExit() { // не понимаю полностью это
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // требование ждать подачи сигнала
	fmt.Println("Exiting...")
}
