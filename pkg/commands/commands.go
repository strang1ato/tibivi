package commands

import (
	"os/exec"
	"strconv"
)

// CurrentDay returns number of current day of the week
func CurrentDay() (int, error) {
	day, err := exec.Command("/bin/sh", "-c", "date +%w").Output()
	if err != nil {
		return 0, err
	}
	currentDay, _ := strconv.Atoi(string(day[:1]))
	if currentDay == 0 {
		currentDay = 6
	} else {
		currentDay--
	}
	return currentDay, nil
}

// CurrentTime returns current hour and minute
func CurrentTime() (int, int, error) {
	time, err := exec.Command("date", "+%R").Output()
	if err != nil {
		return 0, 0, err
	}
	currentHour, _ := strconv.Atoi(string(time[:2]))
	currentMinute, _ := strconv.Atoi(string(time[3:5]))
	return currentHour, currentMinute, nil
}
