name: Embedding and Composition
video: 
description: 
level: Beginner
topic: Go
# Embedding and Composition
## Embedding and Composition

---
name: Embedding and Composition
video: 
thumb:
github:
# Embedding and Composition
## 

	
Go doesn't do the typical Object Oriented inheritance. Instead Go lets you "borrow" smaller pieces of functionality and embed them into larger types. This is called Composition.

### Composition
```
	cd $GOPATH/src/github.com/gophertrain/material/embeddingcomposition/demos/embedding/
	go run main.go
```---
name: Composition
video: 
thumb:
github:
# Composition
## 

Composition works with Interfaces too:

### Embedded Types Interface Example
```
	cd $GOPATH/src/github.com/gophertrain/material/embeddingcomposition/demos/interfaces/
	go run main.go
```---
name: Composition2
video: 
thumb:
github:
# Composition
## 

Put composition, interfaces and embedding together to get a powerful abstraction:

### Composition, Interfaces, Embedding Example
```
	cd $GOPATH/src/github.com/gophertrain/material/embeddingcomposition/demos/both/
	go run main.go
```---
name: Important
video: 
thumb:
github:
# Important
## 


This concept is crucial to your success as a Go programmer.

Interfaces aren't explicitly declared, they're checked at compile time. You don't have to announce that a Crowbar is a NailPuller, the program won't compile if you try to pass a Crowbar to a function that requires a NailPuller if the Crowbar doesn't implement the PullNail function.

Using interfaces is a confusing change from traditional OO programming concepts. After some time, you'll find them refreshing. You create types that implement behaviors, and create functions that operate on things that implement those behaviors.

Let's go back and look at that example again and reason out loud how we'd do that in Java/Ruby with abstract classes and interfaces.

It's a lot easier to keep the Go version in your head.
---
name: Recap
video: 
thumb:
github:
# Recap
## 

In Go, we don't create an inheritance chain of functionality:

Transport -> Wheeled -> Car -> Sedan

We create behaviors and apply them to types:

### Behaviors
```
	cd $GOPATH/src/github.com/gophertrain/material/embeddingcomposition/demos/behaviors/
	go run main.go
```---
name: Exercise-embdcomp
video: 
thumb:
github:
# Exercise
## 

Modify the *behaviors* demo to add another type of transportation, implementing the Storer interface.


### Behaviors
```
	cd $GOPATH/src/github.com/gophertrain/material/embeddingcomposition/exercises/behaviors/
	go run main.go
```