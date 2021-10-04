package reusable

var count = 0

func Calc() int {
	if count < 5 {
		count++
		Calc2()
	}
	return 0
}
