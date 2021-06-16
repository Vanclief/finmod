package utils

func Abs(a int64) int {
	if a >= 0 {
		return (int)(a)
	}
	return (int)(-a)
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
