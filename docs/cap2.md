# Trabalhando com variáveis



___

<details>
  <summary>source code</summary>
  


```shell
package main

import (
	"fmt"
	"os/exec"
	"reflect"
)

func main() {
	// var nome string = "Alex"
	// var idade int = 29
	// var versao float32 = 1.1

	// or

	nome := "Alex"
	cmd := exec.Command("bash", "-c", "go version")
	version, err := cmd.Output()
	if err != nil {
		fmt.Println((err.Error()))
		return
	}

	fmt.Println("Hola ", nome, " !")
	fmt.Println("O tipo de variable 'nome'", reflect.TypeOf(nome))
	fmt.Println("Voce esta usando a GO ", string(version))

	fmt.Println("1-Iniciar Monitoramento")
	fmt.Println("2-Exibir Logs")
	fmt.Println("0-Sair do Programa")

	// aceitar variables do input
	var comando int
	// fmt.Scanf("%d", &comando)
	// or
	fmt.Scan(&comando)
	fmt.Println("Voce escolheu a opcao: ", comando, "-> com endereco: ", &comando)

}

// to build run: go build hello.go
// compile and run : go run hello.go
```
</details>
