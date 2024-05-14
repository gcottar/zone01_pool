package piscine

func Max(a []int) int {
	maxNumber := 0
	for _, n := range a {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber
}
