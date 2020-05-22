package actions

import (
	"github.com/oltarzewskik/tibivi-gocui"
)

// QuitIgnore exits tibivi without saving current state
func QuitIgnore(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
