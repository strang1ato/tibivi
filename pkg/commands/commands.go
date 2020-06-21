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

// CurrentHour returns current hour
func CurrentHour() (int, error) {
	hour, err := exec.Command("date", "+%H").Output()
	if err != nil {
		return 0, err
	}
	currentHour, _ := strconv.Atoi(string(hour[:2]))
	return currentHour, nil
}

// CurrentMinute returns current minute
func CurrentMinute() (int, error) {
	minute, err := exec.Command("date", "+%M").Output()
	if err != nil {
		return 0, err
	}
	currentMinute, _ := strconv.Atoi(string(minute[:2]))
	return currentMinute, nil
}
