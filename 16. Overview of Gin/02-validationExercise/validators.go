package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


var validNationalID validator.Func = func(fl validator.FieldLevel) bool {
	nid, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if len(nid) != 10 {
		return false
	}
	checkNumberStr := nid[9:10]
	checkNumber, err := strconv.Atoi(checkNumberStr)
	if err != nil {
		return false
	}
	sum := 0
	for i := 2; i <= 10; i++ {
		n, err := strconv.Atoi(nid[10-i:11-i])
		if err != nil {
			return false
		}
		sum += i * n
	}
	rem := sum - 11 * (sum / 11)
	if rem > 2 {
		rem = 11 - rem
	}
	return rem == checkNumber
}
	

var validBirthDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	today := time.Now()
	if date.Before(today.Add(-100 * 365 * 24 * time.Hour)) {
		return false
	}
	if date.After(today.Add(-7 * 365 * 24 * time.Hour)) {
		return false
	}
    return true
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("validbirthdate", validBirthDate)
		v.RegisterValidation("validnationalid", validNationalID)
    } else {
		fmt.Println("registerValidators is ok?", ok)
	}
}