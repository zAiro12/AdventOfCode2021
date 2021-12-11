package main

import (
	"fmt"
)

func main() {
	var c [8]int
	//k := 256
	for {
		var in int
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		switch in {
		case 1:
			c[0]++
		case 2:
			c[1]++
		case 3:
			c[2]++
		case 4:
			c[3]++
		case 5:
			c[4]++

		}
	}
	ctot := c[0] + c[1] + c[2] + c[3] + c[4] + c[5] + c[6] + c[7]
	for i := 0; i < 2; i++ {
		c = scala(c)
	}
	ctot += c[0] + c[1] + c[2] + c[3] + c[4] + c[5] + c[6] + c[7]
	fmt.Println(ctot)
}

func scala(c [8]int) [8]int {
	scambio := 0
	for i := 7; i >= 0; i-- {
		somm := c[0]
		if i != 0 {
			scambio = c[i-1]
			c[i-1] = c[i]
		} else {
			c[7] = scambio
		}
		if i == 5 || i == 7 {
			c[i] += somm
		}
	}
	return c
}
