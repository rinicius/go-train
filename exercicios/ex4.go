package main

import (
	"bufio"
	"fmt"
	"os"
)

func ex4() {
	// abrir o arquivo e o armazenar em "file" através do "os"
	file, err := os.Open("exercicios/input.txt")

	// nil é a mesma coisa que "null"
	if err != nil {
		fmt.Println(err)
	}

	// bufio.NewScanner é a função que lê e armazena a leitura do arquivo (exemplo: as letras do arquivo)
	sc := bufio.NewScanner(file)

	// split separa a leitura. no caso, bufio.scanlines está separando em quebras de linha
	// é possível usar também outros métodos já prontos como o "scanbytes", "scanwords" e assim em diante
	sc.Split(bufio.ScanLines)

	var linhas []string

	// o for após o scan ter sido "splitado" vai rodar até não encontrar mais nenhuma linha (token) restante
	for sc.Scan() {
		linhas = append(linhas, sc.Text())
	}

	for i, l := range linhas {
		fmt.Printf("Linha %d: %s\n", i, l)
	}

	file.Close()
}
