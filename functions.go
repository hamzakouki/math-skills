package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func readingfile(str string) []int {
	var slice []int
	data, error := os.ReadFile(str)
	if error != nil {
		fmt.Println("there is a problem reading file")
		os.Exit(0)
	}
	f := strings.Split(string(data), "\n")
	for i := 0; i < len(f); i++ {
		intg, error := strconv.Atoi(f[i])
		if error != nil {
			fmt.Println("there is a problem in convert byte to int")
			os.Exit(1)
		}
		slice = append(slice, intg)
	}
	return slice
}