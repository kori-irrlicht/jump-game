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

	startButton := newButton(
		func() {
			fmt.Println("Start new game")
			engo.SetSceneByName(sceneGame, false)
		})
	b1 := newButton(func() {
		fmt.Println("b1")
	})
	b2 := newButton(func() {
		fmt.Println("b2")
	})

	buttons := []*ui.Button{startButton, b1, b2}

	buildMainMenu(buttons)

	for _, ele := range buttons {
		for _, s := range world.Systems() {
			switch sys := s.(type) {
			case *common.RenderSystem:
				sys.Add(ele.BasicEntity, ele.RenderComponent, ele.SpaceComponent)
			case *common.MouseSystem:
				sys.Add(ele.BasicEntity, ele.MouseComponent, ele.SpaceComponent, ele.RenderComponent)
			case *ui.ActionSystem:
				sys.Add(ele.BasicEntity, ele.ActionComponent, ele.MouseComponent)
			}
		}

	}

}

func buildMainMenu(buttons []*ui.Button) {
	m := ui.Menu{}
	m.SpaceComponent = &common.SpaceComponent{
		Position: engo.Point{X: 0, Y: 0},
		Height:   100,
		Width:    200,
	}
	m.Center(engo.Point{
		X: 400,
		Y: 320,
	})
	for _, b := range buttons {
		m.Add(b)
	}
	m.Build()

}

func newButton(action func()) *ui.Button {
	e := ecs.NewBasic()
	startButton := ui.Button{BasicEntity: &e}
	startButton.RenderComponent = &common.RenderComponent{
		Drawable: common.Rectangle{
			BorderWidth: 1.,
			BorderColor: color.RGBA{0, 255, 0, 255},
		},
		Hidden: false,
	}
	startButton.MouseComponent = &common.MouseComponent{}

	startButton.ActionComponent = &ui.ActionComponent{}
	startButton.Action = action

	return &startButton

}
