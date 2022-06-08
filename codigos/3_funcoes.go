package main

import "fmt"

// Função simples que executa uma comando de impressão
func helloWorld() {
	fmt.Println("Funcao Simples: Hello World")
}

// Função simples que retorna um valor
func getHelloWorld() string {
	return "Funcao Simples com Retorno: Returning Hello World"
}

// Função com multiplos retornos
func swapIntegers(x int, y int) (int, int) {
	return y, x
}

// Função com retornos nomeados
func swapStrings(xIn string, yIn string) (xOut string, yOut string) {
	xOut = yIn
	yOut = xIn
	return
}

// Função de auto nível(retorna uma outra Função)
func autoNivel() func() {
	return func() {
		fmt.Println("Funcao de Auto Nivel: Hello Em Auto Nível")
	}
}

// Função Variádica
func echo(args ...string) {
	// args é um slice com todos os argumentos passados por ele
	for _, value := range args {
		fmt.Println("Funcao Variadica: ", value)
	}
}

// Função com generics
func swapGeneric[T any](x T, y T) (T, T) {
	return y, x
}

func main() {
	helloWorld()

	mensagem := getHelloWorld()
	fmt.Println(mensagem)

	x, y := swapIntegers(10, 20)
	fmt.Println("Funcao com multiplos retornos: ", x, y)

	s1, s2 := swapStrings("Hello", "World")
	fmt.Println("Funcao com retornos nomeados: ", s1, s2)

	var f func() = autoNivel()
	f()

	inputs := []string{"Hello", "Variadic", "Functions"}
	echo(inputs...) // os tres pontos vazem unpacking do slice para o echo

	fmt.Println(swapGeneric("A", "B"))
	fmt.Println(swapGeneric(1, 2))
}
