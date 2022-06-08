package main

import "fmt"

func main() {
	// Constantes:
	const HELLO string = "Hello World"
	fmt.Println(HELLO)

	// Definindo uma variável especificando tipo:
	var mensagem1 string = "Mensagem 1"
	fmt.Println(mensagem1)

	// Definindo uma varíavel inferindo o tipo:
	mensagem2 := "Mensagem 2"
	fmt.Println(mensagem2)

	// Ponteiros:

	var (
		x  int  // variável comum do tipo int
		xp *int // variável que aponta para outra variável do tipo int
	)

	x = 19
	xp = &x // & é usado para buscar o endereço de memória da variável

	fmt.Println("Ponteiro: ", xp)  // imprime o endenreco apontado
	fmt.Println("Ponteiro: ", *xp) // imprime o valor usando o operador *

	x = 10

	fmt.Println("Ponteiro: ", xp)
	fmt.Println("Ponteiro: ", *xp)

	*xp = 55 // usando o operador * podemos alterar o valor da variável apontada

	fmt.Println("Variavel apontada: ", x)

	// Conversão de tipos

	// Estruturas condicionais:
	if true {
		fmt.Println("Entrou no if")
	} else if true {
		fmt.Println("Entrou no else if")
	} else {
		fmt.Println("Entrou no else")
	}

	if simplifica := true; simplifica {
		fmt.Println("Entrou no if simplificado")
	}

	switch 1 {
	case 1:
		fmt.Println("Entrou no 1")
		break
	case 2:
		fmt.Println("Entrou no 2")
		break
	case 3:
		fmt.Println("Entrou no 3")
		break
	default:
		fmt.Println("Entrou no default")
	}

	// Estrutura de repetição e iteração:

	// Loop infinito até que encontre um break
	sentinela := 0
	for {
		if sentinela == 10 {
			fmt.Println("break foi chamado no loop infinito")
			break
		}
		sentinela++
	}

	// Loop While
	sentinela = 0
	for sentinela < 10 {
		fmt.Println("While: ", sentinela)
		sentinela++
	}

	// Loop For & continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("For: ", i)
	}

	// Loop For com "range" (foreach)
	for i, value := range []string{"Hello", "World", "of", "Golang"} {
		fmt.Println("For Range: Index ", i, " Value ", value)
	}

	// Arrays e Slices:
	var arraySimples [5]string = [5]string{"a", "b", "c", "d", "e"}
	for _, value := range arraySimples {
		fmt.Println("Array simples: ", value)
	}

	var sliceSimples []string = []string{}
	sliceSimples = append(sliceSimples, "Go")
	sliceSimples = append(sliceSimples, "Java")
	sliceSimples = append(sliceSimples, "Python")
	for _, value := range sliceSimples {
		fmt.Println("Slice Simples: ", value)
	}

	// Arrays são estáticos, Slices trabalham com referências (não armazenam valores diretamente).
	// Ao usar append em um array é provocado um erro por ter tamanho fixo, diferente do slice
	// que torna a utilização de "listas" mais flexíveis.
	// arraySimples = append(arraySimples, "valor")

	// Criando um slice a partir de um array
	slicePorArray := arraySimples[0:3]
	for _, value := range slicePorArray {
		fmt.Println("Slice por Array: ", value)
	}

	// Como slices alocam referências, podemos alterar um valor de um array
	// quando o slice nasce a partir de um array
	slicePorArray[2] = "Valor foi alterardo no arraySimples"
	for _, value := range arraySimples {
		fmt.Println("Array simples: ", value)
	}

	// Podemos ter referencias de arrays difirentes em um único slice
	outroArray := [5]string{"1", "2", "3", "4", "5"}
	slicePorArray = append(slicePorArray, outroArray[:]...) // pega slice de todas as posicoes e faz um unpacking no append
	for _, value := range slicePorArray {
		fmt.Println("Slice por Array: ", value)
	}

	// Atenção: na prática é raramente usado arrays, devido a simplicidade dos slices
	// eles dominam boa parte do código, e em alguns casos nunca verá um array em um projeto
}
