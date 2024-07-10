package main

import "fmt"

// Greeter Interface implements LanguageName method that returns the Language Name and Greet method that returns a greeting string
type greeter interface {
	languageName() string
	greet(string) string
}

// Italian empty Struct
type Italian struct {
}

// Portuguese empty Struct
type Portuguese struct {
}

// greet method for Italian
func (i Italian) greet(name string) string {
	return fmt.Sprintf("Ciao %v!", name)
}

// Greet method for Portuguese
func (i Portuguese) greet(name string) string {

	return fmt.Sprintf("Olá %v!", name)
}

// languageName method for Italian struct
func (i Italian) languageName() string {
	return "Italian"
}

// languageName method for Portuguese struct
func (i Portuguese) languageName() string {
	return "Portuguese"
}

// SayHello function returns a string in format : => "I can speak German: Hallo Dietrich!"
func SayHello(name string, g greeter) string {
	return fmt.Sprintf("I can speak %v: %v", g.languageName(), g.greet(name))
}

// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
