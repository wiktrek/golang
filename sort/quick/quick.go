package main

// bubble sort
import (
	"fmt"
	"math/rand/v2"
)

func main() {
	lista := generate_nums(100)
	lista = sort(lista, 0, len(lista)-1)
	fmt.Print(lista)
}
func partition(lista []int, low int, high int) ([]int, int) {
	pivot := lista[high]
	i := low
	for j := low; j < high; j++ {
		if lista[j] < pivot {
			lista[i], lista[j] = lista[j], lista[i]
			i++
		}
	}
	lista[i], lista[high] = lista[high], lista[i]
	return lista,i;
}
func sort(arr []int, low int, high int) []int{
	if low < high {
		var p int
       	arr,p = partition(arr, low, high)
        arr = sort(arr, low, p-1)
        arr = sort(arr, p+1, high)
    }
	return arr
}
func generate_nums(nums int) []int{
	lista := make([]int, nums) 
	for i := 0; i < nums; i++ {
		lista[i] = rand.IntN(nums)
	}
	return lista[:]
}

