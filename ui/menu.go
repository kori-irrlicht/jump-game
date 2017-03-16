package ui

import (
	"fmt"
	"image/color"

	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
)

var (
	buttonFnt = &common.Font{
		URL:  "Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 20,
	}
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
	part := (m.Height) / float32(amount)

	for i := 0; i < amount; i++ {
		scBorder := &common.SpaceComponent{
			Width:  m.Width,
			Height: m.Height / float32(amount),
			Position: engo.Point{
				X: m.Position.X,
				Y: m.Position.Y + float32(i)*part,
			},
		}
		m.buttons[i].SpaceComponent = scBorder
		m.buttons[i].parts[0].SpaceComponent = scBorder

		text := m.buttons[i].parts[1].RenderComponent.Drawable.(common.Text)
		dimX, dimY, _ := buttonFnt.TextDimensions(text.Text)
		scText := &common.SpaceComponent{
			Width:  float32(dimX),
			Height: float32(dimY),
		}
		scText.Center(engo.Point{
			X: scBorder.Position.X + (scBorder.Width / 2),
			Y: scBorder.Position.Y + (scBorder.Height / 2),
		})
		m.buttons[i].parts[1].SpaceComponent = scText
	}
}

func (m *Menu) createButtons(buttons []Button) {
	err := buttonFnt.CreatePreloaded()
	if err != nil {
		fmt.Println(err)
	}
	for _, button := range buttons {
		basics := ecs.NewBasics(3)
		e := buttonEntity{}
		e.BasicEntity = &basics[0]
		e.MouseComponent = &common.MouseComponent{}
		e.parts = make([]*buttonPart, 2)
		e.ActionComponent = &ActionComponent{}
		e.Action = button.Action

		borderPart := buttonPart{}
		borderPart.BasicEntity = &basics[1]
		borderPart.RenderComponent = &common.RenderComponent{
			Drawable: common.Rectangle{
				BorderWidth: 1.,
				BorderColor: color.RGBA{0, 255, 0, 255},
			},
			Hidden: false,
		}
		borderPart.SetZIndex(5)
		e.parts[0] = &borderPart

		textPart := buttonPart{}
		textPart.BasicEntity = &basics[2]
		textPart.RenderComponent = &common.RenderComponent{
			Drawable: common.Text{
				Font: buttonFnt,
				Text: button.Text,
			},
			Hidden: false,
		}
		e.parts[1] = &textPart
		textPart.SetZIndex(6)
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
	*ecs.BasicEntity
	*common.RenderComponent
	*common.SpaceComponent
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
			sys.Add(ele.parts[0].BasicEntity, ele.parts[0].RenderComponent, ele.parts[0].SpaceComponent)
			sys.Add(ele.parts[1].BasicEntity, ele.parts[1].RenderComponent, ele.parts[1].SpaceComponent)
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
