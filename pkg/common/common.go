package common

import (
	"os"

	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/commands"
	"github.com/oltarzewskik/tibivi/pkg/datatypes"
	"github.com/oltarzewskik/tibivi/pkg/views"
)

// Declare variables used by multiple packages
var (
	G                    *gocui.Gui
	Views                *views.Views
	Schedule             datatypes.Schedule
	Days                 []string
	DotTibivi            string
	CurrentDay           int
	CurrentHour          int
	CurrentMinute        int
	CurrentViewOnTop     string
	Focus                bool
	UpdatedDays          map[string]bool
	BlocksInBuffer       map[string]int
	SelectedBlock        int
	SelectBlockForMod    bool
	SelectBlockForRemove bool
	Shift                map[string]int
)

// SetCommonVars assigns values to common variables
func SetCommonVars() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return err
	}
	home := os.Getenv("HOME")
	currentDay, err := commands.CurrentDay()
	if err != nil {
		return err
	}
	currentHour, err := commands.CurrentHour()
	if err != nil {
		return err
	}
	currentMinute, err := commands.CurrentMinute()
	if err != nil {
		return err
	}

	G = g
	Views = views.SetViews()
	Schedule = datatypes.Schedule{}
	Days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	DotTibivi = home + "/.tibivi/"
	CurrentDay = currentDay
	CurrentHour = currentHour
	CurrentMinute = currentMinute
	CurrentViewOnTop = Days[CurrentDay]
	Focus = true
	UpdatedDays = make(map[string]bool)
	BlocksInBuffer = make(map[string]int)
	Shift = make(map[string]int)
	return nil
}
