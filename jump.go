package main

import (
	"engo.io/engo"
)

const (
	sceneGame     = "scene_game"
	sceneMainMenu = "scene_main_menu"
)

func main() {
	opts := engo.RunOptions{
		Title:         "Jump game",
		Width:         800,
		Height:        640,
		ScaleOnResize: true,
		VSync:         true,
		FPSLimit:      60,
	}

	engo.RegisterScene(&game{})

	engo.Run(opts, &mainMenu{})
}
