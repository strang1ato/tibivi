package tibivi

import (
	"github.com/oltarzewskik/gocui"
)

// quit exit tibivi without saving
func (tbv *Tibivi) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
