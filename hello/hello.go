package main

import "fmt"

const enHelloPrefix = "Hello, "
const spHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "

// Hello world
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "french":
		prefix = frHelloPrefix
	case "spanish":
		prefix = spHelloPrefix
	default:
		prefix = enHelloPrefix
	}

	return prefix
}

func init() {
	fmt.Println("Init Hello")
}

func main() {
	fmt.Println(Hello("Jun", ""))
}
