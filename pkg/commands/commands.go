package commands

import (
	"os/exec"
	"strconv"
)

// CurrentDay returns number of current day of the week
func CurrentDay() int {
	day, _ := exec.Command("/bin/sh", "-c", "date +%w").Output()
	currentDay, _ := strconv.Atoi(string(day[:1]))
	if currentDay == 0 {
		currentDay = 6
	} else {
		currentDay--
	}
	return currentDay
}

// CurrentTime returns current time
func CurrentTime() float32 {
	hour, _ := exec.Command("date", "+%H").Output()
	currentHour, _ := strconv.ParseFloat(string(hour[:2]), 32)
	minute, _ := exec.Command("date", "+%M").Output()
	currentMinute, _ := strconv.ParseFloat(string(minute[:2]), 32)
	return float32(currentHour + currentMinute/60)
}
