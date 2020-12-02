package main

import "fmt"

func main() {
	const name = "world"
	sayHello(name)
}

func sayHello(name string) {
	fmt.Println("Hello,", name, "!")
}

func sayBye(name string) {
	fmt.Println("Bye,", name, "!")
}
