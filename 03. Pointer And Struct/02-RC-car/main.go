package main

import "fmt"

type Car struct {
	speed   int
	battery int
}

func NewCar(speed, battery int) *Car {
	return &Car{speed, battery}
}

func GetSpeed(car *Car) int {
	return car.speed
}

func GetBattery(car *Car) int {
	return car.battery
}

func ChargeCar(car *Car, minutes int) {
	a := car.battery + (minutes / 2)
	if a <= 100 {
		car.battery += (minutes / 2)
	} else {
		car.battery = 100
	}
}

func TryFinish(car *Car, distance int) (string) {
	if car.battery >= distance/2 {
		car.battery -= (distance / 2)
		return fmt.Sprintf("%.2f", float64(distance)/float64(car.speed))
	} else {
		car.battery = 0
	}

	return ""
}
