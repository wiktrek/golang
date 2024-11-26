package main

// bubble sort
import (
	"fmt"
	"math/rand/v2"
)

func main() {
	lista := generate_nums(1000000)
	a := 0
	swapped := false
	for i:=0;i < len(lista); i++ {
		swapped = false
		for j:=0; j < len(lista)-i-1; j++ {
			if lista[j] > lista[j+1] {
				a = lista[j]
                lista[j] = lista[j+1]
                lista[j+1] = a
                swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Print(lista)
}
func generate_nums(nums int) []int{
	lista := make([]int, nums) 
	for i := 0; i < nums; i++ {
		lista[i] = rand.IntN(nums)
	}
	return lista[:]
}


