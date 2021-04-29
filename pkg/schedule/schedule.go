package schedule

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/strang1ato/tibivi/pkg/common"
	"github.com/strang1ato/tibivi/pkg/datatypes"
)

// SetSchedule supplies `Schedule` with data from datafiles
func SetSchedule() error {
	for _, day := range common.Days {
		data, err := read(day + ".txt")
		if err != nil {
			return err
		}
		common.Schedule[day] = data
	}
	return nil
}

// read converts data defined in datafile and returns it
func read(filename string) (datatypes.Day, error) {
	byteData, err := ioutil.ReadFile(common.DotTibivi + filename)
	if err != nil {
		return nil, err
	}
	data := strings.Split(string(byteData), "\n")
	for i, entry := range data {
		if len(entry) == 0 {
			data = append(data[:i], data[i+1:]...)
		}
	}

	day := make(datatypes.Day, len(data))
	for i, entry := range data {
		day[i] = &datatypes.Block{}
		blockFields := []*string{&day[i].StartHour, &day[i].StartMinute, &day[i].FinishHour, &day[i].FinishMinute, &day[i].Description}
		var section int
		for _, char := range entry {
			if section != 4 && (char == ':' || char == ' ' || char == '-') {
				section++
				continue
			}
			*blockFields[section] += string(char)
		}

		day[i], err = AddNumTimes(day[i])
		if err != nil {
			return nil, err
		}
	}

	day = SortDay(day)
	return day, nil
}

// AddNumTimes adds numeric times and check if there are no defects in block
func AddNumTimes(b *datatypes.Block) (*datatypes.Block, error) {
	StartHour, err := strconv.Atoi(b.StartHour)
	if err != nil {
		return nil, err
	}
	StartMinute, err := strconv.Atoi(b.StartMinute)
	if err != nil {
		return nil, err
	}
	FinishHour, err := strconv.Atoi(b.FinishHour)
	if err != nil {
		return nil, err
	}
	FinishMinute, err := strconv.Atoi(b.FinishMinute)
	if err != nil {
		return nil, err
	}

	if StartHour > 23 || StartHour < 0 {
		return nil, errors.New("Block start hour is wrong")
	}
	if StartMinute > 60 || StartMinute < 0 {
		return nil, errors.New("Block start minute is wrong")
	}
	if FinishHour > 23 || FinishHour < 0 {
		return nil, errors.New("Block end hour is wrong")
	}
	if FinishMinute > 60 || FinishMinute < 0 {
		return nil, errors.New("Block end minute is wrong")
	}
	if StartHour > FinishHour || StartHour == FinishHour && StartMinute >= FinishMinute {
		return nil, errors.New("Block start time is equal or greater than finish time")
	}

	b.NumStartHour, b.NumStartMinute = StartHour, StartMinute
	b.NumFinishHour, b.NumFinishMinute = FinishHour, FinishMinute
	return b, nil
}

// SortDay sorts blocks by start time
func SortDay(day datatypes.Day) datatypes.Day {
	for i := 1; i < len(day); i++ {
		j := i
		for j > 0 {
			if day[j].NumStartHour < day[j-1].NumStartHour ||
				(day[j].NumStartHour == day[j-1].NumStartHour && day[j].NumStartMinute < day[j-1].NumStartMinute) {
				day[j], day[j-1] = day[j-1], day[j]
			}
			j--
		}
	}
	return day
}
