package main

import (
	"fmt"
)

func main() {

	stringChanel := make(chan string, 4)

	names := []string{"Andrius", "Dominykas", "Simas", "Mindaugas"}

	for _, name := range names {
		stringChanel <- name
	}

	for n := range stringChanel {
		fmt.Println(n)
	}
}
