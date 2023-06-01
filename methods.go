package main

import "fmt"

type Human struct {
	Name string
	Age  int
	say  func(string) string
}

type Person struct {
	Name string
	Age  int
}

func newPerson() *Person {
	p := new(Person)
	return p
}

func main() {
	//say := func(message string) string {
	//	return "Hello"
	//}

	//rauf := Human{Name: "Rauf", Age: 27, say: say}
	//rustem := Human{Name: "Rustam", Age: 25, say: say}

	//fmt.Println(rauf.say("Hello"))

	x := newPerson()
	x.Age = 10
	fmt.Println(x.Age)
}
