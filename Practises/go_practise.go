package main

import (
	"fmt"
)

func myfunc(ch chan int) {

	fmt.Println(200 + <-ch)
}

func main() {
	fmt.Println("Start main method====")

	myChannel := make(chan int)
	go myfunc(myChannel)
	myChannel <- 23

	fmt.Println("End main method====")
}
