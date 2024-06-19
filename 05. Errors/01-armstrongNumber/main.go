package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isNumber(r byte) bool {
	_, err := strconv.Atoi(string(r))
	if err != nil {
		return false
	}
	return true
}

func isArmstrong(n int) bool {
	str := strconv.Itoa(n)
	var sum int
	for _, v := range str {
		i, _ := strconv.Atoi(string(v))
		sum += int(math.Pow(float64(i), float64(len(str))))
	}

	return n == sum
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	encrypt := scn.Text()

	i := 0
	var numbers []int
	for i < len(encrypt) {
		if isNumber(encrypt[i]) {
			start := i
			for i < len(encrypt) && isNumber(encrypt[i]) {
				i++
			}
			end := i
			number, _ := strconv.Atoi(encrypt[start:end])
			numbers = append(numbers, number)
		}
		i++
	}

	var sum int
	for _, v := range numbers {
		sum += v
	}

	if isArmstrong(sum) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
