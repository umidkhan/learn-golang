package main

import (
	"fmt"
)

// import "time"

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

	/* t := time.Now()
	switch {
	case t.Hour() < 12:
		println("It's before noon")
	default:
		println("It's after noon")
	} */

	name := "umidxon"
	switch {
	case name == "Umidxon", name == "umidxon":
		println("Welcome", name)
	case name == "Anyone":
		println("Actually you're anyone, get out here now!")
	default:
		println("Name is incorrect, try again!")
	}
}

func WhatAmI(i interface{}){
	switch t := i.(type) {
	case bool:
		println("I'm boolean")
	case int:
		println("I'm int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}

func main() {
	WhatAmI("Hi")
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
