package tibivi

import (
	"github.com/oltarzewskik/tibivi-gocui"
)

// quitIgnore exits tibivi without saving current state
func (tbv *Tibivi) quitIgnore(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
