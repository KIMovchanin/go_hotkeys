package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/KIMovchanin/go_hotkeys/internal/config"
	"github.com/KIMovchanin/go_hotkeys/internal/hotkeys"
	"github.com/KIMovchanin/go_hotkeys/internal/launcher"

	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

type Bind struct {
	Name   string
	Mods   []hotkey.Modifier
	Key    hotkey.Key
	Target string
}

func main() {
	// mainthread.Init нужен, чтобы библиотека hotkey нормально работала
	// с системным event loop. В документации он указан как способ запускать
	// self-contained приложения с hotkey.
	mainthread.Init(run)
}

func run() {
	fmt.Println("Hotkeys app started.")

	configBinds, err := config.Load("config.json")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	fmt.Println("Loaded config:")
	for _, bind := range configBinds {
		fmt.Println("-", bind.Name, bind.Hotkey, "->", bind.Target)
	}

	manager := hotkeys.NewManager()
	defer manager.UnregisterAll() // активируется при завершении функции run, что удаляет все хоткеи

	binds := []Bind{
		{
			Name:   "YouTube",
			Mods:   []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt},
			Key:    hotkey.KeyY,
			Target: "https://www.youtube.com",
		},
		{
			Name:   "GitHub",
			Mods:   []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt},
			Key:    hotkey.KeyG,
			Target: "https://www.github.com/KIMovchanin",
		},
		{
			Name:   "Notepad",
			Mods:   []hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt},
			Key:    hotkey.KeyN,
			Target: "notepad.exe",
		},
	}

	for _, bind := range binds {
		currentBind := bind

		err := manager.Register(currentBind.Mods, currentBind.Key,
			func() {
				fmt.Println("Opening", currentBind.Name)

				err := launcher.Launch(currentBind.Target)
				if err != nil {
					log.Println("Failed to launch", currentBind.Name+":", err)
				}
			})

		if err != nil {
			log.Fatal("Failed to register hotkey:", err)
		}

		fmt.Println("Registered", currentBind.Name)
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
