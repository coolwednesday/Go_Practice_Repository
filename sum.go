package main

// Sum function takes an integer x and returns the numbers 1 to x
func Sum(x int) (y int) {
	for i := 1; i <= x; i++ {
		y += i
	}
	return
}
