package main

import (
	"fmt"
	"math"
	"errors"
)

type person struct {
	name string
	age int
	gender string
	etnicity string
}

//array function returns an array of integers
func array() ([]int) {
	arr := []int{0,1,2,3,4}
	return arr
}


//vertices function returns a map of integers and strings
func vertices() (map[int]string) {
	vertices := make(map[int]string)

	vertices[3] = "triangle"
	vertices[4] = "square"
	vertices[5] = "pentagon"
	vertices[6] = "hexagon"

	return vertices
}

//loops function loops through the array and prints each element
func loops() {
	arr := array()
	for index, entry := range arr {
		fmt.Println("index:", index)
		fmt.Println("element:", entry)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func structHandle() (person) {
	p := person{name: "Rodrigo", age: 33, gender: "male", etnicity: "black"}
	return p
}

func pointer() (int) {
	i := 7
	incrementPointer(&i)
	return i
}

func incrementPointer(x *int) {
	*x++
}

func main() {
	fmt.Println(array())
	fmt.Println(vertices())
	loops()
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(structHandle())
	fmt.Println(pointer())
}