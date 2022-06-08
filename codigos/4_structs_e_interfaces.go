package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

// 1 - Structs & Métodos:

type Pessoa struct {
	Nome  string
	Idade int
}

// O (p *Pessoa) é um receiver, o golang passa implícitamente o ponteiro para
// essa variável receiver dando acesso à própria "instância" da struct, é semelhante ao this em Java
func (p *Pessoa) Diz(mensagem string) {
	output := fmt.Sprintf("%s (%d anos): %s", p.Nome, p.Idade, mensagem)
	fmt.Println(output)
}

// 2 - Genéricos

type Echo[T int | string | bool] struct {
	Output T
}

func (e Echo[T]) Execute() {
	fmt.Println(e.Output)
}

// 3 - Interfaces:

// Interfaces são implementadas implícitamente

type Command interface {
	Execute()
}

type HelloCommand struct {
	Message string
}

func (h *HelloCommand) Execute() {
	fmt.Println(h.Message)
}

func Execute(command Command) {
	command.Execute()
}

// 5 - Pseudo-Herança

type Funcionario struct {
	Pessoa
	Salario float64
}

// 6 - Struct Tags

type Response struct {
	Status  int    `json:"status" xml:"response_status"`
	Message string `json:"message" xml:"response_message"`
}

func main() {
	// 1 - Structs & Métodos:
	pessoa := Pessoa{
		Nome:  "Fulano de Tal",
		Idade: 25,
	}
	pessoa.Diz("Olá mundo!")

	// 2 - Genéricos
	echo := Echo[string]{
		Output: "Hello World com Echo",
	}
	echo.Execute()

	// 3 - Interfaces:
	command := HelloCommand{
		Message: "Hello World com Comando",
	}
	Execute(&command)
	Execute(&echo)

	// 5 - Pseudo-Herança
	funcionario := Funcionario{
		Pessoa: Pessoa{
			Nome:  "Pedrinho",
			Idade: 30,
		},
		Salario: 980.5,
	}
	funcionario.Diz(fmt.Sprintf("Meu salário é: %.2f", funcionario.Salario))
	// voce pode acessar os mebros usando operador de membro . (ponto)
	// funcionario.Nome
	// ou acessando pela variável gerada internamente, funcionar.Pessoa.Nome

	// 6 - Struct Tags
	resonse := Response{
		Status:  200,
		Message: "Ok! Sucesso.",
	}
	if jsonBytes, err := json.Marshal(resonse); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonBytes))
	}
	if xmlBytes, err := xml.MarshalIndent(resonse, "", "  "); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(xmlBytes))
	}
}
