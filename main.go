package main

import (
	"fmt"
	mathskills "mathskills/calculfunctions"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("only one file must be entered")
		os.Exit(0)
	}
	str := os.Args[1]
	slice := readingfile(str)
	fmt.Println(mathskills.Average(slice))
	fmt.Println(mathskills.Median(slice))
	fmt.Println(mathskills.Variance(slice, mathskills.Average(slice)))
	fmt.Println(mathskills.StandardDeviation(mathskills.Variance(slice, mathskills.Average(slice))))
}
