package tibivi

import (
	"fmt"

	"github.com/oltarzewskik/tibivi-gocui"
)

// MenuUtils contains utility variables for Menu
type MenuUtils struct {
	options        []string
	selectedOption int
}

// setMenu shows add/modify/remove block menu
func (tbv *Tibivi) setMenu(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := tbv.g.Size()
	if v, err := tbv.g.SetView("menu", maxX/3, maxY/3, maxX*2/3, maxY*2/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		for i, option := range tbv.MenuUtils.options {
			if i == tbv.MenuUtils.selectedOption {
				fmt.Fprintln(v, "\x1b[7m"+option+"\x1b[0m")
				continue
			}
			fmt.Fprintln(v, option)
		}

		tbv.Views.currentViewOnTop = "menu"

		tbv.Views.menu["menu"] = v
	}

	tbv.Views.bar["bar"].Clear()
	return nil
}

// runtbv.MenuUtils.selectedOption runs selected menu option
func (tbv *Tibivi) runSelectedMenuOption(g *gocui.Gui, v *gocui.View) error {
	switch tbv.MenuUtils.selectedOption {
	case 0:
		if err := tbv.setAddBlockForm(); err != nil {
			return err
		}
	case 1:
		tbv.selectBlockForMod()
	case 2:
		tbv.selectBlockForRemove()
	}
	tbv.MenuUtils.selectedOption = 0
	if err := tbv.g.DeleteView("menu"); err != nil {
		return err
	}
	return nil
}

// previousMenuOption goes to previous menu option
func (tbv *Tibivi) previousMenuOption(g *gocui.Gui, v *gocui.View) error {
	previousIndex := tbv.MenuUtils.selectedOption - 1
	if previousIndex < 0 {
		previousIndex = 2
	}
	tbv.Views.menu["menu"].Clear()
	for i, option := range tbv.MenuUtils.options {
		if i == previousIndex {
			fmt.Fprintln(v, "\x1b[7m"+option+"\x1b[0m")
			continue
		}
		fmt.Fprintln(v, option)
	}

	tbv.MenuUtils.selectedOption = previousIndex
	return nil
}

// nextMenuOption goes to next menu option
func (tbv *Tibivi) nextMenuOption(g *gocui.Gui, v *gocui.View) error {
	nextIndex := tbv.MenuUtils.selectedOption + 1
	if nextIndex > 2 {
		nextIndex = 0
	}
	tbv.Views.menu["menu"].Clear()
	for i, option := range tbv.MenuUtils.options {
		if i == nextIndex {
			fmt.Fprintln(v, "\x1b[7m"+option+"\x1b[0m")
			continue
		}
		fmt.Fprintln(v, option)
	}

	tbv.MenuUtils.selectedOption = nextIndex
	return nil
}

// deleteMenu deletes menu
func (tbv *Tibivi) deleteMenu(g *gocui.Gui, v *gocui.View) error {
	tbv.Views.currentViewOnTop = tbv.days[tbv.g.SelectedDay]
	if err := tbv.g.DeleteView("menu"); err != nil {
		return err
	}
	return nil
}
