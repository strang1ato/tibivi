package days

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/data"
	"github.com/oltarzewskik/tibivi/pkg/days/utils"
	"github.com/oltarzewskik/tibivi/pkg/layout/utils"
)

// SetDayView setups day of the week view
func SetDayView(day string, x0, x1, y1 int) error {
	if day == "Sunday" {
		x1--
	}
	if v, err := common.G.SetView(day, x0, 0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = day
		v.Wrap = true

		width, height := v.Size()
		days_utils.SetDayViewContent(v, width, height)

		common.Views.Days[day] = v
	}
	return nil
}

// nextDayView goes to next day view
func nextDayView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := common.G.SelectedDay + 1
	if nextIndex > 6 {
		nextIndex = 0
	}
	common.CurrentViewOnTop = common.Days[nextIndex]

	common.G.SelectedDay = nextIndex

	if common.SelectBlockForRemove || common.SelectBlockForMod {
		moveBlockSelection()
	}
	return nil
}

// previousDayView goes to previous day view
func previousDayView(g *gocui.Gui, v *gocui.View) error {
	previousIndex := common.G.SelectedDay - 1
	if previousIndex < 0 {
		previousIndex = 6
	}
	common.CurrentViewOnTop = common.Days[previousIndex]

	common.G.SelectedDay = previousIndex

	if common.SelectBlockForRemove || common.SelectBlockForMod {
		moveBlockSelection()
	}
	return nil
}

func moveBlockSelection() {
	dayLen := len(data.Schedule[common.Days[common.G.SelectedDay]])
	if common.SelectedBlock >= dayLen && dayLen != 0 {
		common.SelectedBlock = len(data.Schedule[common.Days[common.G.SelectedDay]]) - 1
	}
	layout_utils.UpdateLayout()
}
