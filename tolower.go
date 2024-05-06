package piscine

func ToLower(s string) string {
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] += 32
		}
	}
	return string(bytes)
}
