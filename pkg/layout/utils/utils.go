package layout_utils

import (
	"sync"
	"time"

	"github.com/strang1ato/tibivi-gocui"
	"github.com/strang1ato/tibivi/pkg/commands"
	"github.com/strang1ato/tibivi/pkg/common"
	"github.com/strang1ato/tibivi/pkg/days/utils"
)

// UpdateLayout updates layout
func UpdateLayout() {
	common.G.Update(func(g *gocui.Gui) error {
		if common.Focus {
			day := common.Days[common.G.SelectedDay]
			width, height := common.Views.Days[day].Size()
			if common.SelectBlockForMod || common.SelectBlockForRemove {
				days_utils.SetDayViewSelectionContent(day, width, height)
			} else {
				days_utils.SetDayViewContent(day, width, height)
			}
		} else {
			for _, day := range common.Days {
				width, height := common.Views.Days[day].Size()
				if common.SelectBlockForMod || common.SelectBlockForRemove {
					days_utils.SetDayViewSelectionContent(day, width, height)
				} else {
					days_utils.SetDayViewContent(day, width, height)
				}
			}
		}
		return nil
	})
}

var mutex = sync.Mutex{}

// UpdateLayoutOnCurrentBlockChange when run in goroutine updates layout if current time block changed
func UpdateLayoutOnCurrentBlockChange() {
	for {
		common.CurrentHour, common.CurrentMinute, _ = commands.CurrentTime()
		common.CurrentDay, _ = commands.CurrentDay()

		mutex.Lock()
		UpdateLayout()
		mutex.Unlock()

		time.Sleep(time.Second)
	}
}

// UpdateLayoutOnResize when run in goroutine updates layout on resize
func UpdateLayoutOnResize() {
	lastMaxX, _ := common.G.Size()
	for {
		time.Sleep(500 * time.Millisecond)

		maxX, _ := common.G.Size()
		if lastMaxX != maxX {
			mutex.Lock()
			UpdateLayout()
			mutex.Unlock()
			lastMaxX = maxX
		}
	}
}
