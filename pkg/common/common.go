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
	CurrentTime          float32
	CurrentViewOnTop     string
	UpdatedDays          map[string]bool
	SelectedBlock        int
	SelectBlockForMod    bool
	SelectBlockForRemove bool
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
	currentTime, err := commands.CurrentTime()
	if err != nil {
		return err
	}

	G = g
	Views = views.SetViews()
	Schedule = datatypes.Schedule{}
	Days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	DotTibivi = home + "/.tibivi/"
	CurrentDay = currentDay
	CurrentTime = currentTime
	CurrentViewOnTop = Days[CurrentDay]
	UpdatedDays = make(map[string]bool)
	return nil
}
