package main

import "fmt"

func main() {
	var nums [4]int
	fmt.Println(nums)

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	fmt.Println(nums)
	fmt.Println(nums[3])

	//arrays con valores strings

	numes := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(numes)

	colors := [...]string{
		"red",
		"green",
		"blue",
		"white",
	}
	fmt.Println(colors, len(colors))

	//indices definidos
	monedas := [...]string{0: "Us", 2: "soles", 3: "pesos", 6: "euro"}

	fmt.Println(monedas, len(monedas))

	//sub arrays

	subMoneda := monedas[2:]
	fmt.Println("sub aarray", subMoneda, len(subMoneda))

}
