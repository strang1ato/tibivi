package tibivi

import (
	"errors"
	"strings"

	"github.com/oltarzewskik/tibivi-gocui"
)

// addBlock adds new block to `tbv.Schedule`
func (tbv *Tibivi) addBlock(startTime, finishTime, description string) error {
	startTime = strings.TrimSuffix(startTime, "\n")
	finishTime = strings.TrimSuffix(finishTime, "\n")
	description = strings.TrimSuffix(description, "\n")

	block := &Block{}
	blockFields := []*string{&block.startHour, &block.startMinute, &block.finishHour, &block.finishMinute, &block.content}

	var section int
	for _, char := range startTime {
		if char == ':' || char == '-' {
			section++
			continue
		}
		*blockFields[section] += string(char)
	}
	if section != 1 {
		return errors.New("Start time is invalid")
	} else {
		section++
	}
	for _, char := range finishTime {
		if char == ':' || char == '-' {
			section++
			continue
		}
		*blockFields[section] += string(char)
	}
	if section != 3 {
		return errors.New("Finish time is invalid")
	} else {
		if len(description) != 0 {
			*blockFields[4] = description
		} else {
			return errors.New("Description is invalid")
		}
	}

	block, err := addNumTimes(block)
	if err != nil {
		return err
	}
	day := append(tbv.Schedule[tbv.days[tbv.g.SelectedDay]], block)
	tbv.Schedule[tbv.days[tbv.g.SelectedDay]] = sortDay(day)
	return nil
}

// quitIgnore exits tibivi without saving current state
func (tbv *Tibivi) quitIgnore(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
