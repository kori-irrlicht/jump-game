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
func (g *mainMenu) Preload()     {}
func (g *mainMenu) Setup(world *ecs.World) {

	common.SetBackground(color.White)

	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&ui.ActionSystem{})

	e := ecs.NewBasic()
	startButton := ui.Button{BasicEntity: &e}
	startButton.SpaceComponent = &common.SpaceComponent{
		Position: engo.Point{X: 200, Y: 200},
		Width:    200,
		Height:   100,
	}
	startButton.RenderComponent = &common.RenderComponent{
		Drawable: common.Rectangle{
			BorderWidth: 1.,
			BorderColor: color.RGBA{0, 255, 0, 255},
		},
		Hidden: false,
	}
	startButton.MouseComponent = &common.MouseComponent{}

	startButton.ActionComponent = &ui.ActionComponent{}
	startButton.Action = func() {
		fmt.Println("Start new game")
		engo.SetSceneByName(sceneGame, false)
	}

	for _, s := range world.Systems() {
		switch sys := s.(type) {
		case *common.RenderSystem:
			sys.Add(startButton.BasicEntity, startButton.RenderComponent, startButton.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(startButton.BasicEntity, startButton.MouseComponent, startButton.SpaceComponent, startButton.RenderComponent)
		case *ui.ActionSystem:
			sys.Add(startButton.BasicEntity, startButton.ActionComponent, startButton.MouseComponent)
		}
	}

}
