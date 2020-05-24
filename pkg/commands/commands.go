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

// CurrentTime returns current time
func CurrentTime() (float32, error) {
	hour, err := exec.Command("date", "+%H").Output()
	if err != nil {
		return 0, err
	}
	currentHour, _ := strconv.ParseFloat(string(hour[:2]), 32)
	minute, err := exec.Command("date", "+%M").Output()
	if err != nil {
		return 0, err
	}
	currentMinute, _ := strconv.ParseFloat(string(minute[:2]), 32)
	return float32(currentHour + currentMinute/60), nil
}
