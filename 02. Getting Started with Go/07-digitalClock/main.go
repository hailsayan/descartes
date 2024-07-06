package main

import (
	"fmt"
	"strconv"
)

func addZero(number int) string {
	if number < 10 {
		return "0" + strconv.Itoa(number)
	}
	return strconv.Itoa(number)
}

func ConvertToDigitalFormat(hour, minute, second int) string {
	return fmt.Sprintf("%s:%s:%s", addZero(hour), addZero(minute), addZero(second))
}
func ExtractTimeUnits(seconds int) (int, int, int) {
	hour := int(seconds / 3600)
	seconds %= 3600
	minute := int(seconds / 60)
	seconds %= 60
	return hour, minute, seconds
}
