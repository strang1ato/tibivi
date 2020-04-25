package tibivi

import (
	"fmt"

	"github.com/oltarzewskik/gocui"
)

// setBar setups "bar" and ":" views
func (tbv *Tibivi) setBar(maxX, maxY int) error {
	if v, err := tbv.g.SetView(":", -1, maxY-2, 1, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Visible = false
		fmt.Fprint(v, ":")
	}

	if v, err := tbv.g.SetView("bar", 0, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Frame = false
	}
	return nil
}
