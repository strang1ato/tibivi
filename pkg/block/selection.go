package block

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/layout/utils"
)

// selectNextBlock selects next block
func selectNextBlock(g *gocui.Gui, v *gocui.View) error {
	if common.SelectBlockForRemove || common.SelectBlockForMod {
		// dayLen := len(common.Schedule[common.Days[common.G.SelectedDay]])
		blocksInBuffer := common.BlocksInBuffer[common.Days[common.G.SelectedDay]]
		if common.SelectedBlock < blocksInBuffer-1 {
			common.SelectedBlock++
		} else if common.SelectedBlock > blocksInBuffer-1 {
			common.SelectedBlock = blocksInBuffer - 1
		}
		layout_utils.UpdateLayout()
	}
	return nil
}

// selectPreviousBlock selects previous block
func selectPreviousBlock(g *gocui.Gui, v *gocui.View) error {
	if common.SelectBlockForRemove || common.SelectBlockForMod {
		if common.SelectedBlock > 0 {
			common.SelectedBlock--
			layout_utils.UpdateLayout()
		}
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
