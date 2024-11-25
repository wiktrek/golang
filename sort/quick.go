package main

// bubble sort
import (
	"fmt"
	"math/rand/v2"
)

func main() {
	lista := generate_nums(1000000)
	swapped := false
	
	fmt.Print(lista)
}
func generate_nums(nums int) []int{
	lista := make([]int, nums) 
	for i := 0; i < nums; i++ {
		lista[i] = rand.IntN(nums)
	}
	return lista[:]
}

