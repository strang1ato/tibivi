package tibivi

import (
	"github.com/oltarzewskik/gocui"
)

// keybindings create keyboard keybindings
func (tbv *Tibivi) keybindings() error {
	if err := tbv.g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, tbv.quit); err != nil {
		return err
	}

	if err := tbv.g.SetKeybinding("", 'h', gocui.ModNone, tbv.previousDayView); err != nil {
		return err
	}
	if err := tbv.g.SetKeybinding("", 'l', gocui.ModNone, tbv.nextDayView); err != nil {
		return err
	}
	if err := tbv.g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, tbv.nextDayView); err != nil {
		return err
	}
	return nil
}
