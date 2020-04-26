package tibivi

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Block consists fields of time block
type Block struct {
	startHour, startMinute,
	endHour, endMinute,
	content string
	numStartTime, numEndTime float32
}

// Day consists time blocks of day of the week
type Day []*Block

// Schedule consists whole week schedule
type Schedule map[string]Day

// createDatafiles creates tibivi's datafiles if they don't exist
func (tbv *Tibivi) createDatafiles() error {
	for _, day := range tbv.days {
		if _, err := os.Stat(tbv.dotTibivi + day + ".txt"); os.IsNotExist(err) {
			if _, err := os.Create(tbv.dotTibivi + day + ".txt"); err != nil {
				return err
			}
		}
	}
	return nil
}

// setSchedule supply `tbv.Schedule` with data from datafiles
func (tbv *Tibivi) setSchedule() error {
	for _, day := range tbv.days {
		data, err := tbv.getData(day + ".txt")
		if err != nil {
			return err
		}
		tbv.Schedule[day] = data
	}
	return nil
}

// getData converts data defined in datafile and returns it
func (tbv *Tibivi) getData(filename string) (Day, error) {
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
		blockFields := []*string{&day[i].startHour, &day[i].startMinute, &day[i].endHour, &day[i].endMinute, &day[i].content}
		var section int
		for _, char := range entry {
			if section != 4 && (char == ':' || char == ' ' || char == '-') {
				section++
				continue
			}
			*blockFields[section] += string(char)
		}

		day[i], err = tbv.addNumTimes(day[i])
		if err != nil {
			return nil, err
		}
	}

	day = tbv.sortDay(day)
	return day, nil
}

// addNumTimes adds numeric times and check if there are no defects in block
func (tbv *Tibivi) addNumTimes(b *Block) (*Block, error) {
	startHour, err := strconv.Atoi(b.startHour)
	if err != nil {
		return nil, err
	}
	startMinute, err := strconv.Atoi(b.startMinute)
	if err != nil {
		return nil, err
	}
	endHour, err := strconv.Atoi(b.endHour)
	if err != nil {
		return nil, err
	}
	endMinute, err := strconv.Atoi(b.endMinute)
	if err != nil {
		return nil, err
	}

	if startHour > 23 || startHour < 0 {
		return b, errors.New("Block start hour is wrong")
	}
	if startMinute > 60 || startMinute < 0 {
		return b, errors.New("Block start minute is wrong")
	}
	if endHour > 23 || endHour < 0 {
		return b, errors.New("Block end hour is wrong")
	}
	if endMinute > 60 || endMinute < 0 {
		return b, errors.New("Block end minute is wrong")
	}

	numStartTime := float32(startHour) + float32(startMinute)/60
	numEndTime := float32(endHour) + float32(endMinute)/60
	if numStartTime >= numEndTime {
		return b, errors.New("Block start time is equal or greater than end time")
	}
	b.numStartTime, b.numEndTime = numStartTime, numEndTime
	return b, nil
}

// sortDay sorts blocks by start time
func (tbv *Tibivi) sortDay(day Day) Day {
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
