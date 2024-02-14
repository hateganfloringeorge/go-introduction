package iteration

func Repeater(char string, repeatCount int) string {
	var repeated string
	for range repeatCount {
		repeated += char
	}
	return repeated
}