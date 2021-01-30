package main

import "fmt"

type Animal interface {
	Eats() string
	Says() string
}

type Cow struct {}
func (*Cow) Eats() string {
	return "grass"
}
func (Cow) Says() string {
	return "moo"
}

type Cat struct {}
func (*Cat) Eats() string {
	return "mice"
}
func (*Cat) Says() string {
	return "meow"
}

type Dog struct {}
func (Dog) Eats() string {
	return "bones"
}
func (Dog) Says() string {
	return "woof"
}

func Speak(a Animal) {
	fmt.Println(a.Says())
}

func main() {

	// Type Cow implements the Animal interface methods using one pointer receiver and one value receiver.
	// POINTERS of type Cow support both value receivers and pointer receivers in their method sets.
	// VALUES support ONLY value receivers

	// So this will work...
	bessy := &Cow{} // pointer supports both pointer and value receivers for interface methods
	Speak(bessy)

	// But this won't compile...
	// daisy := Cow{} // value only supports value receivers, but cow has one pointer receiver
	// Speak(daisy)

	// Cat uses only pointer receivers to implement the Animal interface so this is ok
	pussy := &Cat{}
	Speak(pussy)
	// felix := Cat{} // again, this won't compile

	// Dog uses only value receivers, so can use pointer or value for concrete type
	fido := Dog{}
	Speak(fido)
	rex := &Dog{}
	Speak(rex)

}
