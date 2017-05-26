package main

import "time"

func expensiveProcess(done chan bool) {
	// simulate long task
	time.Sleep(5 * time.Second)
	done <- true
}

func main() {
	done := make(chan bool)
	go expensiveProcess(done)
	<-done
}
