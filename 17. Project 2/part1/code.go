package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	flights := initializeFlights()
	processPassengers(flights)
	processOpinions(flights)
	calculateAndPrintAverages(flights)
}

func initializeFlights() map[string]map[string][2]interface{} {
	scanner.Scan()
	flightsNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	flights := make(map[string]map[string][2]interface{}, flightsNumber)
	for i := 0; i < flightsNumber; i++ {
		scanner.Scan()
		flightName := scanner.Text()
		flights[flightName] = make(map[string][2]interface{})
	}
	return flights
}

func processPassengers(flights map[string]map[string][2]interface{}) {
	scanner.Scan()
	passengersNumber, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	for i := 0; i < passengersNumber; i++ {
		scanner.Scan()
		pf := scanner.Text()
		spf := strings.Split(pf, " ")
		if len(spf) < 2 {
			continue
		}
		passengerName, flightName := spf[0], spf[1]
		if _, exists := flights[flightName]; !exists {
			fmt.Printf("Invalid flight %s\n", flightName)
			continue
		}
		if _, exists := flights[flightName][passengerName]; exists {
			fmt.Printf("Duplicate ticket for %s %s\n", flightName, passengerName)
			continue
		}
		flights[flightName][passengerName] = [2]interface{}{}
	}
}

func processOpinions(flights map[string]map[string][2]interface{}) {
	scanner.Scan()
	opinionNumbers, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	for i := 0; i < opinionNumbers; i++ {
		scanner.Scan()
		op := scanner.Text()
		sop := strings.SplitN(op, " ", 4)
		if len(sop) < 4 {
			continue
		}
		passengerName, flightName, score, comment := sop[0], sop[1], sop[2], sop[3]
		if _, exists := flights[flightName]; !exists {
			fmt.Printf("Invalid flight %s\n", flightName)
			continue
		}
		if _, exists := flights[flightName][passengerName]; !exists {
			fmt.Printf("Invalid passenger for %s %s\n", flightName, passengerName)
			continue
		}
		if flights[flightName][passengerName][1] != nil {
			fmt.Printf("Duplicate comment for %s by %s\n", flightName, passengerName)
			continue
		}
		integerScore, err := strconv.Atoi(score)
		if err != nil {
			panic(err)
		}
		flights[flightName][passengerName] = [2]interface{}{integerScore, comment}
		fmt.Printf("Accepted comment for %s by %s\n", flightName, passengerName)
	}
}

func calculateAndPrintAverages(flights map[string]map[string][2]interface{}) {
	keys := make([]string, 0, len(flights))
	for key := range flights {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		if len(flights[key]) != 0 {
			flightScoreSum := 0
			opinionNumber := 0
			for _, opinion := range flights[key] {
				if score, ok := opinion[0].(int); ok {
					flightScoreSum += score
					opinionNumber++
				}
			}
			if opinionNumber > 0 {
				avgSum := float64(flightScoreSum) / float64(opinionNumber)
				fmt.Printf("Average score for %s is %.2f\n", key, avgSum)
			}
		}
	}
}
