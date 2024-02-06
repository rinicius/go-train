package main

import "fmt"

func ex2() {
	var vetor [10]int
	for i := 0; i < len(vetor); i++ {
		vetor[i] = i + 1
	}

	fmt.Println(vetor)

	//

	vetor2 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	fmt.Println(vetor2)

}
