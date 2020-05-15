package tibivi

import (
	"fmt"

	"github.com/oltarzewskik/tibivi-gocui"
)

// setDayView setups day of the week view
func (tbv *Tibivi) setDayView(day string, x0, x1, y1 int) error {
	if day == "Sunday" {
		x1--
	}
	if v, err := tbv.g.SetView(day, x0, 0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = day
		v.Wrap = true

		width, height := v.Size()
		tbv.setDayViewContent(v, width, height)

		tbv.Views.days[day] = v
	}
	return nil
}

// setDayViewContent sets content of day view
func (tbv *Tibivi) setDayViewContent(v *gocui.View, width, height int) {
	v.Clear()
	var blockLength int
	for _, b := range tbv.Schedule[v.Name()] {
		if v.Name() == tbv.days[tbv.currentDay] && (tbv.currentTime >= b.numStartTime && tbv.currentTime < b.numFinishTime) {
			fmt.Fprintln(v, "\x1b[31m"+b.description+"\x1b[0m")
		} else {
			fmt.Fprintln(v, b.description)
		}
		fmt.Fprint(v, "\x1b[33m"+newTimeLine(b, width)+"\x1b[0m")
		fmt.Fprint(v, newSeparator(width))
		blockLength += len(b.description)/width + 3
	}
	freeSpace := height - blockLength
	fmt.Fprint(v, "\x1b[33m")
	for i := 0; i < freeSpace; i++ {
		fmt.Fprintln(v, "~")
	}
	fmt.Fprint(v, "\x1b[0m")
}

// setDayViewSelectionContent sets content of day view with higlighting of selected block for remove/modification
func (tbv *Tibivi) setDayViewSelectionContent(v *gocui.View, width, height int) {
	v.Clear()
	var blockLength int
	for i, b := range tbv.Schedule[v.Name()] {
		if v.Name() == tbv.days[tbv.g.SelectedDay] && tbv.BlockUtils.selectedBlock == i {
			fmt.Fprint(v, "\x1b[7m")
			fmt.Fprint(v, b.description)
			whiteSpaces := width - len(b.description)%width
			for i := 0; i < whiteSpaces; i++ {
				fmt.Fprint(v, " ")
			}
			fmt.Fprint(v, newTimeLine(b, width))
			fmt.Fprint(v, "\x1b[0m")
		} else {
			fmt.Fprintln(v, b.description)
			fmt.Fprint(v, "\x1b[33m"+newTimeLine(b, width)+"\x1b[0m")
		}
		fmt.Fprint(v, newSeparator(width))
		blockLength += len(b.description)/width + 3
	}
	freeSpace := height - blockLength
	fmt.Fprint(v, "\x1b[33m")
	for i := 0; i < freeSpace; i++ {
		fmt.Fprintln(v, "~")
	}
	fmt.Fprint(v, "\x1b[0m")
}

// nextDayView goes to next day view
func (tbv *Tibivi) nextDayView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := tbv.g.SelectedDay + 1
	if nextIndex > 6 {
		nextIndex = 0
	}
	tbv.Views.currentViewOnTop = tbv.days[nextIndex]

	tbv.g.SelectedDay = nextIndex

	if tbv.BlockUtils.selectedForRemove || tbv.BlockUtils.selectedForMod {
		tbv.moveBlockSelection()
	}
	return nil
}

// previousDayView goes to previous day view
func (tbv *Tibivi) previousDayView(g *gocui.Gui, v *gocui.View) error {
	previousIndex := tbv.g.SelectedDay - 1
	if previousIndex < 0 {
		previousIndex = 6
	}
	tbv.Views.currentViewOnTop = tbv.days[previousIndex]

	tbv.g.SelectedDay = previousIndex

	if tbv.BlockUtils.selectedForRemove || tbv.BlockUtils.selectedForMod {
		tbv.moveBlockSelection()
	}
	return nil
}

// newSeparator returns string separator of given width
func newSeparator(width int) string {
	var separator string
	for i := 0; i < width; i++ {
		separator += "-"
	}
	return separator
}

// newTimeLine returns string line with time of given width
func newTimeLine(b *Block, width int) string {
	whiteSpaces := width - (len(b.startHour) + len(b.finishHour) + 7)
	var line string
	for i := 0; i < whiteSpaces; i++ {
		line += " "
	}
	line += b.startHour + ":" + b.startMinute + "-" + b.finishHour + ":" + b.finishMinute
	return line
}

func (tbv *Tibivi) moveBlockSelection() {
	dayLen := len(tbv.Schedule[tbv.days[tbv.g.SelectedDay]])
	if tbv.BlockUtils.selectedBlock >= dayLen && dayLen != 0 {
		tbv.BlockUtils.selectedBlock = len(tbv.Schedule[tbv.days[tbv.g.SelectedDay]]) - 1
	}
	tbv.updateLayout()
}
