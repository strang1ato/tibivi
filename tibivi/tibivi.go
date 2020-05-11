package tibivi

import (
	"os"

	"github.com/oltarzewskik/tibivi-gocui"
)

// Tibivi struct is tibivi project wrapper
type Tibivi struct {
	g           *gocui.Gui
	Views       *Views
	Schedule    Schedule
	days        []string
	dotTibivi   string
	currentTime float32
	currentDay  int
}

// newTibivi returns new Tibivi object
func newTibivi() *Tibivi {
	tbv := &Tibivi{
		Views: &Views{
			bar:            map[string]*gocui.View{},
			days:           map[string]*gocui.View{},
			menu:           map[string]*gocui.View{},
			addBlockFields: []string{"addBlockDescription", "addBlockStartTime", "addBlockFinishTime"},
		},
		Schedule:    Schedule{},
		days:        []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"},
		dotTibivi:   os.Getenv("HOME") + "/.tibivi/",
		currentDay:  currentDay(),
		currentTime: currentTime(),
	}
	tbv.Views.currentViewOnTop = tbv.days[tbv.currentDay]
	return tbv
}
