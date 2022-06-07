package iteration

func Repeat(value string, times int) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += value
	}
	return repeated
}
