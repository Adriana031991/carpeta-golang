package main

import "fmt"

func main() {
	names := make([]string, 4, 5)

	names[0] = "100"
	names[1] = "100"
	names[2] = "100"
	names[3] = "100"

	names = append(names, "0", "500", "500")

	fmt.Println("names-*", names, len(names), cap(names))
}
