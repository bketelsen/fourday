// Declare an interface named speaker with a method named sayHello. Declare a struct
// named English that represents a person who speaks english and declare a struct named
// Chinese for someone who speaks chinese. Implement the speaker interface for each
// struct using a value receiver and these literal strings "Hello World" and "你好世界".
// Declare a variable of type speaker and assign the _address of_ a value of type English
// and call the method. Do it again for a value of type Chinese.
//
// Add a new function named sayHello that accepts a value of type speaker.
// Implement that function to call the sayHello method on the interface value. Then create
// new values of each type and use the function.
package main

import "fmt"

// speaker implements the voice of anyone.
type speaker interface {
	sayHello()
}



// english represents an english speaking person.
type english struct{
	name string

}

// sayHello implements the speaker interface.
func (e english) sayHello() {
	fmt.Println("Hello ", e.name)
}

// chinese represents a chinese speaking person.
type chinese struct{
	morning bool
}

// sayHello implements the speaker interface.
func (c chinese) sayHello() {
	if c.morning {
		fmt.Println("Good morning in Chinese")
		return
	}

	fmt.Println("你好世界")
}

// main is the entry point for the application.
func main() {
	// Declare a variable of the interface type.
	var sp speaker

	// Assign a value to the interface type variable and
	// call the interface method.
	var e english
	e.name = "William"
	sp = e
	sp.sayHello()

	// Assign a different value to the interface type
	// variable and call the interface method.
	var c chinese
	c.morning = true
	sp = c
	sp.sayHello()

	// Create new values and call the function.
	sayHello(&english{name: "Godfrey"})
	sayHello(&chinese{morning: false})

	// The use of new() and the empty literal is there
	// as a talking point about these options.
}

// SatHello abstracts speaking functionality.
func sayHello(sp speaker) {

	sp.sayHello()
}
