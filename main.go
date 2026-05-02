package main

import (
	"fmt"
	"log"

	"github.com/KIMovchanin/go_hotkeys/internal/launcher"
)

func main() {
	fmt.Println("HotLaunch started")

	err := launcher.Launch("D:/Steam/steam.exe")
	if err != nil {
		log.Println("Ошибка запуска:", err)
	}
}
