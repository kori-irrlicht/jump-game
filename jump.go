package main

import (
	"engo.io/ecs"
	"engo.io/engo"
)

const (
	sceneGame     = "scene_game"
	sceneMainMenu = "scene_main_menu"
)

type game struct{}

func (g *game) Type() string     { return sceneGame }
func (g *game) Preload()         {}
func (g *game) Setup(*ecs.World) {}

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
