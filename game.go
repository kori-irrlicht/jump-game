package main

import (
	"image/color"

	"engo.io/ecs"
	"engo.io/engo/common"
)

type game struct{}

func (g *game) Type() string { return sceneGame }
func (g *game) Preload()     {}
func (g *game) Setup(*ecs.World) {

	common.SetBackground(color.White)
}
