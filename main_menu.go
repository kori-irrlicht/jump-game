package main

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"github.com/kori-irrlicht/jump-game/ui"
)

type mainMenu struct{}

func (g *mainMenu) Type() string { return sceneMainMenu }
func (g *mainMenu) Preload() {
	err := engo.Files.Load("Roboto-Regular.ttf")
	fmt.Println(err)
}
func (g *mainMenu) Setup(world *ecs.World) {

	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&ui.ActionSystem{})

	menu := buildMainMenu()

	for _, s := range world.Systems() {
		menu.AddTo(s)
	}

}

func buildMainMenu() *ui.Menu {
	buttons := []ui.Button{
		{
			Text: "Hallo",
			Action: func() {
				fmt.Println("test")
			},
		},
	}
	m := &ui.Menu{}
	m.SpaceComponent = &common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Height:   100,
		Width:    200,
	}
	m.Center(engo.Point{
		X: 400,
		Y: 320,
	})

	m.BuildMenu(buttons)
	return m

}
