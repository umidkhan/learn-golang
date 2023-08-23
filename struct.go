package main

type Car struct {
	name string;
	processor string;
	ram int;
	isGrahpics bool;
	price int;
}

func main(){
	var pc1 Car

	pc1.name = "ASUS Zenbook"
	pc1.processor = "Intel Core i9"
	pc1.ram = 16
	pc1.isGrahpics = true
	pc1.price = 900

	println("PC Name: " + pc1.name)
	println("Processor: " + pc1.processor)
	println("Is grahpic card:", pc1.isGrahpics)
	println("RAM:", pc1.ram)
	println("Price", pc1.price)
}