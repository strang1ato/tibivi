package layout

import (
	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/bar"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/days"
)

// Layout setups CUI layout
func Layout(g *gocui.Gui) error {
	maxX, maxY := common.G.Size()
	for i, day := range common.Days {
		if err := days.SetDayView(day, i*maxX/7, (i+1)*maxX/7, maxY-2); err != nil {
			return err
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
