package main

import (
	"fmt"
	// "image"
	// "image/color"
	"io"
	"strings"
	"time"
	// "golang.org/x/tour/pic"
)

type Animal interface {
	// Return the noise of the animal
	MakeNoise() string
}

// Empty interface accept anything
// Probably the same as "any" in typescript
type AnyInterface interface{}

type DogName string

type CatName string

type Coordinate struct {
	x, y float64
}

type Image struct{}

func main() {
	implicitInterface()
	// This also works;
	var x interface{} = CatName("NewCat")
	catName := x.(CatName)
	customPrint(catName.MakeNoise())
	customPrint(&Coordinate{12.23, 2.8282})

	// v, err := stableFunction(-1);
	// fmt.Println(v)
	// fmt.Println(err)
	if _, err := stableFunction(-1); err != nil {
		fmt.Println(err)
	}

	readTextStream()

	p1, p2, p3 := "James", "Luke", "Mark"
	s1, s2, s3 := 0, 0, 0
	foodCount := 7
	go eatFood(&foodCount, &p3, &s3)
	go eatFood(&foodCount, &p1, &s1)
	eatFood(&foodCount, &p2, &s2)

	fmt.Printf("Score:\nPlayer 1: %v\nPlayer 2: %v\nPlayer 3: %v\n", s1, s2, s3);
}

func customPrint(x interface{}) {
	fmt.Println(x)
}

// Animal will take any type as long as it has
// the same methods. Like typescript
func implicitInterface() {
	var animal Animal

	boi := DogName("Boi")
	animal = &boi
	fmt.Println(animal)

	smaug := CatName("Smaug")
	animal = &smaug
	fmt.Println(animal.MakeNoise())
	fmt.Println(animal.MakeNoise())
}

func (a *DogName) MakeNoise() string {
	return string(*a) + " Say: Bork!"
}

func (c *CatName) MakeNoise() string {
	return string(*c) + " Say: Miaw!"
}

func (a DogName) String() string {
	return fmt.Sprintf("%v is a Dog, and %v is a very good boye.\n", string(a), string(a))
}

// Error Handling
type CustomInternalError struct {
	Name string
	Desc string
}

type CustomIntegerError float64

//	func (e CustomInternalError) Error() string {
//		return fmt.Sprintf("Error: %v\nDescription: %v\n", e.Name, e.Desc)
//	}
func (e CustomIntegerError) Error() string {
	return fmt.Sprintf("Error: %v\n", float64(e))
	// return fmt.Sprintf("Error: %v\nDescription: %v\n", e)
}

func stableFunction(x float64) (float64, error) {
	if x < 0 {
		return 0, CustomIntegerError(x)
	} else {
		return x + 10, nil
	}
}

// Stream reader
func readTextStream() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// Goroutines
func eatFood(foodLeft *int, participant *string, score *int) {
	for *foodLeft > 0 {
		if *foodLeft > 0 {
			time.Sleep(100 * time.Millisecond)
			*foodLeft--
			if *foodLeft <= -1 {
				*score = -1;
				fmt.Printf("%v tried to eat a food but there is none left, they starved and died.\n", *participant)
			} else {
				*score++;
				fmt.Printf("%v has eaten 1 food: %v food left remaining.\n", *participant, *foodLeft)
			}
		} else {
			fmt.Printf("%v stopped eating.\n", *participant)
		}
	}
}

// Channel
func sumChannel(s []int, c chan int) {
	// Write channel function
}
