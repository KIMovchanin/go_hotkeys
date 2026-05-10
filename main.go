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

	"golang.design/x/hotkey/mainthread"
)

func main() {
	// mainthread.Init нужен, чтобы библиотека hotkey нормально работала
	// с системным event loop. В документации он указан как способ запускать
	// self-contained приложения с hotkey.
	mainthread.Init(run)
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
