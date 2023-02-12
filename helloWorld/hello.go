package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

func main() {
	fmt.Println(quote.Hello())

	// This
	var n = 2
	loopHello(&n)
	// Or this
	n2 := 3
	loopHello(&n2)

	var x = "I don't believe in magic!"
	fmt.Println(x)
	fmt.Println(doYouBelieveInMagic(&x))
	pointerMagic(&x)
	fmt.Println(x)
	fmt.Println(doYouBelieveInMagic(&x))
	// Enable this to see infinite loop
	// infiniteLoop()

	n3 := 6
	scopingTheVariable(&n3)

	dumbStatement := "FFXIV Stormblood is the best expansion."
	switchTheStatement(&dumbStatement)

	deferTheCallback("Rio", "Arswendo", func(lastname *string) {
		fmt.Println(*lastname)
	})

	n4 := 10
	stackedDefer(&n4)

	class := aStructuredClassroom("Class A", 30, 27.8)
	fmt.Println("Class name: ", class.className)
	fmt.Println("Student count: ", class.studentCount)
	fmt.Println("Average grade: ", class.averageGrade)

	arrayAndSlice()
	weirdSliceProperty()
	dynamicSizedArray(5)

	rangeLooping()
	fmt.Println("Treasure Coordinate: ", aMapOfTreasure("Gulag"))
	fmt.Println("Treasure Coordinate: ", aMapOfTreasure("Pentagon"))
	fmt.Println("Treasure Coordinate: ", aMapOfTreasure("Indonesia"))
	mapMutation()

	printFibonacci();
	boi := Dog{name: "Boi"}
	boi.bork();
}

func loopHello(n *int) {
	for i := 0; i < *n; i++ {
		fmt.Printf("Hello, number: %v\n", i)
	}
}

func pointerMagic(x *string) {
	*x = "This is Pointer magic."
}

func doYouBelieveInMagic(statement *string) string {
	if *statement == "This is Pointer magic." {
		return "See, now you too believe in magic."
	} else {
		return "We'll see later man."
	}
}

func scopingTheVariable(x *int) {
	// v can be accesed from else or if as well as long as it within this statement scope
	if v := 2; *x > 5 {
		// y can only accessed within this scope
		y := *x * 20
		fmt.Println("Y: ", y)
		fmt.Println("V: ", v)
	} else {
		fmt.Println("V FROM ELSE: ", v)
	}
	// Undefined Error
	// fmt.Println("Y2: ", y);
}

func switchTheStatement(dumbStatement *string) {
	switch *dumbStatement {
	case "Earth is flat.":
		fmt.Println("No Earth isn't flat")
	case "Assasin's Creed is a good franchise.":
		fmt.Println("Bruh")
	default:
		fmt.Println("Get some help.")
	}
}

func deferTheCallback(firstname, lastname string, callback func(lastname *string)) {
	// 3. Then this
	defer callback(&lastname)
	// 1. This first
	fmt.Print(firstname, " ")
	// 2. Return null
}

func stackedDefer(n *int) {
	// 1. This
	fmt.Println("Counting Backward: ")

	// 2. Place the whole defer in a stack
	for i := 0; i < *n; i++ {
		// 4-n this get executed in a Last-in-first-out manner
		// 9, 8, 7, 6, 5, 4, 3, 2, 1, 0
		defer fmt.Println(i)
	}
	// 3. Return null
}

type Classroom struct {
	className    string
	studentCount int
	averageGrade float64
}

func aStructuredClassroom(className string, studentCount int, averageGrade float64) Classroom {
	return Classroom{
		className,
		studentCount,
		averageGrade,
	}
}

func arrayAndSlice() {
	// This is an array, the size is static and can't change
	fruits := [5]string{"Apple", "Blueberry", "Cherries", "Durian", "Eggplant"}

	// This is a slice, the size is dynamic. Apparently this is more common
	// Apparently this is just a reference to the array and does not store values
	// if i change one value here it should change on fruits as well.
	var fruitSlices []string = fruits[0:3]
	// Or
	// fruitSlices := fruits[1:4]
	fruitSlices[0] = "Mango"

	fmt.Println(fruitSlices)
	fmt.Println(fruits) // Should be: [Mango, Blueberry, Cherries, Durian, Eggplant]
}

func weirdSliceProperty() {
	// you have to remember that slice is a reference
	// if one change everything change
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	// this doesnt change the capacity
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	// if we slice from the low
	// then assign it, it actually affect the underlying array
	// decreasing its capacity
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func dynamicSizedArray(n int) {
	// Makes a slice with len 0 and capacity n
	dynamicArr := make([]int, 0, n)

	// Extending the slice to the length of the capacity
	dynamicArr = dynamicArr[:cap(dynamicArr)]

	dynamicArr[0] = 1
	fmt.Println(dynamicArr)
}

func rangeLooping() {
	fruits := []string{"Apple", "Blueberry", "Cherries", "Durian", "Eggplant"}
	for i, v := range fruits {
		// Doesnt affect the array/slice
		// because range return the COPY
		// not the reference
		v = "Dragonfruit"
		fmt.Printf("%d. %v\n", i, v)
	}
	fmt.Println("FULL FRUITS: ", fruits)
}

type Coordinate struct {
	x, y, z float64
}

func aMapOfTreasure(key string) Coordinate {
	// treasureMap := make(map[string]Coordinate)
	// treasureMap["Gulag"] = Coordinate{
	// 	x: 1.002,
	// 	y: 2.22323,
	// 	z: 932.21,
	// }
	// treasureMap["Pentagon"] = Coordinate{
	// 	x: 999,
	// 	y: 220,
	// 	z: 123.002,
	// }

	// OR
	var treasureMap = map[string]Coordinate{
		"Gulag": {
			x: 1.002,
			y: 2.22323,
			z: 932.21,
		},
		"Pentagon": {
			x: 999,
			y: 220,
			z: 123.002,
		},
	}
	return treasureMap[key]
}

func mapMutation() {
	// Making a dynamic map
	m := make(map[string]int)

	// Assigning value to a key
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// Deleting a value from a key
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Check if the value inside a key exist
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func printFibonacci() {
	fib := closureFibonacci();
	for i := 0; i < 10; i++ {
		fib();
	}
	fmt.Println();
}

// This is a closure function, a function that return a function
func closureFibonacci() func() {
	// These variable is kept independently every loop
	i := -1;
	sequence := []int{}
	return func() {
		i++
		if i > 1 {
			sequence = append(sequence, sequence[i - 1] + sequence[i - 2])
		} else {
			sequence = append(sequence, i)
		}
		fmt.Printf("%v ", sequence[i])
	}
}

type Dog struct {
	name string
}
// This is a method on type Dog, any type can have a method, it doesnt have to be a struct
func (d Dog) bork() {
	fmt.Printf("%v said: %v\n", d.name, "Bork!");
}

// 2D Array example
// func Pic(dx, dy int) [][]uint8 {
// 	img := make([][]uint8, dy)
// 	for i := range img {
// 		img[i] = make([]uint8, dx)
// 		for j := range img[i] {
// 			img[i][j] = uint8(10^2/2*3)
// 		}
// 	}
// 	return img
// }

// func infiniteLoop() {
// 	// This is the while on Go
// 	for {
// 		fmt.Println("This is looping infinitely.")
// 	}
// }
