package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var m = make(map[string]string)
	for i:=0; i<n; i++ {
		var c1, c2 string
		fmt.Scanf("%s %s", &c1, &c2)
		m[c2] = c1
	}
	var q int
	fmt.Scanf("%d", &q)
	res := []string{}
	for i:=0; i<q; i++ {
		var phone string
		fmt.Scanf("%s", &phone)
		c, ok := m[phone[:3]]
		if ok {
			res = append(res, c)
		} else {
			res = append(res, "Invalid Number")
		}
	}

	for _, r := range(res) {
		fmt.Println(r)
	}
}
