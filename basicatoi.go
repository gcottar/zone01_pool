package piscine

func BasicAtoi(s string) int {
	isNeg := false
	result := 0
	index := 0
	len := StrLen(s)

	if index < len && (s[index] == '+' || s[index] == '-') {
		if s[index] == '-' {
			isNeg = true
		}
		index++
	}
	for index < len && s[index] == '0' {
		index++
	}
	for index < len && s[index] >= '0' && s[index] <= '9' {
		result *= 10
		result += int(s[index] - '0')
		index++
	}
	if isNeg {
		result *= -1
	}
	return result
}
