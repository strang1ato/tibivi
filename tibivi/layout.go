package tibivi

import (
	"time"

	"github.com/oltarzewskik/gocui"
)

// layout setups CUI layout
func (tbv *Tibivi) layout(g *gocui.Gui) error {
	maxX, maxY := tbv.g.Size()
	for i, day := range tbv.days {
		if err := tbv.setDayView(day, i*maxX/7, (i+1)*maxX/7, maxY-2); err != nil {
			return err
		}
	}

	if _, err := tbv.setCurrentViewOnTop(tbv.days[tbv.selectedDay]); err != nil {
		return err
	}

	if err := tbv.setBar(maxX, maxY); err != nil {
		return err
	}

	// Updates layout on resize
	go func() {
		lastGWidth, _ := tbv.g.Size()
		for {
			time.Sleep(500 * time.Millisecond)
			gWidth, _ := tbv.g.Size()
			if lastGWidth != gWidth {
				tbv.g.Update(func(g *gocui.Gui) error {
					for _, day := range tbv.Views.days {
						width, height := day.Size()
						tbv.setDayViewContent(day, width, height)
					}
					return nil
				})
				lastGWidth = gWidth
			}
		}
	}()
	return nil
}
