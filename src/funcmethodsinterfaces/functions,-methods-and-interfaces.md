name: Functions, Methods and Interfaces
video: 
description: 
level: Beginner
topic: Go
# Functions, Methods and Interfaces
## Functions, Methods and Interfaces

---
name: Functions
video: 
thumb:
github:
# Functions
## 

Functions are declared with the "func" keyword.
Functions have a name, optional input parameters, and optional return values.

### Function Examples
```
cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/funcs
go run main.go
```

Functions are first-class types in Go. You can assign a function to a variable, you can pass functions as parameters.
### Function Values
```
cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/funcvalues
go run main.go
```---
name: Methods1
video: 
thumb:
github:
# Methods
## 


Methods are syntactic sugar for a function with a type as the first parameter.

WHAT?

	func ChangeEmail(u *User, newEmail string) { ... } // Ugly

	func (u *User) ChangeEmail(newEmail string) { ... } // Clear

These are equivalent in functionality, but one of them is much more clear.
---
name: Methods2
video: 
thumb:
github:
# Methods
## 

Methods can be declared on any named type. Use a pointer receiver when you want to modify the existing type. Use a value receiver when you don't need to modify the type.

### Method Example
```
cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/method
go run main.go
```

Methods are First Class Citizens in Go

That means you can create variables of type method, assign to them, and operate on those variables.

### First Class Methods 
```
cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/firstmethod
go run main.go
```---
name: Interfaces
video: 
thumb:
github:
# Interfaces
## 

Interfaces allow you to specify BEHAVIOR. *If something can do this, then it can be used here.*  Interfaces are types. They are declared as types.

Interfaces usually have a very small number of Methods, 1 or 2 is most common.

	The larger the interface, the weaker the abstraction. -- Rob Pike

*Interface names try to describe the action.*
---
name: Interface Examples
video: 
thumb:
github:
# Interface Examples
## 


Stringer - a type that has a method that returns a string

io.Writer - a type that has a method that writes to a buffer

io.ReadCloser - a type that has a method that reads from a stream and closes it when done
---
name: Creating Good Interfaces
video: 
thumb:
github:
# Creating Good Interfaces
## 


Good interfaces define a very small set of specific actions:

- Writing bytes to a buffer (io.Writer)
- Returning a String representing a type (fmt.Stringer)
---
name: Standard Library Interfaces
video: 
thumb:
github:
# Standard Library Interfaces
## 

Examples of Interfaces in Go's standard library:

- [db/sql: connection interface](https://golang.org/pkg/database/sql/driver/#Conn) 
- [encoding: Marshaler Interfaces](https://golang.org/pkg/encoding/) 
- [net/http: HTTP Handler Interface](https://golang.org/pkg/net/http/#Handler)

By these patterns, you can see that interfaces are intended to represent a small set of behaviors.
---
name: Interfaces in Practice
video: 
thumb:
github:
# Interfaces in Practice
## 

### Interface Example
```
cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/interfaces
go run main.go
```---
name: The Empty Interface
video: 
thumb:
github:
# The Empty Interface
## 

All the interfaces we've seen up to now have declared one or more functions. Interfaces don't have to declare any functions, though.

If you declare an interface with an empty set of functions, then any type will satisfy that interface.

In Go we use the empty interface to represent "anything".

### Empty Interface Example
```
cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/empty
go run main.go
```---
name: Type Assertions
video: 
thumb:
github:
# Type Assertions
## 

That's pretty powerful. You can pass any type around without losing the type safety. But when you receive an "interface{}" how do you know what to do with it?

### Type Assertion Example
```
cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/demos/assert
go run main.go
```---
name: Exercise-fncmthdsint
video: 
thumb:
github:
# Exercise
## 

Fix the last example to recognize and print the float value.

	cd $GOPATH/src/github.com/gophertrain/material/funcmethodsinterfaces/exercises/assert
	go run main.go
