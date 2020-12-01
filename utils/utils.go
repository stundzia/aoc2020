package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type solutionGetter func()

func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Println("Err msg: ", msg, " Err: ", err)
		log.Fatalf("Fatal Exception `%s` : %s", msg, err)
	}
}

func LoadInputTxtFromFile(dayNum int) string {
	fname := fmt.Sprintf("./day%d/input.txt", dayNum)
	absPath, _ := filepath.Abs(fname)
	dat, err := ioutil.ReadFile(absPath)
	if err != nil {
		FailOnError(err, "Could not load file")
	}
	return string(dat)
}


func AllIntsUnique(ints []int) bool {
	mm := map[int]struct{}{}
	for _, i := range ints {
		mm[i] = struct{}{}
	}
	return len(mm) == len(ints)
}

func LoadInputAsStringSlice(day int, sep string) []string {
	input := LoadInputTxtFromFile(day)
	return strings.Split(input, sep)
}

func LoadInputAsIntSlice(day int, sep string) []int {
	res := []int{}
	input := LoadInputTxtFromFile(day)
	strs := strings.Split(input, sep)
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		FailOnError(err, "loading input as ints failed")
		res = append(res, num)
	}
	return res
}

func SliceStrToInt(slice []string) []int {
	res := []int{}
	for _, item := range slice {
		i, err := strconv.Atoi(item)
		FailOnError(err, "Could not parse int")
		res = append(res, i)
	}
	return res
}

func EvaluateRuntime(function solutionGetter) {
	start := time.Now()
	function()
	fmt.Println("Solution evaluation took: ", time.Now().Sub(start))
}