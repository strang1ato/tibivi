package keybindings

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/actions"
	"github.com/oltarzewskik/tibivi/pkg/bar"
	"github.com/oltarzewskik/tibivi/pkg/block"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/days"
	"github.com/oltarzewskik/tibivi/pkg/menu"
)

// keybindings create keyboard keybindings
func Keybindings() error {
	// Keybinding for exiting tibivi without saving current state
	if err := common.G.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, actions.QuitIgnore); err != nil {
		return err
	}

	days.Keybindings()
	if err := bar.Keybindings(); err != nil {
		return err
	}
	if err := menu.Keybindings(); err != nil {
		return err
	}
	if err := block.Keybindings(); err != nil {
		return err
	}
	return nil
}
