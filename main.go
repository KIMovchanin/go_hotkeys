package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/KIMovchanin/go_hotkeys/internal/hotkeys"
	"github.com/KIMovchanin/go_hotkeys/internal/launcher"

	"golang.design/x/hotkey"
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

	err := manager.Register([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModAlt}, hotkey.KeyY,
		func() {
			fmt.Println("Opening youtube...")

			err := launcher.Launch("https://www.youtube.com")
			if err != nil {
				log.Println("Failed to launch youtube:", err)
			}
		},
	)

	if err != nil {
		log.Fatal("Failed to register hotkey:", err)
	}

	fmt.Println("Ctrl + Alt + Y registered")
	fmt.Println("Press Ctrl + C to exit")

	waitForExit()
}

func waitForExit() { // не понимаю полностью это
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // требование ждать подачи сигнала
	fmt.Println("Exiting...")
}
