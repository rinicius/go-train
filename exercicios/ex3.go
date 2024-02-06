package main

func ex3(frase string) (resultado string) {
	for _, l := range frase {
		resultado = string(l) + resultado
	}

	return resultado
}
