package tibivi

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/oltarzewskik/gocui"
)

// Tibivi struct is tibivi project wrapper
type Tibivi struct {
	g           *gocui.Gui
	currentDay  int
	selectedDay int
	days        []string
	dotTibivi   string
	Schedule    Schedule
	Views       *Views
}

// newTibivi returns new Tibivi object
func newTibivi() *Tibivi {
	tbv := &Tibivi{
		currentDay:  currentDay(),
		selectedDay: currentDay(),
		days:        []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"},
		dotTibivi:   os.Getenv("HOME") + "/.tibivi/",
		Schedule:    Schedule{},
		Views: &Views{
			bar:  map[string]*gocui.View{},
			days: map[string]*gocui.View{},
		},
	}
	tbv.Views.currentViewOnTop = tbv.days[tbv.selectedDay]
	return tbv
}

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
