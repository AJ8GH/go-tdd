package main

import "fmt"

const (
	defaultName = "world"

	french  = "French"
	spanish = "Spanish"

	englishHelloTemplate = "Hello, %v"
	frenchHelloTemplate  = "Bonjour, %v"
	spanishHelloTemplate = "Hola, %v"
)

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}
	return fmt.Sprintf(getTemplate(language), name)
}

func main() {
	fmt.Println(Hello("Adam", "English"))
}

func getTemplate(language string) string {
	switch language {
	case spanish:
		return spanishHelloTemplate
	case french:
		return frenchHelloTemplate
	default:
		return englishHelloTemplate
	}
}
