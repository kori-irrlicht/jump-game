package ui

import (
	"engo.io/engo"
	"engo.io/engo/common"
)

//  Element is displayed in the UI
type Element interface {
	AddSpaceComponent(*common.SpaceComponent)
}

// Menu is a collection of UIElements arranged to build a menu
// It only sets the values of the elements
type Menu struct {
	*common.SpaceComponent
	elements []Element
}

//Add adds an Element to the menu
func (m *Menu) Add(e Element) {
	m.elements = append(m.elements, e)
}

// Build arranges the elements relative to the location of the menu
// Currently only vertical arrangement is supported
// The SpaceComponent of the elements will be changed
func (m *Menu) Build() {
	amount := len(m.elements)
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
		m.elements[i].AddSpaceComponent(sc)
	}

}

type Button interface {
	Text() string
	Action()
}

func BuildMenu(buttons []Button) (menu Menu) {
	return
}
