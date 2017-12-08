package main

import (
	"fmt"
)

func main() {

	names := []string{"Andrius", "Dominykas", "Simas", "Mindaugas"}

	for _, name := range names {
		fmt.Println(name)
	}
}
