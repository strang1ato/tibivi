package block

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/data"
	"github.com/oltarzewskik/tibivi/pkg/layout/utils"
)

// selectNextBlock selects next block
func selectNextBlock(g *gocui.Gui, v *gocui.View) error {
	if common.SelectBlockForRemove || common.SelectBlockForMod {
		nextIndex := common.SelectedBlock + 1
		dayLen := len(data.Schedule[common.Days[common.G.SelectedDay]])
		if nextIndex >= dayLen {
			nextIndex = dayLen - 1
		}
		common.SelectedBlock = nextIndex
		layout_utils.UpdateLayout()
	}
	return nil
}

// selectPreviousBlock selects previous block
func selectPreviousBlock(g *gocui.Gui, v *gocui.View) error {
	if common.SelectBlockForRemove || common.SelectBlockForMod {
		previousIndex := common.SelectedBlock - 1
		if previousIndex < 0 {
			previousIndex = 0
		}
		common.SelectedBlock = previousIndex
		layout_utils.UpdateLayout()
	}
	return nil
}

// ExitFromBlockSelection exits from block selection
func ExitFromBlockSelection(g *gocui.Gui, v *gocui.View) error {
	common.SelectedBlock = 0
	common.SelectBlockForRemove = false
	common.SelectBlockForMod = false
	layout_utils.UpdateLayout()
	return nil
}
