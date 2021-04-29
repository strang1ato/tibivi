package bar

import (
	"fmt"
	"strings"

	"github.com/strang1ato/tibivi-gocui"
	"github.com/strang1ato/tibivi/pkg/actions"
	"github.com/strang1ato/tibivi/pkg/common"
)

// SetBar setups "bar" and ":" views
func SetBar(maxX, maxY int) error {
	if v, err := common.G.SetView(":", -1, maxY-2, 1, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Visible = false

		fmt.Fprint(v, ":")

		common.Views.Bar[":"] = v
	}

	if v, err := common.G.SetView("bar", 0, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Frame = false

		common.Views.Bar["bar"] = v
	}
	return nil
}

// focusBar focuses on bar view
func focusBar(g *gocui.Gui, v *gocui.View) error {
	common.G.Cursor = true

	common.Views.Bar[":"].Visible = true

	common.Views.Bar["bar"].Clear()
	common.Views.Bar["bar"].SetCursor(0, 0)

	common.CurrentViewOnTop = "bar"
	return nil
}

// unFocusBar unfocuses from bar view
func unFocusBar(g *gocui.Gui, v *gocui.View) error {
	common.G.Cursor = false

	common.Views.Bar[":"].Visible = false

	common.Views.Bar["bar"].Clear()

	common.CurrentViewOnTop = common.Days[common.G.SelectedDay]
	return nil
}

// executeCommand executes command defined in command bar
func executeCommand(g *gocui.Gui, v *gocui.View) error {
	command := common.Views.Bar["bar"].Buffer()

	unFocusBar(g, v)

	var notEmpty bool
	if len(command) > 0 {
		command = strings.TrimSuffix(command, "\n")
		notEmpty = true
	}

	switch command {
	case "q":
		return actions.Quit()
	case "q!":
		return actions.QuitIgnore(g, v)
	case "w", "w!":
		return actions.Write()
	case "wq", "wq!":
		return actions.WriteQuit()
	case "day":
		fmt.Fprint(common.Views.Bar["bar"], actions.SelectedDay())
	default:
		if notEmpty {
			fmt.Fprint(common.Views.Bar["bar"], "\x1b[41m"+"Not a tibivi command: "+command+"\x1b[0m")
		}
	}
	return nil
}
