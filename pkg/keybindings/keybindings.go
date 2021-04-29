package keybindings

import (
	"github.com/strang1ato/tibivi-gocui"
	"github.com/strang1ato/tibivi/pkg/actions"
	"github.com/strang1ato/tibivi/pkg/bar"
	"github.com/strang1ato/tibivi/pkg/block"
	"github.com/strang1ato/tibivi/pkg/common"
	"github.com/strang1ato/tibivi/pkg/days"
	"github.com/strang1ato/tibivi/pkg/menu"
)

// Keybindings create keyboard keybindings
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
