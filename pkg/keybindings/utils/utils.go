package keybindings_utils

import (
	"github.com/strang1ato/tibivi-gocui"
	"github.com/strang1ato/tibivi/pkg/common"
)

// SetViewsKeybinding sets keybinding to set of views
func SetViewsKeybinding(viewnames []string, key interface{}, mod gocui.Modifier,
	handler func(*gocui.Gui, *gocui.View) error) error {
	for _, viewname := range viewnames {
		if err := common.G.SetKeybinding(viewname, key, mod, handler); err != nil {
			return err
		}
	}
	return nil
}

// SetViewsRuneKeybindings sets rune keybindings to set of views
func SetViewsRuneKeybindings(viewnames []string, keys []rune, mod gocui.Modifier,
	handler func(*gocui.Gui, *gocui.View) error) error {
	for _, viewname := range viewnames {
		if err := SetRuneKeybindings(viewname, keys, mod, handler); err != nil {
			return err
		}
	}
	return nil
}

// SetRuneKeybindings attaches set of rune keybindings to one function
func SetRuneKeybindings(viewname string, keys []rune, mod gocui.Modifier,
	handler func(*gocui.Gui, *gocui.View) error) error {
	for _, key := range keys {
		if err := common.G.SetKeybinding(viewname, key, mod, handler); err != nil {
			return err
		}
	}
	return nil
}

// DeleteViewsRuneKeybindings deletes set of rune keybindings from set of views
func DeleteViewsRuneKeybindings(viewnames []string, keys []rune, mod gocui.Modifier) error {
	for _, viewname := range viewnames {
		for _, key := range keys {
			if err := common.G.DeleteKeybinding(viewname, key, mod); err != nil {
				return err
			}
		}
	}
	return nil
}
