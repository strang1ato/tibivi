package tibivi

import (
	"fmt"

	"github.com/oltarzewskik/tibivi-gocui"
)

// BlockUtils contains utility variables for adding, removing and modifying blocks
type BlockUtils struct {
	formFields                        []string
	selectedFormField                 int
	addingBlock                       bool
	selectedForMod, selectedForRemove bool
	selectedBlock                     int
}

// setAddBlockForm shows user view based add block form
func (tbv *Tibivi) setAddBlockForm() error {
	maxX, maxY := tbv.g.Size()
	if err := tbv.setForm(maxX, maxY); err != nil {
		return err
	}

	tbv.BlockUtils.selectedFormField = 0
	tbv.Views.currentViewOnTop = tbv.BlockUtils.formFields[tbv.BlockUtils.selectedFormField]

	tbv.BlockUtils.addingBlock = true
	return nil
}

// formFieldsNormalMode changes to normal mode
func (tbv *Tibivi) formFieldsNormalMode(g *gocui.Gui, v *gocui.View) error {
	if tbv.g.Cursor {
		if err := tbv.setViewsRuneKeybindings(tbv.BlockUtils.formFields, []rune{'l', 'L', 'j', 'J'}, gocui.ModNone,
			tbv.nextFormField); err != nil {
			return err
		}
		if err := tbv.setViewsRuneKeybindings(tbv.BlockUtils.formFields, []rune{'h', 'H', 'k', 'K'}, gocui.ModNone,
			tbv.previousFormField); err != nil {
			return err
		}
		if err := tbv.setViewsRuneKeybindings(tbv.BlockUtils.formFields, []rune{'i', 'I'}, gocui.ModNone,
			tbv.formFieldsInsertMode); err != nil {
			return err
		}

		tbv.Views.bar["bar"].Clear()

		for _, name := range tbv.BlockUtils.formFields {
			tbv.Views.block[name].Editable = false
		}
		tbv.g.Cursor = false
	}
	return nil
}

// formFieldsInsertMode changes to vi like insert mode
func (tbv *Tibivi) formFieldsInsertMode(g *gocui.Gui, v *gocui.View) error {
	if !tbv.g.Cursor {
		if err := tbv.deleteViewsRuneKeybindings(tbv.BlockUtils.formFields,
			[]rune{'l', 'L', 'j', 'J', 'h', 'H', 'k', 'K', 'i', 'I'}, gocui.ModNone); err != nil {
			return err
		}

		tbv.Views.bar["bar"].Clear()
		fmt.Fprint(tbv.Views.bar["bar"], "\x1b[1m"+"-- INSERT --"+"\x1b[0m")

		for _, name := range tbv.BlockUtils.formFields {
			tbv.Views.block[name].Editable = true
		}
		tbv.g.Cursor = true
	}
	return nil
}

// nextFormField goes to next form field
func (tbv *Tibivi) nextFormField(g *gocui.Gui, v *gocui.View) error {
	if !tbv.g.Cursor {
		tbv.BlockUtils.selectedFormField++
		if tbv.BlockUtils.selectedFormField > 2 {
			tbv.BlockUtils.selectedFormField = 0
		}
		tbv.Views.currentViewOnTop = tbv.BlockUtils.formFields[tbv.BlockUtils.selectedFormField]
	}
	return nil
}

// previousFormField goes to previous form field
func (tbv *Tibivi) previousFormField(g *gocui.Gui, v *gocui.View) error {
	if !tbv.g.Cursor {
		tbv.BlockUtils.selectedFormField--
		if tbv.BlockUtils.selectedFormField < 0 {
			tbv.BlockUtils.selectedFormField = 2
		}
		tbv.Views.currentViewOnTop = tbv.BlockUtils.formFields[tbv.BlockUtils.selectedFormField]
	}
	return nil
}

// submitAddBlockForm submits addBlock form
func (tbv *Tibivi) submitAddBlockForm(g *gocui.Gui, v *gocui.View) error {
	if tbv.BlockUtils.addingBlock {
		tbv.submitForm(tbv.addBlock, g, v)
	}
	return nil
}

// deleteForm deletes form view and all its field views
func (tbv *Tibivi) deleteForm(g *gocui.Gui, v *gocui.View) error {
	if !tbv.g.Cursor {
		tbv.Views.currentViewOnTop = tbv.days[tbv.g.SelectedDay]

		if err := tbv.g.DeleteView("form"); err != nil {
			return err
		}

		for _, name := range tbv.BlockUtils.formFields {
			if err := tbv.g.DeleteView(name); err != nil {
				return err
			}
		}

		tbv.BlockUtils.addingBlock = false
		tbv.BlockUtils.selectedForMod = false
	}
	return nil
}

// selectBlockForRemove selects block for modification
func (tbv *Tibivi) selectBlockForMod() {
	tbv.BlockUtils.selectedForMod = true
	tbv.updateLayout()

	tbv.Views.currentViewOnTop = tbv.days[tbv.g.SelectedDay]
}

// setModBlockForm shows user view based modify block form
func (tbv *Tibivi) setModBlockForm(g *gocui.Gui, v *gocui.View) error {
	day := tbv.Schedule[tbv.days[tbv.g.SelectedDay]]
	dayLen := len(day)
	if tbv.BlockUtils.selectedForMod && tbv.BlockUtils.selectedBlock < dayLen {
		maxX, maxY := tbv.g.Size()
		if err := tbv.setForm(maxX, maxY); err != nil {
			return err
		}
		block := day[tbv.BlockUtils.selectedBlock]
		fmt.Fprint(tbv.Views.block["formDescription"], block.description)

		startTime := block.startHour + ":" + block.startMinute
		finishTime := block.finishHour + ":" + block.finishMinute
		fmt.Fprint(tbv.Views.block["formStartTime"], startTime)
		fmt.Fprint(tbv.Views.block["formFinishTime"], finishTime)

		tbv.BlockUtils.selectedFormField = 0
		tbv.Views.currentViewOnTop = tbv.BlockUtils.formFields[tbv.BlockUtils.selectedFormField]
	}
	return nil
}

