package main

import "fmt"

func ex1() {
	count := 100
	for i := 1; i < count+1; i++ {
		fmt.Print(i, " ")
	}
}

func ex2() {
	count := 101
	var soma int
	for i := 1; i < count; i++ {
		soma += i
		fmt.Println(soma)
	}
}

func ex3() {
	count := 51
	for i := 1; i < count; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}

func ex4() bool {
	var num int
	fmt.Scan(&num)

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func ex5() {
	seq := [20]int{0, 1, 1}

	for i := 3; i < 20; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}

	fmt.Print(seq)
}

/*
* @description resposta alternativa do exercício 5 usando um "slice" ao invés do array
* pros: podemos definir o tamanho e, caso quisermos, podemos usar um slice sem tamanho pré definido
* cons: usa mais memória e não podemos inicializar os valores de maneira pré definida
 */
func ex5Alt() {
	seqSize := 20
	seq := make([]int, seqSize)
	seq[0] = 0
	seq[1] = 1
	seq[2] = 1

	for i := 3; i < seqSize; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}

	fmt.Print(seq)
}

func ex6() {
	// Fatorial... to be continued :p
}
