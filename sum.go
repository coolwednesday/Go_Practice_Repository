package main

// Sum function takes an integer x and returns the numbers 1 to x
func Sum(x int) (y int) {
	if x <= 0 {
		return 0
	}
	for i := 1; i <= x; i++ {
		y += i
	}
	return
}
