package piscine

func Unmatch(arr []int) int {
	var result int

	for _, el := range arr {
		result = 0
		for _, v := range arr {
			if v == el {
				result++
			}
		}
		if result%2 != 0 {
			return el
		}
	}
	return -1
}
