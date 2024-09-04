package mathskills

import (
	"fmt"
	"math"
)

func Average(slice []int) int {
	var intg int
	for i := 0; i < len(slice); i++ {
		intg = intg + slice[i]
	}
	fmt.Println(intg)
	return int(math.Round(float64(intg /(len(slice)))))
}
