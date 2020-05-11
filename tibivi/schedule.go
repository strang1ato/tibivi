package tibivi

import (
	"errors"
	"io/ioutil"
	"strconv"
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
	description string
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

// read converts data defined in datafile and returns it
func (tbv *Tibivi) read(filename string) (Day, error) {
	byteData, err := ioutil.ReadFile(tbv.dotTibivi + filename)
	if err != nil {
		return nil, err
	}
	data := strings.Split(string(byteData), "\n")
	for i, entry := range data {
		if len(entry) == 0 {
			data = append(data[:i], data[i+1:]...)
		}
	}

	day := make(Day, len(data))
	for i, entry := range data {
		day[i] = &Block{}
		blockFields := []*string{&day[i].startHour, &day[i].startMinute, &day[i].finishHour, &day[i].finishMinute, &day[i].description}
		var section int
		for _, char := range entry {
			if section != 4 && (char == ':' || char == ' ' || char == '-') {
				section++
				continue
			}
			*blockFields[section] += string(char)
		}

		day[i], err = addNumTimes(day[i])
		if err != nil {
			return nil, err
		}
	}

	day = sortDay(day)
	return day, nil
}

// addNumTimes adds numeric times and check if there are no defects in block
func addNumTimes(b *Block) (*Block, error) {
	startHour, err := strconv.Atoi(b.startHour)
	if err != nil {
		return nil, err
	}
	startMinute, err := strconv.Atoi(b.startMinute)
	if err != nil {
		return nil, err
	}
	finishHour, err := strconv.Atoi(b.finishHour)
	if err != nil {
		return nil, err
	}
	finishMinute, err := strconv.Atoi(b.finishMinute)
	if err != nil {
		return nil, err
	}

	if startHour > 23 || startHour < 0 {
		return nil, errors.New("Block start hour is wrong")
	}
	if startMinute > 60 || startMinute < 0 {
		return nil, errors.New("Block start minute is wrong")
	}
	if finishHour > 23 || finishHour < 0 {
		return nil, errors.New("Block end hour is wrong")
	}
	if finishMinute > 60 || finishMinute < 0 {
		return nil, errors.New("Block end minute is wrong")
	}

	numStartTime := float32(startHour) + float32(startMinute)/60
	numFinishTime := float32(finishHour) + float32(finishMinute)/60
	if numStartTime >= numFinishTime {
		return nil, errors.New("Block start time is equal or greater than end time")
	}
	b.numStartTime, b.numFinishTime = numStartTime, numFinishTime
	return b, nil
}

// sortDay sorts blocks by start time
func sortDay(day Day) Day {
	for i := 1; i < len(day); i++ {
		j := i
		for j > 0 {
			if day[j].numStartTime < day[j-1].numStartTime {
				day[j], day[j-1] = day[j-1], day[j]
			}
			j--
		}
	}
	return day
}
