package tibivi

import (
	"errors"
	"strings"

	"github.com/oltarzewskik/tibivi-gocui"
)

// addBlock adds new block to `tbv.Schedule`
func (tbv *Tibivi) addBlock(startTime, finishTime, description string) error {
	block, err := tbv.createBlock(startTime, finishTime, description)
	if err != nil {
		return err
	}
	day := tbv.Schedule[tbv.days[tbv.g.SelectedDay]]
	tbv.Schedule[tbv.days[tbv.g.SelectedDay]] = sortDay(append(day, block))
	return nil
}

// quitIgnore exits tibivi without saving current state
func (tbv *Tibivi) quitIgnore(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// createBlock create, test and return block
func (tbv *Tibivi) createBlock(startTime, finishTime, description string) (*Block, error) {
	startTime = strings.TrimSuffix(startTime, "\n")
	finishTime = strings.TrimSuffix(finishTime, "\n")
	description = strings.TrimSuffix(description, "\n")

	block := &Block{}
	blockFields := []*string{&block.startHour, &block.startMinute, &block.finishHour, &block.finishMinute, &block.description}

	var section int
	for _, char := range startTime {
		if char == ':' || char == '-' {
			section++
			continue
		}
		*blockFields[section] += string(char)
	}
	if section != 1 {
		return nil, errors.New("Start time is invalid")
	}
	section++
	for _, char := range finishTime {
		if char == ':' || char == '-' {
			section++
			continue
		}
		*blockFields[section] += string(char)
	}
	if section != 3 {
		return nil, errors.New("Finish time is invalid")
	}
	if len(description) == 0 {
		return nil, errors.New("Description is invalid")
	}
	*blockFields[4] = description

	block, err := addNumTimes(block)
	if err != nil {
		return nil, err
	}
	return block, nil
}
