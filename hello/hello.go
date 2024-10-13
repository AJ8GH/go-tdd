package main

import "fmt"

const english = "English"
const spanish = "Spanish"
const englishHelloTemplate = "Hello, %v"
const spanishHelloTemplate = "Hola, %v"
const defaultName = "world"

var languages = map[string]string{
	english: englishHelloTemplate,
	spanish: spanishHelloTemplate,
}

func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}
	return fmt.Sprintf(languages[language], name)
}

func main() {
	fmt.Println(Hello("Adam", "English"))
}
