package tibivi

import (
	"github.com/oltarzewskik/tibivi-gocui"
)

// keybindings create keyboard keybindings
func (tbv *Tibivi) keybindings() error {
	// Keybinding for exiting tibivi without saving current state
	if err := tbv.g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, tbv.quitIgnore); err != nil {
		return err
	}

	// Keybindings for selecting day view
	if err := tbv.setViewsRuneKeybindings(tbv.days, []rune{'h', 'H'}, gocui.ModNone, tbv.previousDayView); err != nil {
		return err
	}
	if err := tbv.setViewsRuneKeybindings(tbv.days, []rune{'l', 'L'}, gocui.ModNone, tbv.nextDayView); err != nil {
		return err
	}
	if err := tbv.setViewsKeybinding(tbv.days, gocui.KeyTab, gocui.ModNone, tbv.nextDayView); err != nil {
		return err
	}

	// Keybindings related to command bar
	if err := tbv.setViewsKeybinding(tbv.days, ':', gocui.ModNone, tbv.focusBar); err != nil {
		return err
	}
	if err := tbv.g.SetKeybinding("bar", gocui.KeyEsc, gocui.ModNone, tbv.unfocusBar); err != nil {
		return err
	}
	if err := tbv.g.SetKeybinding("bar", gocui.KeyEnter, gocui.ModNone, tbv.executeCommand); err != nil {
		return err
	}

	// Keybindings for menu
	if err := tbv.setViewsRuneKeybindings(tbv.days, []rune{'m', 'M'}, gocui.ModNone, tbv.setMenu); err != nil {
		return err
	}
	if err := tbv.g.SetKeybinding("menu", gocui.KeyEsc, gocui.ModNone, tbv.deleteMenu); err != nil {
		return err
	}
	if err := tbv.setRuneKeybindings("menu", []rune{'k', 'K'}, gocui.ModNone, tbv.previousMenuOption); err != nil {
		return err
	}
	if err := tbv.setRuneKeybindings("menu", []rune{'j', 'J'}, gocui.ModNone, tbv.nextMenuOption); err != nil {
		return err
	}
	if err := tbv.g.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, tbv.runSelectedMenuOption); err != nil {
		return err
	}

	// Keybindings for add block
	if err := tbv.setViewsKeybinding(tbv.BlockUtils.formFields, gocui.KeyEsc, gocui.ModNone, tbv.deleteForm); err != nil {
		return err
	}
	if err := tbv.setViewsKeybinding(tbv.BlockUtils.formFields, gocui.KeyEsc, gocui.ModNone, tbv.formFieldsNormalMode); err != nil {
		return err
	}
	if err := tbv.setViewsRuneKeybindings(tbv.BlockUtils.formFields, []rune{'i', 'I'}, gocui.ModNone,
		tbv.formFieldsInsertMode); err != nil {
		return err
	}
	if err := tbv.setViewsKeybinding(tbv.BlockUtils.formFields, gocui.KeyEnter, gocui.ModNone, tbv.submitAddBlockForm); err != nil {
		return err
	}
	if err := tbv.setViewsRuneKeybindings(tbv.BlockUtils.formFields, []rune{'l', 'L', 'j', 'J'}, gocui.ModNone,
		tbv.nextFormField); err != nil {
		return err
	}
	if err := tbv.setViewsRuneKeybindings(tbv.BlockUtils.formFields, []rune{'h', 'H', 'k', 'K'}, gocui.ModNone,
		tbv.previousFormField); err != nil {
		return err
	}

	// Keybindings for modify/remove block
	if err := tbv.setViewsKeybinding(tbv.days, gocui.KeyEnter, gocui.ModNone, tbv.setModBlockForm); err != nil {
		return err
	}
	if err := tbv.setViewsKeybinding(tbv.BlockUtils.formFields, gocui.KeyEnter, gocui.ModNone, tbv.submitModBlockForm); err != nil {
		return err
	}
	if err := tbv.setViewsRuneKeybindings(tbv.days, []rune{'j', 'J'}, gocui.ModNone, tbv.selectNextBlock); err != nil {
		return err
	}
	if err := tbv.setViewsRuneKeybindings(tbv.days, []rune{'k', 'K'}, gocui.ModNone, tbv.selectPreviousBlock); err != nil {
		return err
	}
	if err := tbv.setViewsKeybinding(tbv.days, gocui.KeyEsc, gocui.ModNone, tbv.exitFromBlockSelection); err != nil {
		return err
	}
	if err := tbv.setViewsKeybinding(tbv.days, gocui.KeyEnter, gocui.ModNone, tbv.removeSelctedBlock); err != nil {
		return err
	}
	return nil
}

// setViewsKeybinding sets keybinding to set of views
func (tbv *Tibivi) setViewsKeybinding(viewnames []string, key interface{}, mod gocui.Modifier,
	handler func(*gocui.Gui, *gocui.View) error) error {
	for _, viewname := range viewnames {
		if err := tbv.g.SetKeybinding(viewname, key, mod, handler); err != nil {
			return err
		}
	}
	return nil
}

// setViewsRuneKeybindings sets rune keybindings to set of views
func (tbv *Tibivi) setViewsRuneKeybindings(viewnames []string, keys []rune, mod gocui.Modifier,
	handler func(*gocui.Gui, *gocui.View) error) error {
	for _, viewname := range viewnames {
		if err := tbv.setRuneKeybindings(viewname, keys, mod, handler); err != nil {
			return err
		}
	}
	return nil
}

// setRuneKeybindings attaches set of rune keybindings to one function
func (tbv *Tibivi) setRuneKeybindings(viewname string, keys []rune, mod gocui.Modifier,
	handler func(*gocui.Gui, *gocui.View) error) error {
	for _, key := range keys {
		if err := tbv.g.SetKeybinding(viewname, key, mod, handler); err != nil {
			return err
		}
	}
	return nil
}

// deleteViewsRuneKeybindings deletes set of rune keybindings from set of views
func (tbv *Tibivi) deleteViewsRuneKeybindings(viewnames []string, keys []rune, mod gocui.Modifier) error {
	for _, viewname := range viewnames {
		for _, key := range keys {
			if err := tbv.g.DeleteKeybinding(viewname, key, mod); err != nil {
				return err
			}
		}
	}
	return nil
}
