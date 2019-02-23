package main

import "fmt"

type square struct {
	X float64
}

type circle struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square:
		fmt.Println("This is a square")
	case circle:
		fmt.Printf("%v is a circle\n", v)
	case rectangle:
		fmt.Println("This is a rectangle!")
	default:
		fmt.Printf("Unknown type %T!", v)
	}
}

func main() {

}
