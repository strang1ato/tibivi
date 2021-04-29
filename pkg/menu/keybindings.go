package menu

import (
	"github.com/strang1ato/tibivi-gocui"
	"github.com/strang1ato/tibivi/pkg/common"
	"github.com/strang1ato/tibivi/pkg/keybindings/utils"
)

// Keybindings sets keybindings for menu
func Keybindings() error {
	keybindings_utils.SetViewsRuneKeybindings(common.Days, []rune{'m', 'M'}, gocui.ModNone, setMenu)
	if err := common.G.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, runSelectedMenuOption); err != nil {
		return err
	}
	keybindings_utils.SetRuneKeybindings("menu", []rune{'j', 'J'}, gocui.ModNone, nextMenuOption)
	keybindings_utils.SetRuneKeybindings("menu", []rune{'k', 'K'}, gocui.ModNone, previousMenuOption)
	if err := common.G.SetKeybinding("menu", gocui.KeyEsc, gocui.ModNone, deleteMenu); err != nil {
		return err
	}
	if err := keybindings_utils.SetRuneKeybindings("menu", []rune{'q', 'Q'}, gocui.ModNone, deleteMenu); err != nil {
		return err
	}
	return nil
}
