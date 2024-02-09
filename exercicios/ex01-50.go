package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

func ex4(num int) bool {

	for i := 2; i*i <= num; i++ {
		if math.Mod(float64(num), float64(i)) == 0 {
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
	var num int
	fmt.Println("Digite o número: ")
	fmt.Scan(&num)

	for i := num - 1; i > 0; i-- {
		num *= i
	}

	fmt.Println("Resultado: ", num)
}

func ex7() {
	// 7. Implemente um programa que calcule o máximo divisor comum (MDC) de dois números.
	var num1 int
	var num2 int

	fmt.Scan(&num1)
	fmt.Scan(&num2)

	div := 1

	for div != 0 {

		if num2 > num1 {
			num1, num2 = num2, num1
		}

		div = num1 % num2
		num1 = num1 % num2

	}

	fmt.Println("MDC: ", num2)
}

func ex8() {
	// Faça um programa que inverta um número inteiro.
	var num string

	fmt.Scan(&num)

	numstring := strings.Split(num, "")

	var reverse string

	for i := len(numstring); i > 0; i-- {
		reverse += numstring[i-1]
	}

	fmt.Print(reverse)

}

func ex9() {
	// Escreva um programa que verifique se uma string é um palíndromo.
	var checkstring string

	fmt.Scan(&checkstring)

	checkstringSplit := strings.Split(checkstring, "")

	var reverse string

	for i := len(checkstringSplit); i > 0; i-- {
		reverse += checkstringSplit[i-1]
	}

	if reverse == checkstring {
		fmt.Println("É palíndromo")
	} else {
		fmt.Println("Não é palíndromo")
	}
}

func ex10() {
	// Implemente um programa que converta um número decimal para binário.
	var num int

	fmt.Scan(&num)

	div := 1
	leftover := []int{}

	for num != 0 {
		div = num % 2
		num /= 2
		leftover = append(leftover, div)
	}

	for i, j := 0, len(leftover)-1; i < j; i, j = i+1, j-1 {
		leftover[i], leftover[j] = leftover[j], leftover[i]
	}

	fmt.Println(leftover)
}

func ex11() {
	// Crie um programa que calcule a média de um conjunto de números.
	var input string

	println("Insira os números (use o espaço para separar)")
	fmt.Scan(&input)

	inputSlice := strings.Split(input, " ")
	fmt.Println(inputSlice)

	var average int

	for i := 0; i < len(inputSlice); i++ {
		num, _ := strconv.Atoi(inputSlice[i])
		average += num
	}

	average = average / len(inputSlice)

	println(average)
}

func ex12() bool {
	var num int

	fmt.Scan(&num)

	sum := 1

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			sum += i

			if i != num/i {
				println(num/i, i)
				sum += num / i
			}
		}
	}

	if sum == num {
		return true
	} else {
		return false
	}

}
