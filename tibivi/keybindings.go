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
	return nil
}

// setViewsKeybinding sets certain keybinding to set of views
func (tbv *Tibivi) setViewsKeybinding(viewnames []string, key interface{}, mod gocui.Modifier,
	handler func(*gocui.Gui, *gocui.View) error) error {
	for _, viewname := range viewnames {
		if err := tbv.g.SetKeybinding(viewname, key, mod, handler); err != nil {
			return err
		}
	}
	return nil
}

// setViewsRuneKeybindings sets certain rune keybindings to set of views
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
