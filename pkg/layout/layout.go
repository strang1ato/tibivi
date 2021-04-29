package layout

import (
	"github.com/strang1ato/tibivi-gocui"
	"github.com/strang1ato/tibivi/pkg/bar"
	"github.com/strang1ato/tibivi/pkg/common"
	"github.com/strang1ato/tibivi/pkg/days"
	"github.com/strang1ato/tibivi/pkg/days/utils"
)

// Layout setups CUI layout
func Layout(g *gocui.Gui) error {
	maxX, maxY := common.G.Size()
	if common.Focus {
		day := common.Days[common.G.SelectedDay]
		if day == "Sunday" {
			maxX++
		}
		if err := days.SetDayView(day, -1, -1, maxX, maxY-1); err != nil {
			return err
		}
		common.Views.Days[day].Frame = false
		width, height := common.Views.Days[day].Size()
		if common.SelectBlockForMod || common.SelectBlockForRemove {
			days_utils.SetDayViewSelectionContent(day, width, height)
		} else {
			days_utils.SetDayViewContent(day, width, height)
		}
	} else {
		for i, day := range common.Days {
			if err := days.SetDayView(day, i*maxX/7, 0, (i+1)*maxX/7, maxY-2); err != nil {
				return err
			}
		}
	}

	if _, err := common.G.SetCurrentView(common.CurrentViewOnTop); err != nil {
		return err
	}
	if _, err := common.G.SetViewOnTop(common.CurrentViewOnTop); err != nil {
		return err
	}

	if err := bar.SetBar(maxX, maxY); err != nil {
		return err
	}
	return nil
}
