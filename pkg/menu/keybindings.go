package menu

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/keybindings/utils"
)

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
	return nil
}
