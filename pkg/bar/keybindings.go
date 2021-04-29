package bar

import (
	"github.com/strang1ato/tibivi-gocui"
	"github.com/strang1ato/tibivi/pkg/common"
	"github.com/strang1ato/tibivi/pkg/keybindings/utils"
)

// Keybindings sets keybindings for bar
func Keybindings() error {
	keybindings_utils.SetViewsKeybinding(common.Days, ':', gocui.ModNone, focusBar)
	common.G.SetKeybinding("bar", gocui.KeyEsc, gocui.ModNone, unFocusBar)
	if err := common.G.SetKeybinding("bar", gocui.KeyEnter, gocui.ModNone, executeCommand); err != nil {
		return err
	}
	return nil
}
