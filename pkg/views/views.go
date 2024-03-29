package views

import (
	"github.com/strang1ato/tibivi-gocui"
)

// Views is struct of tibivi views
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
