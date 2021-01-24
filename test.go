
package main

import (
	"fmt"
	"strconv"
)

func method1(x1, x2 int) int {
	res := x1 + x2
	return res
}

func method2(x1, x2 int) (int, int) {
	return x1, x2
}

func methodMethod(myMethod func(string, int, byte) int) {
	// you can also have a func type return
	fmt.Println(myMethod("z", 1, 'z'))
}

func changeVariable(x string) {
	x = "tim"
}

func changeMap(x map[string]int) {
	x["tim"] = 0
}

func changeArray(x []int) {
	x[0] = 0
}
func changeVariableForReal(x *string) {
	*x = "tim"
}

type Coordination struct {
	x int16
	y int16
}

func changeX(input Coordination) {
	input.x = 0
}

func changeXForReal(input *Coordination) {
	input.x = 0
}

type Square struct {
	point1 *Coordination
	point2 *Coordination
	point3 *Coordination
	point4 *Coordination
}

func main() {
	name1 := "tim"
	name2 := "tim"
	res := name1 == name2
	fmt.Printf("%t", res)
	arr := [6]int{1, 2, 3, 4, 5, 6}
	var expectedAssignmentArray [3]int
	arrSlice1 := arr[:1]
	arrSlice2 := arr[1:]
	arrSlice3 := arr[2:cap(arrSlice2)]
	// arrSlice3 := arr[:-1] invalid
	// arrSlice3 := arr[-1:] invalid
	twoDimensionalArray := [3][]int{arrSlice1, arrSlice2, arrSlice3}
	for i := 1; i <= len(twoDimensionalArray); i++ {
		fmt.Println(twoDimensionalArray[i - 1])
	}
	for i := 1; i <= len(expectedAssignmentArray); i++ {
		expectedAssignmentArray[i - 1] = i + 1
	}
	for index, value := range arrSlice1 {
		fmt.Println(index, value)
	}
	var relation map[string]int = map[string]int{
		"jin1":1,
		"jin2":2,
		// "jin1":2, invalid
	}
	fmt.Println(relation["jin1"])
	keyFromRelation, valueFromrelation := relation["jin2"]
	fmt.Println(keyFromRelation, valueFromrelation)
	fmt.Println(len(relation))
	for key, value := range relation {
		fmt.Println(key, value)
	} 
	defer fmt.Println(method1(1, 2))
	defer fmt.Println(method2(1, 2))
	defer fmt.Println(method1(3, 4))
	defer fmt.Println(method2(3, 4))
	defer fmt.Println("---------------")
	variableAssignedbyMethod := method1
	fmt.Println(variableAssignedbyMethod(4, 5))
	variableAssignedByMethodDefination1 := func(x, y int) int {
		return x * y
	}(6, 6)
	fmt.Println(variableAssignedByMethodDefination1)
	variableAssignedByMethodDefination2 := func(x, y int) int {
		return x * y
	}
	fmt.Println(variableAssignedByMethodDefination2(7, 7))
	s := "97"
	_, err := strconv.ParseInt(s, 10, 64)
	fmt.Printf("this is the type of error %T \n", err)
	methodUsedForMethodMethod := func(x string, y int, z byte) int {
		inter, _ := strconv.Atoi(x)
		return inter + y + int(z)
	}
	methodMethod(methodUsedForMethodMethod)
	variable := "tom"
	mapVariable := map[string]int {
		"tim": 1,
	}
	arrayVariable := []int{1, 1, 1, 1}
	changeVariable(variable)
	fmt.Println(variable)
	changeMap(mapVariable)
	fmt.Println(mapVariable)
	changeArray(arrayVariable)
	fmt.Println(arrayVariable)
	location := 1
	fmt.Println(location)
	fmt.Println(&location)
	addressOfLocationVariable := &location
	fmt.Println(addressOfLocationVariable)
	location = 8
	fmt.Println(location)
	fmt.Println(&location)
	fmt.Println(addressOfLocationVariable)
	*addressOfLocationVariable = 10
	fmt.Println(location)
	fmt.Println(&location)
	fmt.Println(addressOfLocationVariable)
	changeVariableForReal(&variable)
	fmt.Println(variable)
	change := "world"
	var pointer *string = &change
	fmt.Println(change, &change)
	fmt.Println(pointer, &pointer)
	*pointer = "after change"
	fmt.Println(change, &change)
	fmt.Println(pointer, &pointer)
	coordination := Coordination{x : 1, y : 1}
	fmt.Println(coordination)
	changeX(coordination)
	fmt.Println(coordination)
	coordinationLocation := &Coordination{x : 1, y : 1}
	fmt.Println(coordinationLocation)
	changeXForReal(coordinationLocation)
	fmt.Println(*coordinationLocation)
	fmt.Println(coordination.x)
	fmt.Println(coordinationLocation.x)
	square := Square{point1 : &Coordination{1, 1}, point2 : &Coordination{2, 2}, point3 : &Coordination{x : 3, y : 3}, point4 : &Coordination{x : 4, y : 4}}
	fmt.Println(square)
	fmt.Println(square.point1)
	fmt.Println(square.point1.x)
} 