package main

import "fmt"

func main() {
	i := 5
	b := &i
	*b = 6
	fmt.Println(i)
	return
}
