package tibivi

import (
	"fmt"
	"strings"

	"github.com/oltarzewskik/tibivi-gocui"
)

// setBar setups "bar" and ":" views
func (tbv *Tibivi) setBar(maxX, maxY int) error {
	if v, err := tbv.g.SetView(":", -1, maxY-2, 1, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Visible = false

		fmt.Fprint(v, ":")

		tbv.Views.bar[":"] = v
	}

	if v, err := tbv.g.SetView("bar", 0, maxY-2, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Frame = false

		tbv.Views.bar["bar"] = v
	}
	return nil
}

// focusBar focuses on bar view
func (tbv *Tibivi) focusBar(g *gocui.Gui, v *gocui.View) error {
	tbv.g.Cursor = true

	tbv.Views.bar[":"].Visible = true

	tbv.Views.bar["bar"].Clear()
	tbv.Views.bar["bar"].SetCursor(0, 0)

	tbv.Views.currentViewOnTop = "bar"
	return nil
}

// unfocusBar unfocuses from bar view
func (tbv *Tibivi) unfocusBar(g *gocui.Gui, v *gocui.View) error {
	tbv.g.Cursor = false

	tbv.Views.bar[":"].Visible = false

	tbv.Views.bar["bar"].Clear()

	tbv.Views.currentViewOnTop = tbv.days[tbv.g.SelectedDay]
	return nil
}

// executeCommand executes command defined in command bar
func (tbv *Tibivi) executeCommand(g *gocui.Gui, v *gocui.View) error {
	command := tbv.Views.bar["bar"].Buffer()

	tbv.unfocusBar(g, v)

	var notEmpty bool
	if len(command) > 0 {
		command = strings.TrimSuffix(command, "\n")
		notEmpty = true
	}

	switch command {
	case "q", "q!":
		return tbv.quitIgnore(g, v)
	default:
		if notEmpty {
			fmt.Fprint(tbv.Views.bar["bar"], "\x1b[41m"+"Not a tibivi command: "+command+"\x1b[0m")
		}
	}
	return nil
}
