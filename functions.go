package main

func main() {
	message := "Hello World"
	changeVariable(&message)
	println(message)
}

func sayHelloName(name string) {
	println("Hello", name)
}

func sayHello() {
	println("Hello")
}

func changeVariable(message *string) {
	println(*message)
	*message = "Hello Go"
}
