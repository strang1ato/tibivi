package block

import (
	"errors"
	"strings"

	"github.com/oltarzewskik/tibivi-gocui"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"github.com/oltarzewskik/tibivi/pkg/datatypes"
	"github.com/oltarzewskik/tibivi/pkg/schedule"
)

var addingBlock bool

// SetAddBlockForm shows user view based add block form
func SetAddBlockForm() error {
	maxX, maxY := common.G.Size()
	if err := setForm(maxX, maxY); err != nil {
		return err
	}

	selectedFormField = 0
	common.CurrentViewOnTop = formFields[selectedFormField]

	addingBlock = true
	return nil
}

// submitAddBlockForm submits addBlock form
func submitAddBlockForm(g *gocui.Gui, v *gocui.View) error {
	if addingBlock {
		submitForm(addBlock, g, v)
	}
	return nil
}

// addBlock adds new block to schedule
func addBlock(startTime, finishTime, Description string) error {
	block, err := createBlock(startTime, finishTime, Description)
	if err != nil {
		return err
	}
	day := common.Schedule[common.Days[common.G.SelectedDay]]
	common.Schedule[common.Days[common.G.SelectedDay]] = schedule.SortDay(append(day, block))

	common.UpdatedDays[common.Days[common.G.SelectedDay]] = true
	return nil
}

// createBlock create, test and return new block
func createBlock(startTime, finishTime, Description string) (*datatypes.Block, error) {
	startTime = strings.TrimSuffix(startTime, "\n")
	finishTime = strings.TrimSuffix(finishTime, "\n")
	Description = strings.TrimSuffix(Description, "\n")

	block := &datatypes.Block{}
	blockFields := []*string{&block.StartHour, &block.StartMinute, &block.FinishHour, &block.FinishMinute, &block.Description}

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
	if len(Description) == 0 {
		return nil, errors.New("Description is invalid")
	}
	*blockFields[4] = Description

	block, err := schedule.AddNumTimes(block)
	if err != nil {
		return nil, err
	}
	return block, nil
}
