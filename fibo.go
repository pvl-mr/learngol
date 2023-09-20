package learngol

// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55,
func fibonacci(a int) int {
	x1 := 1
	x2 := 1
	if a == 1 || a == 2 {
		return 1
	}
	for i := 3; i < 10; i++ {
		x1, x2 = x2, x2+x1
		if i == a {
			return x2
		}
	}
	return -1
}
