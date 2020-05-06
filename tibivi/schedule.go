package tibivi

import (
	"errors"
	"strings"
)

// Schedule consists whole week schedule
type Schedule map[string]Day

// Day consists time blocks of day of the week
type Day []*Block

// Block consists fields of time block
type Block struct {
	startHour, startMinute,
	finishHour, finishMinute,
	content string
	numStartTime, numFinishTime float32
}

// setSchedule supplies `tbv.Schedule` with data from datafiles
func (tbv *Tibivi) setSchedule() error {
	for _, day := range tbv.days {
		data, err := tbv.read(day + ".txt")
		if err != nil {
			return err
		}
		tbv.Schedule[day] = data
	}
	return nil
}

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
	block, err := tbv.addNumTimes(block)
	if err != nil {
		return err
	}
	day := append(tbv.Schedule[tbv.days[tbv.g.SelectedDay]], block)
	tbv.Schedule[tbv.days[tbv.g.SelectedDay]] = tbv.sortDay(day)
	return nil
}
