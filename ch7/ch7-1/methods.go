package main

import "fmt"

type twoInts struct {
	X int64
	Y int64
}

func regularFunctions(a, b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

// 객체의 method 화
func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

func main() {
	i := twoInts{X: 1, Y: 2}
	j := twoInts{X: -5, Y: -2}

	fmt.Println(regularFunctions(i, j))
	fmt.Println(i.method(j))
}
