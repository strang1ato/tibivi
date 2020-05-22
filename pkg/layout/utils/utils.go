package layout_utils

import (
	"time"

	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/days/utils"
)

// UpdateLayout updates layout
func UpdateLayout() {
	common.G.Update(func(g *gocui.Gui) error {
		for _, day := range common.Views.Days {
			width, height := day.Size()
			if common.SelectBlockForMod || common.SelectBlockForRemove {
				days_utils.SetDayViewSelectionContent(day, width, height)
			} else {
				days_utils.SetDayViewContent(day, width, height)
			}
		}
		return nil
	})
}

// UpdateLayoutOnCurrentBlockChange when run in goroutine updates layout if current time block changed
func UpdateLayoutOnCurrentBlockChange() {
	for {
		time.Sleep(time.Minute)

		common.CurrentTime = common.CurrentTime + float32(1)/60
		if common.CurrentTime >= 24 {
			common.CurrentTime = float32(0)
			common.CurrentDay++
			if common.CurrentDay > 6 {
				common.CurrentDay = 0
			}
		}

		UpdateLayout()
	}
}

// UpdateLayoutOnResize when run in goroutine updates layout on resize
func UpdateLayoutOnResize() {
	lastMaxX, _ := common.G.Size()
	for {
		time.Sleep(500 * time.Millisecond)

		maxX, _ := common.G.Size()
		if lastMaxX != maxX {
			UpdateLayout()
			lastMaxX = maxX
		}
	}
}
