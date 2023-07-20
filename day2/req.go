package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"

func HelloReq(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(HelloReq("", "Spanish"))
}
