package allocation

func CreateSlice() []int {
	// Arrays are pass by values
	var p []int = make([]int, 10, 100)
	return p
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return // return sum explicity
}
