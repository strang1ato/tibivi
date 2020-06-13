package menu

import (
	"fmt"

	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/block"
	"github.com/oltarzewskik/tibivi/pkg/common"
)

var (
	options        = []string{"Add block", "Modify block", "Remove block"}
	selectedOption int
)

// setMenu shows add/modify/remove block menu
func setMenu(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := common.G.Size()
	if v, err := common.G.SetView("menu", maxX/3, maxY/3, maxX*2/3, maxY*2/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		for i, option := range options {
			if i == selectedOption {
				fmt.Fprintln(v, "\x1b[7m"+option+"\x1b[0m")
				continue
			}
			fmt.Fprintln(v, option)
		}

		common.CurrentViewOnTop = "menu"

		common.Views.Menu = v
	}
	if common.SelectBlockForRemove || common.SelectBlockForMod {
		block.ExitFromBlockSelection(g, v)
	}
	common.Views.Bar["bar"].Clear()
	return nil
}

// runselectedOption runs selected menu option
func runSelectedMenuOption(g *gocui.Gui, v *gocui.View) error {
	switch selectedOption {
	case 0:
		if err := block.SetAddBlockForm(); err != nil {
			return err
		}
	case 1:
		block.SelectBlockForMod()
	case 2:
		block.SelectBlockForRemove()
	}
	selectedOption = 0
	if err := common.G.DeleteView("menu"); err != nil {
		return err
	}
	return nil
}

// nextMenuOption goes to next Menu option
func nextMenuOption(g *gocui.Gui, v *gocui.View) error {
	nextIndex := selectedOption + 1
	if nextIndex > 2 {
		nextIndex = 0
	}
	common.Views.Menu.Clear()
	for i, option := range options {
		if i == nextIndex {
			fmt.Fprintln(v, "\x1b[7m"+option+"\x1b[0m")
			continue
		}
		fmt.Fprintln(v, option)
	}

	selectedOption = nextIndex
	return nil
}

// previousMenuOption goes to previous Menu option
func previousMenuOption(g *gocui.Gui, v *gocui.View) error {
	previousIndex := selectedOption - 1
	if previousIndex < 0 {
		previousIndex = 2
	}
	common.Views.Menu.Clear()
	for i, option := range options {
		if i == previousIndex {
			fmt.Fprintln(v, "\x1b[7m"+option+"\x1b[0m")
			continue
		}
		fmt.Fprintln(v, option)
	}

	selectedOption = previousIndex
	return nil
}

// deleteMenu deletes menu
func deleteMenu(g *gocui.Gui, v *gocui.View) error {
	common.CurrentViewOnTop = common.Days[common.G.SelectedDay]
	if err := common.G.DeleteView("menu"); err != nil {
		return err
	}
	return nil
}
