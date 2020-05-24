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
	G                    = g
	Views                = views.SetViews()
	Schedule             = datatypes.Schedule{}
	Days                 = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	DotTibivi            = os.Getenv("HOME") + "/.tibivi/"
	CurrentDay           = commands.CurrentDay()
	CurrentTime          = commands.CurrentTime()
	CurrentViewOnTop     = Days[CurrentDay]
	UpdatedDays          = make(map[string]bool)
	SelectedBlock        int
	SelectBlockForMod    bool
	SelectBlockForRemove bool
)

var g, _ = gocui.NewGui(gocui.OutputNormal)
