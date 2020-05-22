package views

import (
	"github.com/oltarzewskik/tibivi-gocui"
)

type Views struct {
	Bar   map[string]*gocui.View
	Days  map[string]*gocui.View
	Menu  *gocui.View
	Block map[string]*gocui.View
}

// SetViews returns new Views struct
func SetViews() *Views {
	Views := &Views{
		Bar:   map[string]*gocui.View{},
		Days:  map[string]*gocui.View{},
		Menu:  &gocui.View{},
		Block: map[string]*gocui.View{},
	}
	return Views
}
