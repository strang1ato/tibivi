package days

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/days/utils"
	"github.com/oltarzewskik/tibivi/pkg/layout/utils"
)

// SetDayView setups day of the week view
func SetDayView(day string, x0, y0, x1, y1 int) error {
	if day == "Sunday" {
		x1--
	}
	if v, err := common.G.SetView(day, x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = day
		v.Wrap = true

		width, height := v.Size()
		common.Views.Days[day] = v
		days_utils.SetDayViewContent(day, width, height)
	}
	return nil
}

// switchDayViewFocus switches focus
func switchDayViewFocus(g *gocui.Gui, v *gocui.View) error {
	if common.Focus {
		unFocusDayView(g, v)
	} else {
		common.Focus = true
		layout_utils.UpdateLayout()
	}

	return nil
}

// unFocusDayView unfocuses from day view to whole week schedule
func unFocusDayView(g *gocui.Gui, v *gocui.View) error {
	if common.SelectBlockForRemove || common.SelectBlockForMod {
		return nil
	}
	if common.Focus {
		common.Focus = false
		for _, day := range common.Views.Days {
			day.Frame = true
		}
		layout_utils.UpdateLayout()
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

func shiftBlock(g *gocui.Gui, v *gocui.View) error {
	currentDay := common.Days[common.G.SelectedDay]
	dayLen := len(common.Schedule[currentDay])
	if !(common.SelectBlockForRemove || common.SelectBlockForMod) || common.SelectedBlock >= common.BlocksInBuffer[currentDay]-1 {
		if common.Shift[currentDay] < dayLen-common.BlocksInBuffer[common.Days[common.G.SelectedDay]] {
			common.Shift[currentDay]++
			layout_utils.UpdateLayout()
		}
	}
	return nil
}

func unShiftBlock(g *gocui.Gui, v *gocui.View) error {
	if !(common.SelectBlockForRemove || common.SelectBlockForMod) || common.SelectedBlock <= 0 {
		currentDay := common.Days[common.G.SelectedDay]
		if common.Shift[currentDay] > 0 {
			common.Shift[currentDay]--
			layout_utils.UpdateLayout()
		}
	}
	return nil
}

func moveBlockSelection() {
	blocksInBuffer := common.BlocksInBuffer[common.Days[common.G.SelectedDay]]
	if common.SelectedBlock >= blocksInBuffer && blocksInBuffer != 0 {
		common.SelectedBlock = len(common.Schedule[common.Days[common.G.SelectedDay]]) - 1
	}
	layout_utils.UpdateLayout()
}
