package ui

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

// Menu is a collection of UIElements arranged to build a menu
// It only sets the values of the elements
type Menu struct {
	*common.SpaceComponent
	buttons []buttonEntity
}

// Build arranges the elements relative to the location of the menu
// Currently only vertical arrangement is supported
// The SpaceComponent of the elements will be changed
func (m *Menu) positionEntries() {
	amount := len(m.buttons)
	fmt.Println(m)
	part := (m.Height) / float32(amount)

	for i := 0; i < amount; i++ {
		sc := &common.SpaceComponent{
			Width:  m.Width,
			Height: m.Height / float32(amount),
			Position: engo.Point{
				X: m.Position.X,
				Y: m.Position.Y + float32(i)*part,
			},
		}
		m.buttons[i].SpaceComponent = sc
	}
}

func (m *Menu) createButtons(buttons []Button) {
	fnt := &common.Font{
		URL:  "Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 64,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		fmt.Println(err)
	}
	for _, button := range buttons {
		basic := ecs.NewBasic()
		e := buttonEntity{}
		e.BasicEntity = &basic
		e.MouseComponent = &common.MouseComponent{}
		e.parts = make([]*buttonPart, 2)
		e.ActionComponent = &ActionComponent{}
		e.Action = button.Action

		borderPart := buttonPart{}
		borderPart.RenderComponent = &common.RenderComponent{
			Drawable: common.Rectangle{
				BorderWidth: 1.,
				BorderColor: color.RGBA{0, 255, 0, 255},
			},
			Hidden: false,
		}
		e.parts[0] = &borderPart

		textPart := buttonPart{}
		textPart.RenderComponent = &common.RenderComponent{
			Drawable: common.Text{
				Font: fnt,
				Text: button.Text,
			},
		}
		e.parts[1] = &textPart
		m.buttons = append(m.buttons, e)
	}

}

type Button struct {
	Text   string
	Action func()
}

type buttonEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*common.MouseComponent
	*ActionComponent
	parts []*buttonPart
}

type buttonPart struct {
	*common.RenderComponent
}

//BuildMenu creates the menu from the given buttons
func (m *Menu) BuildMenu(buttons []Button) {
	m.buttons = make([]buttonEntity, 0)
	m.createButtons(buttons)
	m.positionEntries()
	return
}

// AddTo adds the menu elements to the given system
func (m *Menu) AddTo(system ecs.System) {
	switch sys := system.(type) {
	case *common.RenderSystem:
		for _, ele := range m.buttons {
			sys.Add(ele.BasicEntity, ele.parts[0].RenderComponent, ele.SpaceComponent)
			sys.Add(ele.BasicEntity, ele.parts[1].RenderComponent, ele.SpaceComponent)
		}
	case *common.MouseSystem:
		for _, ele := range m.buttons {
			sys.Add(ele.BasicEntity, ele.MouseComponent, ele.SpaceComponent, ele.parts[1].RenderComponent)
		}
	case *ActionSystem:
		for _, ele := range m.buttons {
			sys.Add(ele.BasicEntity, ele.ActionComponent, ele.MouseComponent)
		}
	}
}
