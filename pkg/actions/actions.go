package actions

import (
	"fmt"
	"os"

	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
)

// Quit closes tibivi if schedule wasn't changed
func Quit() error {
	for _, day := range common.Days {
		if common.UpdatedDays[day] {
			common.Views.Bar["bar"].Clear()
			message := "No write since last change (add ! to override)"
			fmt.Fprint(common.Views.Bar["bar"], "\x1b[41m"+message+"\x1b[0m")
			return nil
		}
	}
	return gocui.ErrQuit
}

// QuitIgnore closes tibivi without saving current state
func QuitIgnore(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// Write writes schedule to datafiles
func Write() error {
	for _, day := range common.Days {
		if common.UpdatedDays[day] {
			datafile, err := os.Create(common.DotTibivi + day + ".txt")
			if err != nil {
				return err
			}
			for _, b := range common.Schedule[day] {
				time := b.StartHour + ":" + b.StartMinute + "-" + b.FinishHour + ":" + b.FinishMinute
				_, err := datafile.WriteString(time + " " + b.Description + "\n")
				if err != nil {
					return err
				}
			}
			common.UpdatedDays[day] = false
		}
	}
	return nil
}

// WriteQuit writes schedule to datafiles and then closes tibivi
func WriteQuit() error {
	if err := Write(); err != nil {
		return err
	}
	return gocui.ErrQuit
}
