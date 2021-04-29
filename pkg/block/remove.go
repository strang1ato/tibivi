package block

import (
	"github.com/strang1ato/tibivi-gocui"
	"github.com/strang1ato/tibivi/pkg/common"
	"github.com/strang1ato/tibivi/pkg/layout/utils"
)

// SelectBlockForRemove selects block for remove
func SelectBlockForRemove() {
	common.SelectBlockForRemove = true
	layout_utils.UpdateLayout()

	common.CurrentViewOnTop = common.Days[common.G.SelectedDay]
}

// removeSelctedBlock removes selected block
func removeSelctedBlock(g *gocui.Gui, v *gocui.View) error {
	if common.SelectBlockForRemove {
		removeBlock()

		layout_utils.UpdateLayout()
	}
	return nil
}

// removeBlock removes selected block
func removeBlock() {
	selectedDay := common.Days[common.G.SelectedDay]
	day := common.Schedule[selectedDay][common.Shift[selectedDay]:]
	for i := range day {
		if i == common.SelectedBlock {
			common.Schedule[selectedDay] = append(common.Schedule[selectedDay][:i+common.Shift[selectedDay]],
				common.Schedule[selectedDay][i+common.Shift[selectedDay]+1:]...)

			common.UpdatedDays[selectedDay] = true
			break
		}
	}

	common.SelectedBlock = 0
	common.SelectBlockForMod = false
	common.SelectBlockForRemove = false
}
