package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("slicen -*", nums, len(nums))

	//agregar datos

	nums = append(nums, 6)
	nums = append(nums, 7)

	fmt.Println("slicen -*", nums, len(nums))

	//sub slicen

	subslicen := nums[:2]

	nums[0] = 10

	fmt.Println("subslicen -*", subslicen, len(subslicen))
	fmt.Println("slicen -*", nums, len(nums))

	//capaciad
	//longitud
	//puntero

	meses := []string{"ene", "feb", "mar"}
	fmt.Printf("len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "aabr", "re", "por")
	fmt.Printf("len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}
