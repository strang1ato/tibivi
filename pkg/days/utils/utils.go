package days_utils

import (
	"fmt"
	"unicode/utf8"

	"github.com/strang1ato/tibivi/pkg/common"
	"github.com/strang1ato/tibivi/pkg/datatypes"
)

// SetDayViewContent sets content of day view
func SetDayViewContent(day string, width, height int) {
	v := common.Views.Days[day]
	v.Clear()
	freeSpace := height
	common.BlocksInBuffer[day] = 0
	for _, b := range common.Schedule[v.Name()][common.Shift[day]:] {
		if v.Name() == common.Days[common.CurrentDay] &&
			((common.CurrentHour > b.NumStartHour || (common.CurrentHour == b.NumStartHour && common.CurrentMinute >= b.NumStartMinute)) &&
				(common.CurrentHour < b.NumFinishHour || (common.CurrentHour == b.NumFinishHour && common.CurrentMinute < b.NumFinishMinute))) {
			fmt.Fprintln(v, "\x1b[31m"+b.Description+"\x1b[0m")
		} else {
			fmt.Fprintln(v, b.Description)
		}
		fmt.Fprint(v, "\x1b[33m"+newTimeLine(b, width)+"\x1b[0m")
		fmt.Fprint(v, newSeparator(width))
		freeSpace -= utf8.RuneCountInString(b.Description)/width + 3
		if freeSpace >= 0 {
			common.BlocksInBuffer[day]++
		}
	}
	fmt.Fprint(v, "\x1b[33m")
	for i := 0; i < freeSpace; i++ {
		fmt.Fprintln(v, "~")
	}
	fmt.Fprint(v, "\x1b[0m")
}

// SetDayViewSelectionContent sets content of day view with higlighting of selected block for remove/modification
func SetDayViewSelectionContent(day string, width, height int) {
	v := common.Views.Days[day]
	v.Clear()
	freeSpace := height
	common.BlocksInBuffer[day] = 0
	for i, b := range common.Schedule[v.Name()][common.Shift[day]:] {
		if v.Name() == common.Days[common.G.SelectedDay] && common.SelectedBlock == i {
			fmt.Fprint(v, "\x1b[7m")
			fmt.Fprint(v, b.Description)
			whiteSpaces := width - utf8.RuneCountInString(b.Description)%width
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
		freeSpace -= utf8.RuneCountInString(b.Description)/width + 3
		if freeSpace >= 0 {
			common.BlocksInBuffer[day]++
		}
	}
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
