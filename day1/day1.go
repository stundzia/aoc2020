package day1

import (
	"aoc2020/utils"
	"fmt"
)


func DoSilver() {
	nums := utils.LoadInputAsIntSlice(1, "\n")
	for i, num := range nums {
		for t, num2 := range nums {
			if i == t {
				continue
			}
			if num + num2 == 2020 {
				fmt.Println("Solution: ", num * num2)
				return
			}
		}
	}
}

func DoGold() {
	nums := utils.LoadInputAsIntSlice(1, "\n")
	for i, num := range nums {
		for t, num2 := range nums {
			for z, num3 := range nums {
				if i == t || t == z || i == z {
					continue
				}
				if num + num2 + num3 == 2020 {
					fmt.Println("Solution: ", num * num2 * num3)
					return
				}
			}
		}
	}
}