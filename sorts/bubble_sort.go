//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func bubbleSort(arrayzor []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arrayzor)-1; i++ {
			if arrayzor[i+1] < arrayzor[i] {
				arrayzor[i+1], arrayzor[i] =  arrayzor[i], arrayzor[i+1]
				swapped = true
			}
		}
	}
	return arrayzor
}

func swap(arrayzor []int, i, j int) {
	arrayzor[i],arrayzor[j] = arrayzor[j],arrayzor[i]
}
