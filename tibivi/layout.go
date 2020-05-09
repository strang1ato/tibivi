package tibivi

import (
	"time"

	"github.com/oltarzewskik/tibivi-gocui"
)

// layout setups CUI layout
func (tbv *Tibivi) layout(g *gocui.Gui) error {
	maxX, maxY := tbv.g.Size()
	for i, day := range tbv.days {
		if err := tbv.setDayView(day, i*maxX/7, (i+1)*maxX/7, maxY-2); err != nil {
			return err
		}
	}

	if err := tbv.setCurrentViewOnTop(tbv.Views.currentViewOnTop); err != nil {
		return err
	}

	if err := tbv.setBar(maxX, maxY); err != nil {
		return err
	}
	return nil
}

// updateLayoutOnCurrentBlockChange when run in goroutine updates layout if current time block changed
func (tbv *Tibivi) updateLayoutOnCurrentBlockChange() {
	for {
		time.Sleep(time.Minute)

		tbv.currentTime = tbv.currentTime + float32(1)/60
		if tbv.currentTime >= 24 {
			tbv.currentTime = float32(0)
			tbv.currentDay++
			if tbv.currentDay > 6 {
				tbv.currentDay = 0
			}
		}

		tbv.updateLayout()
	}
}

// updateLayoutOnResize when run in goroutine updates layout on resize
func (tbv *Tibivi) updateLayoutOnResize() {
	lastMaxX, _ := tbv.g.Size()
	for {
		time.Sleep(500 * time.Millisecond)

		maxX, _ := tbv.g.Size()
		if lastMaxX != maxX {
			tbv.updateLayout()
			lastMaxX = maxX
		}
	}
}

// updateLayout updates layout
func (tbv *Tibivi) updateLayout() {
	tbv.g.Update(func(g *gocui.Gui) error {
		for _, day := range tbv.Views.days {
			width, height := day.Size()
			tbv.setDayViewContent(day, width, height)
		}
		return nil
	})
}