// submitModBlockForm submits modBlock form
func (tbv *Tibivi) submitModBlockForm(g *gocui.Gui, v *gocui.View) error {
	if tbv.BlockUtils.selectedForMod {
		tbv.submitForm(tbv.modBlock, g, v)
	}
	return nil
}

// modBlock modifies selected block
func (tbv *Tibivi) modBlock(startTime, finishTime, description string) error {
	block, err := tbv.createBlock(startTime, finishTime, description)
	if err != nil {
		return err
	}
	tbv.removeBlock()
	day := tbv.Schedule[tbv.days[tbv.g.SelectedDay]]
	tbv.Schedule[tbv.days[tbv.g.SelectedDay]] = sortDay(append(day, block))
	return nil
}

// selectBlockForRemove selects block for remove
func (tbv *Tibivi) selectBlockForRemove() {
	tbv.BlockUtils.selectedForRemove = true
	tbv.updateLayout()

	tbv.Views.currentViewOnTop = tbv.days[tbv.g.SelectedDay]
}

// removeSelctedBlock removes selected block
func (tbv *Tibivi) removeSelctedBlock(g *gocui.Gui, v *gocui.View) error {
	if tbv.BlockUtils.selectedForRemove {
		tbv.removeBlock()

		tbv.updateLayout()
	}
	return nil
}

// selectNextBlock selects next block
func (tbv *Tibivi) selectNextBlock(g *gocui.Gui, v *gocui.View) error {
	if tbv.BlockUtils.selectedForRemove || tbv.BlockUtils.selectedForMod {
		nextIndex := tbv.BlockUtils.selectedBlock + 1
		dayLen := len(tbv.Schedule[tbv.days[tbv.g.SelectedDay]])
		if nextIndex >= dayLen {
			nextIndex = dayLen - 1
		}
		tbv.BlockUtils.selectedBlock = nextIndex
		tbv.updateLayout()
	}
	return nil
}

// selectNextBlock selects next block
func (tbv *Tibivi) selectPreviousBlock(g *gocui.Gui, v *gocui.View) error {
	if tbv.BlockUtils.selectedForRemove || tbv.BlockUtils.selectedForMod {
		previousIndex := tbv.BlockUtils.selectedBlock - 1
		if previousIndex < 0 {
			previousIndex = 0
		}
		tbv.BlockUtils.selectedBlock = previousIndex
		tbv.updateLayout()
	}
	return nil
}

// exitFromBlockSelection exits from block selection
func (tbv *Tibivi) exitFromBlockSelection(g *gocui.Gui, v *gocui.View) error {
	tbv.BlockUtils.selectedBlock = 0
	tbv.BlockUtils.selectedForRemove = false
	tbv.BlockUtils.selectedForMod = false
	tbv.updateLayout()
	return nil
}

// setForm sets basis add/modify block form
func (tbv *Tibivi) setForm(maxX, maxY int) error {
	if v, err := tbv.g.SetView("form", maxX/3, maxY/3, maxX*2/3, maxY*2/3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		tbv.Views.block["form"] = v
	}

	if v, err := tbv.g.SetView("formStartTime", maxX/3+1, maxY/3+1, maxX/2-1, maxY/3+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Start time"

		tbv.Views.block["formStartTime"] = v
	}
	if v, err := tbv.g.SetView("formFinishTime", maxX/2, maxY/3+1, maxX*2/3-1, maxY/3+3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Finish time"

		tbv.Views.block["formFinishTime"] = v
	}
	if v, err := tbv.g.SetView("formDescription", maxX/3+1, maxY/3+4, maxX*2/3-1, maxY*2/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Description"
		v.Wrap = true

		tbv.Views.block["formDescription"] = v
	}
	return nil
}

// submitForm submits form if finish time field is selected
func (tbv *Tibivi) submitForm(f func(string, string, string) error, g *gocui.Gui, v *gocui.View) {
	if tbv.BlockUtils.selectedFormField > 1 {
		tbv.formFieldsNormalMode(g, v)
		err := f(tbv.Views.block["formStartTime"].Buffer(),
			tbv.Views.block["formFinishTime"].Buffer(),
			tbv.Views.block["formDescription"].Buffer())
		if err != nil {
			tbv.Views.bar["bar"].Clear()
			fmt.Fprint(tbv.Views.bar["bar"], "\x1b[41m"+err.Error()+"\x1b[0m")
		} else {
			tbv.deleteForm(g, v)
			tbv.updateLayout()
		}
	} else {
		tbv.BlockUtils.selectedFormField++
		tbv.Views.currentViewOnTop = tbv.BlockUtils.formFields[tbv.BlockUtils.selectedFormField]
	}
}

// removeBlock removes selected block
func (tbv *Tibivi) removeBlock() {
	day := tbv.Schedule[tbv.days[tbv.g.SelectedDay]]
	for i := range day {
		if i == tbv.BlockUtils.selectedBlock {
			tbv.Schedule[tbv.days[tbv.g.SelectedDay]] = append(day[:i], day[i+1:]...)
			break
		}
	}

	tbv.BlockUtils.selectedBlock = 0
	tbv.BlockUtils.selectedForMod = false
	tbv.BlockUtils.selectedForRemove = false
}
