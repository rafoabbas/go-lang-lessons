package main

import "fmt"

func main() {
	//message := "Hello World"
	//changeVariable(&message)
	fmt.Print(sum(5, 6))
}

func sum(x, y int) int {
	return x + y
}

func sayHelloName(name string) {
	println("Hello", name)
}

func sayHello() {
	println("Hello")
}

// globaldan deyismek
func changeVariable(message *string) {
	println(*message)
	*message = "Hello Go"
}
