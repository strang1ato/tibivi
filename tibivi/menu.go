package tibivi

import (
	"fmt"

	"github.com/oltarzewskik/tibivi-gocui"
)

var (
	menuOptions                               []string = []string{"Add block", "Modify block", "Remove block"}
	selectedMenuOption, selectedAddBlockField int
)

// setMenu shows add/modify/remove block menu
func (tbv *Tibivi) setMenu(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := tbv.g.Size()
	if v, err := tbv.g.SetView("menu", maxX/3, maxY/3, maxX*2/3, maxY*2/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		for i, option := range menuOptions {
			if i == selectedMenuOption {
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

// deleteMenu deletes menu
func (tbv *Tibivi) deleteMenu(g *gocui.Gui, v *gocui.View) error {
	tbv.Views.currentViewOnTop = tbv.days[tbv.g.SelectedDay]
	if err := tbv.g.DeleteView("menu"); err != nil {
		return err
	}
	return nil
}

// previousMenuOption goes to previous menu option
func (tbv *Tibivi) previousMenuOption(g *gocui.Gui, v *gocui.View) error {
	previousIndex := selectedMenuOption - 1
	if previousIndex < 0 {
		previousIndex = 2
	}
	tbv.Views.menu["menu"].Clear()
	for i, option := range menuOptions {
		if i == previousIndex {
			fmt.Fprintln(v, "\x1b[7m"+option+"\x1b[0m")
			continue
		}
		fmt.Fprintln(v, option)
	}

	selectedMenuOption = previousIndex
	return nil
}

// nextMenuOption goes to next menu option
func (tbv *Tibivi) nextMenuOption(g *gocui.Gui, v *gocui.View) error {
	nextIndex := selectedMenuOption + 1
	if nextIndex > 2 {
		nextIndex = 0
	}
	tbv.Views.menu["menu"].Clear()
	for i, option := range menuOptions {
		if i == nextIndex {
			fmt.Fprintln(v, "\x1b[7m"+option+"\x1b[0m")
			continue
		}
		fmt.Fprintln(v, option)
	}

	selectedMenuOption = nextIndex
	return nil
}

// runSelectedMenuOption runs selected menu option
func (tbv *Tibivi) runSelectedMenuOption(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := tbv.g.Size()
	switch selectedMenuOption {
	case 0:
		tbv.setAddBlock(maxX, maxY)
	}
	if err := tbv.g.DeleteView("menu"); err != nil {
		return err
	}
	return nil
}

// setAddBlock shows user view based add block form
func (tbv *Tibivi) setAddBlock(maxX, maxY int) error {
	if v, err := tbv.g.SetView("addBlock", maxX/3, maxY/3, maxX*2/3, maxY*2/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		tbv.Views.menu["addBlock"] = v
	}

	if v, err := tbv.g.SetView("addBlockDescription", maxX/3+1, maxY/3+1, maxX*2/3-1, maxY*2/3-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Description"
		v.Wrap = true

		tbv.Views.menu["addBlockDescription"] = v
	}
	if v, err := tbv.g.SetView("addBlockStartTime", maxX/3+1, maxY*2/3-3, maxX/2-1, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Start time"

		tbv.Views.menu["addBlockStartTime"] = v
	}
	if v, err := tbv.g.SetView("addBlockFinishTime", maxX/2, maxY*2/3-3, maxX*2/3-1, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Finish time"

		tbv.Views.menu["addBlockFinishTime"] = v
	}

	selectedAddBlockField = 0
	tbv.Views.currentViewOnTop = tbv.Views.addBlockFields[selectedAddBlockField]
	return nil

}

// addBlockFieldsNormalMode changes to normal mode
func (tbv *Tibivi) addBlockFieldsNormalMode(g *gocui.Gui, v *gocui.View) error {
	tbv.Views.bar["bar"].Clear()

	if tbv.g.Cursor {
		for _, name := range tbv.Views.addBlockFields {
			tbv.Views.menu[name].Editable = false
		}
		tbv.g.Cursor = false
	}

	if err := tbv.setViewsRuneKeybindings(tbv.Views.addBlockFields, []rune{'l', 'L', 'j', 'J'}, gocui.ModNone,
		tbv.nextAddBlockField); err != nil {
		return err
	}
	if err := tbv.setViewsRuneKeybindings(tbv.Views.addBlockFields, []rune{'h', 'H', 'k', 'K'}, gocui.ModNone,
		tbv.previousAddBlockField); err != nil {
		return err
	}
	if err := tbv.setViewsRuneKeybindings(tbv.Views.addBlockFields, []rune{'i', 'I'}, gocui.ModNone,
		tbv.addBlockFieldsInsertMode); err != nil {
		return err
	}
	return nil
}

// addBlockFieldsInsertMode changes to vi like insert mode
func (tbv *Tibivi) addBlockFieldsInsertMode(g *gocui.Gui, v *gocui.View) error {
	if !tbv.g.Cursor {
		tbv.Views.bar["bar"].Clear()
		fmt.Fprint(tbv.Views.bar["bar"], "\x1b[1m"+"-- INSERT --"+"\x1b[0m")
		for _, name := range tbv.Views.addBlockFields {
			tbv.Views.menu[name].Editable = true
		}

		tbv.g.Cursor = true
	}

	if err := tbv.deleteViewsRuneKeybindings(tbv.Views.addBlockFields,
		[]rune{'l', 'L', 'j', 'J', 'h', 'H', 'k', 'K', 'i', 'I'}, gocui.ModNone); err != nil {
		return err
	}
	return nil
}

// deleteAddBlock deletes addBlock view and all its field views
func (tbv *Tibivi) deleteAddBlock(g *gocui.Gui, v *gocui.View) error {
	if !tbv.g.Cursor {
		tbv.Views.currentViewOnTop = tbv.days[tbv.g.SelectedDay]

		if err := tbv.g.DeleteView("addBlock"); err != nil {
			return err
		}

		for _, name := range tbv.Views.addBlockFields {
			if err := tbv.g.DeleteView(name); err != nil {
				return err
			}
		}
	}
	return nil
}

// submitAddBlock submits new block if addBlock finish time field is selected
func (tbv *Tibivi) submitAddBlock(g *gocui.Gui, v *gocui.View) error {
	if selectedAddBlockField > 1 {
		tbv.addBlockFieldsNormalMode(g, v)
		err := tbv.addBlock(tbv.Views.menu["addBlockStartTime"].Buffer(),
			tbv.Views.menu["addBlockFinishTime"].Buffer(),
			tbv.Views.menu["addBlockDescription"].Buffer())
		if err != nil {
			tbv.Views.bar["bar"].Clear()
			fmt.Fprint(tbv.Views.bar["bar"], "\x1b[41m"+err.Error()+"\x1b[0m")
		} else {
			tbv.deleteAddBlock(g, v)
			tbv.updateLayout()
		}
	} else {
		selectedAddBlockField++
		tbv.Views.currentViewOnTop = tbv.Views.addBlockFields[selectedAddBlockField]
	}
	return nil
}

// nextAddBlockField goes to next addBlock field
func (tbv *Tibivi) nextAddBlockField(g *gocui.Gui, v *gocui.View) error {
	if !tbv.g.Cursor {
		selectedAddBlockField++
		if selectedAddBlockField > 2 {
			selectedAddBlockField = 0
		}
		tbv.Views.currentViewOnTop = tbv.Views.addBlockFields[selectedAddBlockField]
	}
	return nil
}

// previousAddBlockField goes to previous addBlock field
func (tbv *Tibivi) previousAddBlockField(g *gocui.Gui, v *gocui.View) error {
	if !tbv.g.Cursor {
		selectedAddBlockField--
		if selectedAddBlockField < 0 {
			selectedAddBlockField = 2
		}
		tbv.Views.currentViewOnTop = tbv.Views.addBlockFields[selectedAddBlockField]
	}
	return nil
}
