package piscine

func ToUpper(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] -= 32
		}
	}
	return string(bytes)
}
