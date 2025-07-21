package main

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName(name string) string
	Greet(visitorName string) string
}

type German struct {
	languageName string
	visitorName  string
}

func (german German) LanguageName(name string) string {
	german.languageName = name
	return german.languageName
}

func (german German) Greet(visitorName string) string {
	german.visitorName = visitorName
	return fmt.Sprintf("I can speak German: Hallo %s!", german.visitorName)
}

type Italian struct {
	languageName string
	visitorName  string
}

func (italian Italian) LanguageName(name string) string {
	italian.languageName = name
	return italian.languageName
}

func (italian Italian) Greet(visitorName string) string {
	italian.visitorName = visitorName
	return fmt.Sprintf("I can speak Italian: Ciao %s!", italian.visitorName)
}

type Portuguese struct {
	languageName string
	visitorName  string
}

func (portuguese Portuguese) LanguageName(name string) string {
	portuguese.languageName = name
	return portuguese.languageName
}

func (portuguese Portuguese) Greet(visitorName string) string {
	portuguese.visitorName = visitorName
	return fmt.Sprintf("I can speak Portuguese: Olá %s!", portuguese.visitorName)
}

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet(name)
}

func main() {
	germanGreeter := German{}
	italianGreeter := Italian{}
	portugueseGreeter := Portuguese{}

	fmt.Println(SayHello("Dietrich", germanGreeter))
	// => "I can speak German: Hallo Dietrich!"

	fmt.Println(SayHello("Giulia", italianGreeter))
	// => "I can speak Italian: Ciao Giulia!"

	fmt.Println(SayHello("Carlos", portugueseGreeter))
	// => "I can speak Portuguese: Olá Carlos!"
}
