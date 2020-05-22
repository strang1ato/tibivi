package bar

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/keybindings/utils"
)

func Keybindings() error {
	keybindings_utils.SetViewsKeybinding(common.Days, ':', gocui.ModNone, focusBar)
	common.G.SetKeybinding("bar", gocui.KeyEsc, gocui.ModNone, unFocusBar)
	if err := common.G.SetKeybinding("bar", gocui.KeyEnter, gocui.ModNone, executeCommand); err != nil {
		return err
	}
	return nil
}
