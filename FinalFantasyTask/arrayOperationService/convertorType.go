package arrayOperationService

func convertFloat64ToInt(ar []float64) []int {
	newar := make([]int, len(ar))
	var v float64
	var i int
	for i, v = range ar {
		newar[i] = int(v)
	}
	return newar
}
