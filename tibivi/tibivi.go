package tibivi

import (
	"os"

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
			days: map[string]*gocui.View{},
		},
	}
	return tbv
}

// Run runs Tibivi
func Run() error {
	tbv := newTibivi()

	if err := tbv.createDotTibivi(); err != nil {
		return err
	}

	if err := tbv.setSchedule(); err != nil {
		return err
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	tbv.g = g
	defer tbv.g.Close()

	tbv.g.Highlight = true
	tbv.g.SelFgColor = gocui.ColorGreen

	tbv.g.SetManagerFunc(tbv.layout)

	if err := tbv.keybindings(); err != nil {
		return err
	}

	if err := tbv.g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}
	return nil
}
