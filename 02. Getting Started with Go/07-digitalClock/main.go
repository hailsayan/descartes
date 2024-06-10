package main

import (
	"strconv"
)

func ConvertToDigitalFormat(hour, minute, second int) string {
	hh := strconv.Itoa(hour)
	mm := strconv.Itoa(minute)
	ss := strconv.Itoa(second)
	
	if hour < 10{
	    hh = "0" + hh
	}
	if minute < 10 {
	    mm = "0" + mm
	}
	if second < 10 {
	    ss = "0" + ss
	}
	
	return hh + ":" + mm + ":" + ss
}

func ExtractTimeUnits(seconds int) (int, int, int) {
	hour := seconds / 3600
	minutes := (seconds % 3600 ) / 60
	second := seconds % 60
	
	return hour, minutes, second
}
