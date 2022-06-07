package main

import "fmt"

const enHelloPrefix = "Hello, "
const espHelloPrefix = "Hola, "
const frHelloPrefix = "Bonjour, "
const exclamation = "!"
const spanish = "Spanish"
const french = "French"

func main() {
	fmt.Println(Greet("", ""))
	fmt.Println(Greet("Chris", ""))
	fmt.Println(Greet("Elodie", "Spanish"))
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frHelloPrefix
	case spanish:
		prefix = espHelloPrefix
	default:
		prefix = enHelloPrefix
	}

	return
}

func Greet(name, lang string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name + exclamation
}
