package main

import (
	"math"
)

func AddElement(numbers *[]int, element int) {
	//TODO
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}
	min := math.MaxInt64
	for _, v := range *numbers {
		if min > v {
			min = v
		}
	}
	return min
}

func ReverseSlice(numbers *[]int) {
	//TODO
	for i := 0; i < len(*numbers)/2; i++ {
		j := len(*numbers) - i - 1
		SwapElements(numbers, i, j)
	}

}

func SwapElements(numbers *[]int, i, j int) {
	//TODO
	if len(*numbers)-1 < i || len(*numbers)-1 < j || i < 0 || j < 0 {
		return
	}
	(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
}
