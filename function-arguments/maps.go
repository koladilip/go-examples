package main

import (
	"fmt"
)

func incMaps(people map[string]int) map[string]int {
	for i, v := range(people) {
		people[i] = v+1
	}
	return people
}

func main() {
	people := map[string]int{}
	people["dilip"] = 1
	people["agnel"] = 2
	fmt.Println("before", people)
	fmt.Println("returned", incMaps(people))
	fmt.Println("after", people)
}
