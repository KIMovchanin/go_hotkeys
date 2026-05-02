package main

import (
	"fmt"
	"log"

	"github.com/KIMovchanin/go_hotkeys/internal/launcher"
)

func main() {
	fmt.Println("HotLaunch started")

	err := launcher.Launch("https://youtube.com")
	if err != nil {
		log.Println("Ошибка запуска:", err)
	}
}
