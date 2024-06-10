package main

import (
	"math"
	"strconv"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	sqrt := math.Sqrt(float64(x))
	return math.Mod(sqrt, 1.0) == 0
}

func IsPalindrome(x int) bool {
	str := strconv.Itoa(int(math.Abs(float64(x))))

	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func Abs(num int) int {
	return int(math.Abs(float64(num)))
}

func Cube(num int) int {
	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	result := []int{}
	for _, v := range input {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map(input []int, m MapperFunc) []int {
	result := []int{}
	for i := range input {
		result = append(result, m(input[i]))
	}
	return result
}
