package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		inp := scanner.Text()
		inpSplit := strings.Split(inp, " ")
		nums := make([]int, 0)
		for j := 1; j < len(inpSplit); j++ {
			tmp, _ := strconv.Atoi(inpSplit[j])
			nums = append(nums, tmp)
		}
		count := 0
		for j := 0; j < len(nums)-2; j++ {
			diff1 := nums[j+1] - nums[j]
			diff2 := nums[j+2] - nums[j+1]
			if diff1 != diff2 {
				continue
			}
			count++
			prev := nums[j+2]
			for k := j + 3; k < len(nums); k++ {
				if nums[k]-prev == diff1 {
					count++
					prev = nums[k]
				} else {
					break
				}
			}
		}
		fmt.Printf("%s %d\n", inpSplit[0], count)
	}
}
