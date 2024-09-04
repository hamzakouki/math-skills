package mathskills

func Variance(slice []int, n int) int {
	moin := 0
	var intg2 int
	for i := 0; i < len(slice); i++ {
		moin = n - slice[i]
		moin = moin * moin
		intg2 = intg2 + moin
		moin = 0
	}
	return intg2 / n
}
