package days

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/keybindings/utils"
)

// Keybindings sets keybindings for menu
func Keybindings() {
	keybindings_utils.SetViewsRuneKeybindings(common.Days, []rune{'f', 'F'}, gocui.ModNone, switchDayViewFocus)
	keybindings_utils.SetViewsKeybinding(common.Days, gocui.KeyEsc, gocui.ModNone, unFocusDayView)
	keybindings_utils.SetViewsKeybinding(common.Days, gocui.KeyTab, gocui.ModNone, nextDayView)
	keybindings_utils.SetViewsRuneKeybindings(common.Days, []rune{'l', 'L'}, gocui.ModNone, nextDayView)
	keybindings_utils.SetViewsRuneKeybindings(common.Days, []rune{'h', 'H'}, gocui.ModNone, previousDayView)
	keybindings_utils.SetViewsRuneKeybindings(common.Days, []rune{'j', 'J'}, gocui.ModNone, shiftBlock)
	keybindings_utils.SetViewsRuneKeybindings(common.Days, []rune{'k', 'K'}, gocui.ModNone, unShiftBlock)
}
