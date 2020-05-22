package data

import (
	"errors"
	"github.com/oltarzewskik/tibivi/pkg/common"
	"io/ioutil"
	"strconv"
	"strings"
)

// Day consists time blocks of day of the week
type Day []*Block

// Block consists fields of time block
type Block struct {
	StartHour, StartMinute,
	FinishHour, FinishMinute,
	Description string
	NumStartTime, NumFinishTime float32
}

var Schedule = map[string]Day{}

// SetSchedule supplies `Schedule` with data from datafiles
func SetSchedule() error {
	for _, day := range common.Days {
		data, err := Read(day + ".txt")
		if err != nil {
			return err
		}
		Schedule[day] = data
	}
	return nil
}

// Read converts data defined in datafile and returns it
func Read(filename string) (Day, error) {
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

	day := make(Day, len(data))
	for i, entry := range data {
		day[i] = &Block{}
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
func AddNumTimes(b *Block) (*Block, error) {
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

	NumStartTime := float32(StartHour) + float32(StartMinute)/60
	NumFinishTime := float32(FinishHour) + float32(FinishMinute)/60
	if NumStartTime >= NumFinishTime {
		return nil, errors.New("Block start time is equal or greater than end time")
	}
	b.NumStartTime, b.NumFinishTime = NumStartTime, NumFinishTime
	return b, nil
}

// sortDay sorts blocks by start time
func SortDay(day Day) Day {
	for i := 1; i < len(day); i++ {
		j := i
		for j > 0 {
			if day[j].NumStartTime < day[j-1].NumStartTime {
				day[j], day[j-1] = day[j-1], day[j]
			}
			j--
		}
	}
	return day
}
