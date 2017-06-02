name: Concurrency
video: 
description: 
level: Beginner
topic: Go
# Concurrency
## Concurrency

---
name: Goroutines
video: 
thumb:
github:
# Goroutines
## 


A goroutine is a function that is capable of running concurrently with other functions.

Goroutines aren't threads. They're smaller, lighter, and less resource intensive. A goroutine only requires a few bytes of memory, compared to a thread. Each thread has a fixed stack size, so there's a finite limit to the number of threads you can run. In contrast, it's common to run millions of goroutines at once.
---
name: Goroutines2
video: 
thumb:
github:
# Goroutines
## 

To create a goroutine, simply prepend the "go" keyword before a function call.

### Create Goroutine
```
	cd $GOPATH/src/github.com/gophertrain/material/concurrency/demos/goroutine/
	go run main.go
```
Why didn't that work?
---
name: Goroutines3
video: 
thumb:
github:
# Goroutines
## 

Goroutines execute concurrently. In this case, `main()` and `sayHello()` were running at the same time. `main()` finished and `sayHello()` never got a chance to run!

To solve this problem you need to use some sort of communication or synchronization device to signal that a goroutine is complete.

We typically do that with channels or waitgroups, depending on whether you need to return a value from the function you're running.
---
name: Waitgroups
video: 
thumb:
github:
# Waitgroups
## 

WaitGroups are defined in the "sync" package in the standard library. They're a simple and powerful way to ensure that all of the goroutines you launch have completed before you move on to do other things.

### Waitgroup Example
```
	cd $GOPATH/src/github.com/gophertrain/material/concurrency/demos/waitgroup/
	go run main.go
```

WaitGroups are also especially powerful when combined with closures (anonymous functions)

[Closure Example](https://golang.org/pkg/sync/#example_WaitGroup)
---
name: Channels
video: 
thumb:
github:
# Channels
## 

Channels are a typed conduit through which you can send and receive values.

[Channel Overview](https://tour.golang.org/concurrency/2)
---
name: Channels2
video: 
thumb:
github:
# Channels
## 

By default channels are unbuffered -- sends on a channel will block until there is something ready to receive from it.

You can create a buffered channel- which allows you to control how many items can sit in a channel unread.

[Buffered Channels](https://tour.golang.org/concurrency/3)
---
name: Channels3
video: 
thumb:
github:
# Channels
## 

You can signal that there are no more values in a channel by closing it. Receivers can test to see if a channel is closed during a read.

[Closing Channels](https://tour.golang.org/concurrency/4)
---
name: Channels4
video: 
thumb:
github:
# Channels
## 

You can use a select statement to allow a goroutine to operate on multiple communication operations:

[Select](https://tour.golang.org/concurrency/5)
---
name: Blocking Channels
video: 
thumb:
github:
# Blocking Channels
## 

If no operations inside a select statement can process, the select will block. To work around this problem, you can add a default statement to your select:

[Select With Default](https://tour.golang.org/concurrency/6)
---
name: Exercise-concurrency
video: 
thumb:
github:
# Exercise
## 

Read Dave Cheney's article on some of the more interesting properties of channels.

[Curious Channels](https://dave.cheney.net/2013/04/30/curious-channels)
