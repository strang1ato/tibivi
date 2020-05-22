package block

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/keybindings/utils"
)

func Keybindings() error {
	// Keybindings for add block
	keybindings_utils.SetViewsKeybinding(formFields, gocui.KeyEnter, gocui.ModNone, submitAddBlockForm)

	// Keybindings for modify block
	keybindings_utils.SetViewsKeybinding(common.Days, gocui.KeyEnter, gocui.ModNone, setModBlockForm)
	keybindings_utils.SetViewsKeybinding(formFields, gocui.KeyEnter, gocui.ModNone, submitModBlockForm)

	// Keybindings for remove block
	keybindings_utils.SetViewsKeybinding(common.Days, gocui.KeyEnter, gocui.ModNone, removeSelctedBlock)

	// Keybindings related to form
	if err := keybindings_utils.SetViewsKeybinding(formFields, gocui.KeyEsc, gocui.ModNone, formFieldsNormalMode); err != nil {
		return err
	}
	if err := keybindings_utils.SetViewsRuneKeybindings(formFields, []rune{'i', 'I'}, gocui.ModNone,
		formFieldsInsertMode); err != nil {
		return err
	}
	keybindings_utils.SetViewsRuneKeybindings(formFields, []rune{'l', 'L', 'j', 'J'}, gocui.ModNone, nextFormField)
	keybindings_utils.SetViewsRuneKeybindings(formFields, []rune{'h', 'H', 'k', 'K'}, gocui.ModNone, previousFormField)
	if err := keybindings_utils.SetViewsKeybinding(formFields, gocui.KeyEsc, gocui.ModNone, deleteForm); err != nil {
		return err
	}

	// Keybindings for block selection
	keybindings_utils.SetViewsRuneKeybindings(common.Days, []rune{'j', 'J'}, gocui.ModNone, selectNextBlock)
	keybindings_utils.SetViewsRuneKeybindings(common.Days, []rune{'k', 'K'}, gocui.ModNone, selectPreviousBlock)
	keybindings_utils.SetViewsKeybinding(common.Days, gocui.KeyEsc, gocui.ModNone, exitFromBlockSelection)
	return nil
}
