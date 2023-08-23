package main
import "fmt"

var myMap = map[int]string{1: "First", 2: "Second", 3: "Third", 4: "Fourth"} // declare map

var makedMap = make(map[string]string) // declare map using make function

func main(){
	makedMap["brand"] = "Lenovo"
	makedMap["procsessor"] = "Intel Celeron"

/* 	val1, ok1 := myMap[1]
	check1, ok2 := myMap[5]
	println(val1, ok1)
	println(check1, ok2) */
	// checking map elements using their names

	// otherMap := myMap
	// fmt.Println(myMap)
	// otherMap[5] = "Fifth"
	// println("myMap changed!")
	// fmt.Println(otherMap)

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	// fmt.Printf("\t%v\n", makedMap) // show map
}