# Trabalhando com variáveis

Neste capítulo vamos criar a mensagem de introdução do script e o menu que será exibido para o usuário, além de capturar a escolha que ele fez.

1 - Para começar a trabalhar com variáveis, vamos declarar o nome do usuário e sua versão em duas variáveis, utilizando o operador declaração de variáveis curto:

```go
//hello.go
package main

import "fmt"

func main() {
    nome := "Alex"
    versao := 1.1
}
```

2 - Vamos exibir as mensagens de boas vindas com estas duas variáveis:

```go
//hello.go
package main

import "fmt"

func main() {
    nome := "Alex"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)
}
```

3 - Agora vamos exibir o Menu para o usuário escolher qual opção ele deseja. Serão 3 `Println` para exibir o Menu, um para cada opção:

```go
//hello.go
package main

import "fmt"

func main() {
    nome := "Alex"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}
```

4 - Com o menu sendo exibido corretamente, precisamos capturar a escolha do usuário do nosso programa. Vamos colocar a opção escolhida na variável int comando e capturar o seu valor com a função `fmt.Scan()`:

```go
//hello.go
package main

import "fmt"

func main() {
    nome := "Alex"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
}
```

> Atenção para não esquecer de passar o endereço (&) da variável comando para a função Scan

5- Por último imprima o valor da variável comando para garantir que estamos capturando com sucesso:

```go
//hello.go
package main

import "fmt"

func main() {
    nome := "Alex"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor da variável comando é:", comando)
}
```

Conseguimos criar nossa mensagem de introdução e o menu de opções do usuário, além de exibir qual opção ele escolheu. Já estamos avançando em nosso programa !

___

<details>
  <summary>source code</summary>
  
```go
package main

import (
	"fmt"
	"os/exec"
	"reflect"
)

func main() {
	// var nome string = "Alex"
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
