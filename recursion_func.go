package main
import "fmt"

func main(){
	println(factorialTest(4))
}

func testCount(x int) int{
	if x == 16 {
		return 0
	}
	fmt.Println(x)

	return testCount(x + 1)
}


func factorialTest(x float64) (y float64){
	if x > 0 {
		y = x * factorialTest(x-1)
	} else {
		y = 1
	}
	return
}
