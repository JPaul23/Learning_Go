package main

import (
	"fmt"
)

func main() {
	// fmt.Println(combat(100.0, 50))
	// fmt.Printf("Answer is === %v\nCharacter is === %c \nCode point is === %U \n", answer, answer, answer)

	// testing interfaces
	var shape Shaper
	shape = myvalue{5, 10}
	fmt.Println("The value of interface is:", shape)
	fmt.Println("Area of shape is:", shape.Area())
	fmt.Println("Perimeter of shape is:", shape.Perimeter())
}

// finding average
func GetGrade(a, b, c int) rune {
	// Code here
	average := (a + b + c) / 3
	//using switch cases
	switch {
	case average >= 90 && average <= 100:
		return 'A'
	case average >= 80 && average < 90:
		return 'B'
	case average >= 70 && average < 80:
		return 'C'
	case average >= 60 && average < 70:
		return 'D'
	default:
		return 'F'
	}
}

// finding the opposite
func Opposite(value int) int {
	return -value
}

// repeating string
func RepeatStr(repetitions int, value string) string {
	var repeated string
	for i := 0; i < repetitions; i++ {
		repeated += value
	}
	return repeated
}

// combat function
func combat(health, damage float64) float64 {
	newHealth := health - damage
	if newHealth < 0 {
		return 0
	}
	return newHealth
}

// Go Interfaces
type Shaper interface {
	// method signature
	Area() float64
	Perimeter() float64
}

type myvalue struct {
	radius float64
	height float64
}

// implementing the method shaper interface
func (m myvalue) Area() float64 {
	return m.radius * m.radius * 3.14
}

func (m myvalue) Perimeter() float64 {
	return 2 * m.radius * 3.14
}
