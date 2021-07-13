/*

Interfaces are named collections of method signatures.
Interfaces are not generic types
In GO Interfaces are tough --- only need understand and practice

Problems with out Interfaces(Refer below Example)  :- we cannot create a function with same name like
		func (englishBot) getGreeting() string{}
		func (englishBot) getGreeting() string{} when we create like this it throws an error called already func name
declared to achieve this we use interfaces
like java method overloading is not possible in go lang (Refer below Example)

Refer below Example for Problems with out Interfaces

package main

import "fmt"

type englishBot struct {}
type  spanishBot struct {}
func main()  {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hi There !!! "
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}
func (spanishBot) getGreeting() string {
	return "Hola !!! "
}
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
*/

package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {
	name string
}
type spanishBot struct{}

func main() {
	eb := englishBot{"English bot"}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (eb englishBot) getGreeting() string {
	return eb.name
}
func (spanishBot) getGreeting() string {
	return "Hola !!! "
}
