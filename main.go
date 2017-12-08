package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 1
func main() {

	// Create channel wthat will accept string and have buffer size of >4

	//How many strings can hold it before consuming it
	stringChanel := make(chan string, 4)

	//Simple string array
	names := []string{"Andrius", "Dominykas", "Simas", "Mindaugas"}

	for _, name := range names {
		stringChanel <- name
	}
	close(stringChanel)

	for n := range stringChanel {
		fmt.Println(len(stringChanel))
		fmt.Println(n)
	}

}
