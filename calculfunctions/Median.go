package mathskills

func Median(slice []int) int {
	var intg int
	var intg1 int
	var intg2 int
	var slice1 []int
	var slice2 []int
	if len(slice)%2 != 0 {
		intg = slice[len(slice)/2]
		return intg
	} else {
		slice1 = slice[:len(slice)/2]
		slice2 = slice[len(slice)/2:]
		intg1 = slice1[len(slice1)-1]
		intg2 = slice2[0]
	}
	return (intg1+intg2)/2
}