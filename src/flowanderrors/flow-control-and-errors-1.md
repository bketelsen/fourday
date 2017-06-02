name: Flow Control and Errors
video: 
description: 
level: Beginner
topic: Go
# Flow Control and Errors
## Flow Control and Errors

---
name: Flow Control and Errors
video: 
thumb:
github:
# Flow Control and Errors
## 

Go's control flow is very similar to C, but it has a few simplifications.

Go's if statements don't require parenthesis. They're easier to read, but you can add the parens if you want them for clarity in more complex statements.

### If Examples
```
cd $GOPATH/src/github.com/gophertrain/material/flowanderrors/demos/if/
go run main.go
```---
name: If Statements
video: 
thumb:
github:
# If Statements
## 

It's idiomatic in go to skip the else where possible:

### Skipping Else
```
cd $GOPATH/src/github.com/gophertrain/material/flowanderrors/demos/skipelse/
go run main.go
```---
name: For
video: 
thumb:
github:
# For
## 

There's a Go joke that goes like this:

How do you spell "WHILE" in go? F O R

There is one looping construct, the "for" statement.

### For
```
cd $GOPATH/src/github.com/gophertrain/material/flowanderrors/demos/for/
go run main.go
```---
name: SWITCH
video: 
thumb:
github:
# SWITCH
## 

Go's switch statement is very powerful. Cases are evaluated top to bottom until the first match is found.

### Switch
```
cd $GOPATH/src/github.com/gophertrain/material/flowanderrors/demos/switch/
go run main.go
```---
name: Type Switching
video: 
thumb:
github:
# Type Switching
## 

One of the more common uses of the switch statement is when determining the type of an interface{} value. 

### Type Switch
```
cd $GOPATH/src/github.com/gophertrain/material/flowanderrors/demos/typeswitch/
go run main.go
```