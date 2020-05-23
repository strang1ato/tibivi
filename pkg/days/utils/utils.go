package days_utils

import (
	"fmt"

	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/datatypes"
)

// SetDayViewContent sets content of day view
func SetDayViewContent(v *gocui.View, width, height int) {
	v.Clear()
	var blockLength int
	for _, b := range common.Schedule[v.Name()] {
		if v.Name() == common.Days[common.CurrentDay] && (common.CurrentTime >= b.NumStartTime && common.CurrentTime < b.NumFinishTime) {
			fmt.Fprintln(v, "\x1b[31m"+b.Description+"\x1b[0m")
		} else {
			fmt.Fprintln(v, b.Description)
		}
		fmt.Fprint(v, "\x1b[33m"+newTimeLine(b, width)+"\x1b[0m")
		fmt.Fprint(v, newSeparator(width))
		blockLength += len(b.Description)/width + 3
	}
	freeSpace := height - blockLength
	fmt.Fprint(v, "\x1b[33m")
	for i := 0; i < freeSpace; i++ {
		fmt.Fprintln(v, "~")
	}
	fmt.Fprint(v, "\x1b[0m")
}

// SetDayViewSelectionContent sets content of day view with higlighting of selected block for remove/modification
func SetDayViewSelectionContent(v *gocui.View, width, height int) {
	v.Clear()
	var blockLength int
	for i, b := range common.Schedule[v.Name()] {
		if v.Name() == common.Days[common.G.SelectedDay] && common.SelectedBlock == i {
			fmt.Fprint(v, "\x1b[7m")
			fmt.Fprint(v, b.Description)
			whiteSpaces := width - len(b.Description)%width
			for i := 0; i < whiteSpaces; i++ {
				fmt.Fprint(v, " ")
			}
			fmt.Fprint(v, newTimeLine(b, width))
			fmt.Fprint(v, "\x1b[0m")
		} else {
			fmt.Fprintln(v, b.Description)
			fmt.Fprint(v, "\x1b[33m"+newTimeLine(b, width)+"\x1b[0m")
		}
		fmt.Fprint(v, newSeparator(width))
		blockLength += len(b.Description)/width + 3
	}
	freeSpace := height - blockLength
	fmt.Fprint(v, "\x1b[33m")
	for i := 0; i < freeSpace; i++ {
		fmt.Fprintln(v, "~")
	}
	fmt.Fprint(v, "\x1b[0m")
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
func newTimeLine(b *datatypes.Block, width int) string {
	whiteSpaces := width - (len(b.StartHour) + len(b.FinishHour) + 7)
	var line string
	for i := 0; i < whiteSpaces; i++ {
		line += " "
	}
	line += b.StartHour + ":" + b.StartMinute + "-" + b.FinishHour + ":" + b.FinishMinute
	return line
}
