package ui

import (
	"fmt"

	"engo.io/ecs"
	"engo.io/engo/common"
)

// Button is a clickable Element, which performs an action when clicked
type Button struct {
	*ecs.BasicEntity
	*common.MouseComponent
	*common.SpaceComponent
	*common.RenderComponent
	*ActionComponent
}

// ActionComponent contains an action to be executed
type ActionComponent struct {
	Action func()
}

type actionEntity struct {
	*ecs.BasicEntity
	*common.MouseComponent
	*ActionComponent
}

// ActionSystem executes the action stored in the ActionComponent
type ActionSystem struct {
	entities []actionEntity
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (a *ActionSystem) Update(float32) {
	for _, e := range a.entities {
		if e.Clicked {
			e.Action()
		}
	}
}

// Remove is called whenever an Entity is removed from the World, in order to remove it from this sytem as well
func (a *ActionSystem) Remove(ecs.BasicEntity) {}

// New initializes the system
func (a *ActionSystem) New(*ecs.World) {
	fmt.Println("ActionSystem was created")
}

// Add adds a new Button to the system
func (a *ActionSystem) Add(basic *ecs.BasicEntity, action *ActionComponent, mouse *common.MouseComponent) {
	a.entities = append(a.entities, actionEntity{basic, mouse, action})
}
