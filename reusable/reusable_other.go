package reusable

func Calc2() int {
	if count < 5 {
		count++
		Calc()
	}
	return 0
}
