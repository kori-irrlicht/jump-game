package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

const buttonOpenMenu = "OpenMenu"

type game struct{}

func (g *game) Type() string { return sceneGame }
func (g *game) Preload()     {}
func (g *game) Setup(world *ecs.World) {

	engo.Input.RegisterButton(buttonOpenMenu, engo.Escape)

	common.SetBackground(color.White)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&inputSystem{})
}

type inputSystem struct{}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (is *inputSystem) Update(float32) {
	if engo.Input.Button(buttonOpenMenu).JustPressed() {
		engo.SetSceneByName(sceneMainMenu, true)
	}
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (is *inputSystem) Remove(ecs.BasicEntity) {}
