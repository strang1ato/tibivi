package tibivi

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/oltarzewskik/gocui"
)

// currentDay returns number of current day of the week
func currentDay() int {
	day, _ := exec.Command("/bin/sh", "-c", "date +%w").Output()
	currentDay, _ := strconv.Atoi(string(day[:1]))
	if currentDay == 0 {
		currentDay = 6
	} else {
		currentDay--
	}
	return currentDay
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
	whiteSpaces := width - (len(b.startHour) + len(b.endHour) + 7)
	var line string
	for i := 0; i < whiteSpaces; i++ {
		line += " "
	}
	line += b.startHour + ":" + b.startMinute + "-" + b.endHour + ":" + b.endMinute
	return line
}

// setDayViewContent sets content of day of the week view
func (tbv *Tibivi) setDayViewContent(v *gocui.View, width, height int) {
	v.Clear()
	var blockLength int
	for _, b := range tbv.Schedule[v.Name()] {
		fmt.Fprintln(v, b.content)
		fmt.Fprint(v, "\x1b[33m"+newTimeLine(b, width)+"\x1b[0m")
		fmt.Fprint(v, newSeparator(width))
		blockLength += len(b.content)/width + 3
	}
	freeSpace := height - blockLength
	fmt.Fprint(v, "\x1b[33m")
	for i := 0; i < freeSpace; i++ {
		fmt.Fprintln(v, "~")
	}
	fmt.Fprint(v, "\x1b[0m")
}

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

// previousDayView goes to previous day of the week view
func (tbv *Tibivi) previousDayView(g *gocui.Gui, v *gocui.View) error {
	previousIndex := tbv.selectedDay - 1
	if previousIndex < 0 {
		previousIndex = 6
	}
	name := tbv.days[previousIndex]
	if _, err := tbv.setCurrentViewOnTop(name); err != nil {
		return err
	}

	tbv.selectedDay = previousIndex
	return nil
}

// nextDayView goes to next day of the week view
func (tbv *Tibivi) nextDayView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := tbv.selectedDay + 1
	if nextIndex > 6 {
		nextIndex = 0
	}
	name := tbv.days[nextIndex]
	if _, err := tbv.setCurrentViewOnTop(name); err != nil {
		return err
	}

	tbv.selectedDay = nextIndex
	return nil
}
