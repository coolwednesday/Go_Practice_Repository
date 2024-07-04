package main

const pi = 3.14

// PerimeterCircle takes the radius x and returns circumference
func PerimeterCircle(x float64) float64 {

	if x < 0 {
		return -1
	}
	return 2 * pi * x
}

// PerimeterRectangle takes the length x and breadth y and returns the perimeter
func PerimeterRectangle(x float64, y float64) float64 {
	if x < 0 || y < 0 {
		return -1
	}
	return 2 * (x + y)
}
