# go-101

Method to declare an slice (abstraction of array)
 
 ```go
func array() ([]int) {
	arr := []int{0,1,2,3,4}
	return arr
}
```

Maps which ressemble `Record` in typescript

```go
func vertices() (map[int]string) {
	vertices := make(map[int]string)

	vertices[3] = "triangle"
	vertices[4] = "square"
	vertices[5] = "pentagon"
	vertices[6] = "hexagon"

	return vertices
}
```

Basic loop

```go
func loops() {
	arr := array()
	for index, entry := range arr {
		fmt.Println("index:", index)
		fmt.Println("element:", entry)
	}
}
```

Calculating srt

```go
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
```

Structs

```go
type person struct {
	name string
	age int
	gender string
	etnicity string
}

func structHandle() (person) {
	p := person{name: "Rodrigo", age: 33, gender: "male", etnicity: "black"}
	return p
}
```

Pointer

```go
func pointer() (int) {
	i := 7
	incrementPointer(&i)
	return i
}

func incrementPointer(x *int) {
	*x++
}
```

Output from main

```bash
[0 1 2 3 4]
map[3:triangle 4:square 5:pentagon 6:hexagon]
index: 0
element: 0
index: 1
element: 1
index: 2
element: 2
index: 3
element: 3
index: 4
element: 4
4
{Rodrigo 33 male black}
8
```
