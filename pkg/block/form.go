package block

import (
	"fmt"

	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/keybindings/utils"
	"github.com/oltarzewskik/tibivi/pkg/layout/utils"
)

var (
	formFields        = []string{"formStartTime", "formFinishTime", "formDescription"}
	selectedFormField int
)

// setForm sets basis add/modify block form
func setForm(maxX, maxY int) error {
	if v, err := common.G.SetView("form", maxX/3, maxY/3, maxX*2/3, maxY*2/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		common.Views.Block["form"] = v
	}

	if v, err := common.G.SetView("formStartTime", maxX/3+1, maxY/3+1, maxX/2-1, maxY/3+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Start time"

		common.Views.Block["formStartTime"] = v
	}
	if v, err := common.G.SetView("formFinishTime", maxX/2, maxY/3+1, maxX*2/3-1, maxY/3+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Finish time"

		common.Views.Block["formFinishTime"] = v
	}
	if v, err := common.G.SetView("formDescription", maxX/3+1, maxY/3+4, maxX*2/3-1, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Description"
		v.Wrap = true

		common.Views.Block["formDescription"] = v
	}
	return nil
}

// submitForm submits form if finish time field is selected
func submitForm(action func(string, string, string) error, g *gocui.Gui, v *gocui.View) {
	if selectedFormField > 1 {
		formNormalModeOrDeleteForm(g, v)
		err := action(common.Views.Block["formStartTime"].Buffer(),
			common.Views.Block["formFinishTime"].Buffer(),
			common.Views.Block["formDescription"].Buffer())
		if err != nil {
			common.Views.Bar["bar"].Clear()
			fmt.Fprint(common.Views.Bar["bar"], "\x1b[41m"+err.Error()+"\x1b[0m")
		} else {
			deleteForm()
			layout_utils.UpdateLayout()
		}
	} else {
		selectedFormField++
		common.CurrentViewOnTop = formFields[selectedFormField]
	}
}

// formNormalModeOrDeleteForm changes to normal mode if insert mode, otherwise deletes form
func formNormalModeOrDeleteForm(g *gocui.Gui, v *gocui.View) error {
	if common.G.Cursor {
		keybindings_utils.SetViewsRuneKeybindings(formFields, []rune{'l', 'L', 'j', 'J'}, gocui.ModNone, nextFormField)
		keybindings_utils.SetViewsRuneKeybindings(formFields, []rune{'h', 'H', 'k', 'K'}, gocui.ModNone, previousFormField)

		if err := keybindings_utils.SetViewsRuneKeybindings(formFields, []rune{'i', 'I'}, gocui.ModNone,
			formInsertMode); err != nil {
			return err
		}

		common.Views.Bar["bar"].Clear()

		for _, name := range formFields {
			common.Views.Block[name].Editable = false
		}
		common.G.Cursor = false
	} else {
		deleteForm()
	}
	return nil
}

// formInsertMode changes to vi like insert mode
func formInsertMode(g *gocui.Gui, v *gocui.View) error {
	if !common.G.Cursor {
		if err := keybindings_utils.DeleteViewsRuneKeybindings(formFields,
			[]rune{'l', 'L', 'j', 'J', 'h', 'H', 'k', 'K', 'i', 'I'}, gocui.ModNone); err != nil {
			return err
		}

		common.Views.Bar["bar"].Clear()
		fmt.Fprint(common.Views.Bar["bar"], "\x1b[1m"+"-- INSERT --"+"\x1b[0m")

		for _, name := range formFields {
			common.Views.Block[name].Editable = true
		}
		common.G.Cursor = true
	}
	return nil
}

// nextFormField goes to next form field
func nextFormField(g *gocui.Gui, v *gocui.View) error {
	if !common.G.Cursor {
		selectedFormField++
		if selectedFormField > 2 {
			selectedFormField = 0
		}
		common.CurrentViewOnTop = formFields[selectedFormField]
	}
	return nil
}

// previousFormField goes to previous form field
func previousFormField(g *gocui.Gui, v *gocui.View) error {
	if !common.G.Cursor {
		selectedFormField--
		if selectedFormField < 0 {
			selectedFormField = 2
		}
		common.CurrentViewOnTop = formFields[selectedFormField]
	}
	return nil
}

// deleteForm deletes form view and all its field views
func deleteForm() error {
	if !common.G.Cursor {
		common.CurrentViewOnTop = common.Days[common.G.SelectedDay]

		if err := common.G.DeleteView("form"); err != nil {
			return err
		}

		for _, name := range formFields {
			if err := common.G.DeleteView(name); err != nil {
				return err
			}
		}

		addingBlock = false
		common.SelectBlockForMod = false
		layout_utils.UpdateLayout()
	}
	return nil
}
