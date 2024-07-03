package main

import (
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

var c, python, java bool

func pow(x, m, lim float64) float64 {
	if v := math.Pow(x, m); v < lim {
		return v
	} else {
		return lim
	}

}

type Vertex struct {
	X int
	Y int
}

func main() {
	/*fmt.Println("Hello World")
	fmt.Println("The time is", time.Now())
	fmt.Println("My Favourite Number is", rand.Intn(10))
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	var i int
	fmt.Println(i, c, python, java)
	*/
	/*sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	*/
	/*for sum < 1000 {
		sum += 10
	}
	fmt.Println(sum)
	*/
	/*
		j := 0
		switch j {
		case 0:
			Println("hello world")
		default:
			Println("hey")
		}
		fmt.Println(pow(2, 3, 5))
	*/
	/*i := 42
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	*/
	/*
		v := Vertex{1, 2}
		v.X = 4
		p := &v
		var a [2]string
		a[0] = "Hello"
		fmt.Println(a[1])

		fmt.Println(p.Y)
		fmt.Println(v.X)
	*/ /*
		names := [4]string{"a", "b", "c"}

		a := names[1:2]
		var s []int
		s = append(s, 1, 2, 3, 4)
	*/
	/*
		var y = make(map[string]Vertex)
		y["Divya"] = Vertex{1, 2}
		elem, ok := y["x"]
		fmt.Println(elem, ok)
	*/

	/*
		fmt.Println(names, a)
		fmt.Println(a[:])
	*/
}
