package main
import (
	"fmt"
	"errors"
	"time"
)

func main() {
	testCount(16)
}

func testCount(num int){
	if num == 16 {
		fmt.Println(time.Now().Second(), errors.New("Unexpected data"))
	} else {
		fmt.Println("True")
	}
}