package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

func SwitchExample() {
	/* 	i := 5
	   	println("Write", i, "as")
	   	switch i {
	   	case 1:
	   		println("one")
	   	case 2:
	   		println("two")
	   	case 3:
	   		println("three")
	   	case 4:
	   		println("four")
	   	default:
	   		println("not added")
	   	} */

	/* 	switch time.Now().Weekday() {
	   	case time.Saturday, time.Sunday:
	   		println("It's the weekend")
	   	default:
	   		println("it's the weekday")
	   	} */

	weekday := time.Now().Weekday()
	switch weekday {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Throusday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Error")
	}

	/* t := time.Now()
	switch {
	case t.Hour() < 12:
		println("It's before noon")
	default:
		println("It's after noon")
	} */

	// name := "umidxon"
	// switch {
	// case name == "Umidxon", name == "umidxon":
	// 	println("Welcome", name)
	// case name == "Anyone":
	// 	println("Actually you're anyone, get out here now!")
	// default:
	// 	println("Name is incorrect, try again!")
	// }
}

func WhatAmI(i interface{}) {
	switch t := i.(type) {
	case bool:
		println("I'm boolean")
	case int:
		println("I'm int")
	case string:
		println("I'm string")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}

func ArrayExample() {
	/* var a [5]int
	fmt.Println("array:", a)

	a[0] = 100
	fmt.Println("set:", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b) */

	var myArray = [...]interface{}{1, "two", 3.14, true}

	fmt.Println(myArray)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}

func SliceExample() {
	/* var s []string
	s = make([]string, 3)
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	fmt.Println(s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s) */

	// var scores = []int{10, 20, 30}
	// scores[2] = 30 / 2
	// scores = append(scores, 40)

	mySlice := make([]string, 3)
	firstRange := mySlice[1:2]
	fmt.Println(firstRange)

	// fmt.Println(scores, len(scores))
	// fmt.Printf("%T\n", scores)
}

func MapExample() {
	statuses := make(map[string]int)

	// Add values to map
	statuses["successful"] = 1
	statuses["fail"] = 2
	statuses["pending"] = 3

	// read map value
	var SuccesfulStatus = statuses["successful"]
	fmt.Println(SuccesfulStatus)
	fmt.Println(statuses["fail"])
}

func ForExample() {
	for i := 1; i < 8; i++ {
		fmt.Println(i)
	}

	// array itaratsiya
	myArr := [3]string{"Go", "C#", "C++"}
	for index, value := range myArr {
		fmt.Println("Index:", index, "Value:", value)
	}

	// map itaratsiya
	myMap := map[int]string{
		1: "Python",
		2: "JavaScript",
		3: "C",
	}

	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func Sum(x, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Manfiy son qabul qilinmaydi!")
	}
	return math.Sqrt(x), nil
}

type Car struct {
	Make, Model, Color string
	Year, Weight       int
	Engine             engine
}

type engine struct {
	name string
	hp   int
}

func main() {
	// var myCar Car
	myCar := Car{"BMW", "X1 M35i", "white", 2023, 1719, engine{"xDrive28i", 312}}
	fmt.Println(myCar)

	// result, err := sqrt(-1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }
	// fmt.Println(Sum(15, 15))
	// ForExample()
	// MapExample()
	// ArrayExample()
	// SwitchExample()
	// SliceExample()
	/* 	fmt.Println("Hello World!")

	   	var a string = "Golang" // o'zgaruvchi e'lon qilish
	   	b := 12 // qisqartma shaklda o'zgaruvchi e'lon qilish
	   	println(a)
	   	println(b) */

	/* const n = 10000 // o'zgarmas o'zgaruvchi e'lon qilish
	const d = 3e20 / n
	println(d)
	println(int64(d))
	println(math.Sin(n)) */

	/* i := 1
	for i <= 10 { // for sintaksis
		println(i)
		i += 1
	} */

	/* for {
		println("Infinite!")
	} */

	/* for n := 0; n <= 5; n++ {
		if n%2 == 0 {
		}
		println(n)
	} */

	// num1 := 8
	// num2 := 3

	// if num1%num2 == 0 {
	// 	println(num1, "soni", num2, "soniga qoldiqsiz bo'linadi")
	// } else if num1%num2 == 1 {
	// 	println("qoldiq 1")
	// } else {
	// 	println(num1, "soni", num2, "soniga qoldiqli bo'linadi")
	// }
}
