package main

import "fmt"

const englishHelloTemplate = "Hello, %v"
const defaultName = "world"

func Hello(name string) string {
	if name == "" {
		name = defaultName
	}
	return fmt.Sprintf(englishHelloTemplate, name)
}

func main() {
	fmt.Println(Hello("Adam"))
}
