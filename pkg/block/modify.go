package block

import (
	"fmt"

	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/data"
	"github.com/oltarzewskik/tibivi/pkg/layout/utils"
)

// SelectBlockForMod selects block for modification
func SelectBlockForMod() {
	common.SelectBlockForMod = true
	layout_utils.UpdateLayout()

	common.CurrentViewOnTop = common.Days[common.G.SelectedDay]
}

// setModBlockForm shows user view based modify block form
func setModBlockForm(g *gocui.Gui, v *gocui.View) error {
	day := data.Schedule[common.Days[common.G.SelectedDay]]
	dayLen := len(day)
	if common.SelectBlockForMod && common.SelectedBlock < dayLen {
		maxX, maxY := common.G.Size()
		if err := setForm(maxX, maxY); err != nil {
			return err
		}
		block := day[common.SelectedBlock]
		fmt.Fprint(common.Views.Block["formDescription"], block.Description)

		startTime := block.StartHour + ":" + block.StartMinute
		finishTime := block.FinishHour + ":" + block.FinishMinute
		fmt.Fprint(common.Views.Block["formStartTime"], startTime)
		fmt.Fprint(common.Views.Block["formFinishTime"], finishTime)

		selectedFormField = 0
		common.CurrentViewOnTop = formFields[selectedFormField]
	}
	return nil
}

// submitModBlockForm submits ModBlock form
func submitModBlockForm(g *gocui.Gui, v *gocui.View) error {
	if common.SelectBlockForMod {
		submitForm(modBlock, g, v)
	}
	return nil
}

// modBlock modifies selected block
func modBlock(startTime, finishTime, Description string) error {
	block, err := createBlock(startTime, finishTime, Description)
	if err != nil {
		return err
	}
	removeBlock()
	day := data.Schedule[common.Days[common.G.SelectedDay]]
	data.Schedule[common.Days[common.G.SelectedDay]] = data.SortDay(append(day, block))
	return nil
}
