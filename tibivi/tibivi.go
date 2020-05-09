package tibivi

import (
	"os"

	"github.com/oltarzewskik/tibivi-gocui"
)

// Tibivi struct is tibivi project wrapper
type Tibivi struct {
	g           *gocui.Gui
	currentDay  int
	currentTime float32
	days        []string
	dotTibivi   string
	Schedule    Schedule
	Views       *Views
}

// newTibivi returns new Tibivi object
func newTibivi() *Tibivi {
	tbv := &Tibivi{
		currentDay:  currentDay(),
		currentTime: currentTime(),
		days:        []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"},
		dotTibivi:   os.Getenv("HOME") + "/.tibivi/",
		Schedule:    Schedule{},
		Views: &Views{
			bar:            map[string]*gocui.View{},
			days:           map[string]*gocui.View{},
			menu:           map[string]*gocui.View{},
			addBlockFields: []string{"addBlockDescription", "addBlockStartTime", "addBlockFinishTime"},
		},
	}
	tbv.Views.currentViewOnTop = tbv.days[tbv.currentDay]
	return tbv
}
