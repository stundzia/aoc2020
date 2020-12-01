package day1

import (
	"aoc2020/utils"
	"fmt"
)

func getProductOfMatchingSumOfTwoItems(intSlice []int, sum int) int {
	for i, num := range intSlice {
		for t, num2 := range intSlice {
			if i == t {
				continue
			}
			if num + num2 == sum {
				return num * num2
			}
		}
	}
	return -1
}

func getProductOfMatchingSumOfThreeItems(intSlice []int, sum int) int {
	for i, num := range intSlice {
		for t, num2 := range intSlice {
			for z, num3 := range intSlice {
				if i == t || t == z || i == z {
					continue
				}
				if num + num2 + num3 == sum {
					return num * num2 * num3
				}
			}
		}
	}
	return -1
}

func DoSilver() {
	nums := utils.LoadInputAsIntSlice(1, "\n")
	fmt.Println("Solution: ", getProductOfMatchingSumOfTwoItems(nums, 2020))
}

func DoGold() {
	nums := utils.LoadInputAsIntSlice(1, "\n")
	fmt.Println("Solution: ", getProductOfMatchingSumOfThreeItems(nums, 2020))
}