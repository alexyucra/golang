# Contro de fluxo

Neste capitulo vamos determinar para qual fluxo do nosso programa o usuário deve seguir a partir do comando que ele escolheu, e além disso vamos começar a organizar nosso código em pequenas funções.

1 - O primeiro passo vai ser criar um `switch` com todas as opções que oferecemos ao usuário até agora, e um caso de `default` para quando o usuário colocar algum comando desconhecido:
<details>
	<summary>source code</summary>


```go
//hello.go

package main

import "fmt"

func main() {
    // ... restante da função main
    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor da variável comando é:", comando)

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
    default:
        fmt.Println("Não conheço este comando")
    }
}
```
</details>

	Aqui colocamos uma opção no `switch` para cada item do menu que tinhamos no menu anteriormente

2 - Agora vamos começar a organizar nosso código em pequenas funções, cada uma com sua responsabilidade. A primeira a ser criada será a função que vai exibir a introdução de boas vindas para o usuário. Extrai este código da função main e coloque na função `view_info()`:

<details>
	<summary>source code</summary>

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {

	view_info()

	comando := read_comando()

	switch comando {
	case 1:
		fmt.Println("[1] Monitorando ...")
	case 2:
		fmt.Println("[2] Exibindo Logs")
	case 0:
		fmt.Println("[0] Volte sempre !!!")
		os.Exit(0)
	default:
		fmt.Println("[?] Comando desconhecido ?")
		os.Exit(-1) // sair com err: exit status 255
	}
}

func view_info() {
	nome := "Alex"
	cmd := exec.Command("bash", "-c", "go version")
	version, err := cmd.Output()
	if err != nil {
		fmt.Println((err.Error()))
		return
	}

	fmt.Println("Hola ", nome, "! ", nome, "es variable de tipo ", reflect.TypeOf(nome))
	fmt.Println("Voce esta usando a GO ", string(version))
}
```
</details>

> Não esqueça de chamar a função `view_info()` na função main.

3 - Vamos extrair também o código que exibe o menu de opções do usuário para uma função externa, chamda de `view_menu()`:

<details>
	<summary>source code</summary>

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {

	view_info()

	view_menu()

	// ...
}

func view_info() {
	// ...
}

func view_menu() {
	// menu
	fmt.Println("+", strings.Repeat("-", 25), "+")
	fmt.Println("| 1-Iniciar Monitoramento   |")
	fmt.Println("| 2-Exibir Logs             |")
	fmt.Println("| 0-Sair do Programa        |")
	fmt.Println("+", strings.Repeat("-", 25), "+")
}
```
</details>

4 - Por último, vamos criar uma função responsável por ler os comandos do usuário e exportar esta função da `main` também. A função `read_comando()` deve retornar um `**int**` com o comando lido pelo usuário:

<details>
	<summary>source code</summary>

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {

	view_info()

	view_menu()

	comando := read_comando()

	// ...
}

func view_info() {
	// ...
}

func view_menu() {
	// ...
}

func read_comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Voce escolheu a opcao: ", comando, "-> com endereco: ", &comando)
	return comando
}
```

</details>

5 - Vamos por último adicionar nos itens finais do switch um ponto de saída para o nosso script, retornando para o sistema operacional se tudo correu bem ou não. Importe o pacote `os` do Go e utilize a função Exit() para informar o status de saída do programa:

<details>
	<summary>source code</summary>

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {

	view_info()
	view_menu()
	comando := read_comando()

	switch comando {
	case 1:
		fmt.Println("[1] Monitorando ...")
	case 2:
		fmt.Println("[2] Exibindo Logs")
	case 0:
		fmt.Println("[0] Volte sempre !!!")
		os.Exit(0)
	default:
		fmt.Println("[?] Comando desconhecido ?")
		os.Exit(-1) // sair com err: exit status 255
	}
}
//... restante das funcoes
```
</details>

Pronto, nosso código está mais organizado e agora só falta implementar as funções de cada um dos casos do nosso switch.

___ 

## Exercicio

Implemente na linguagem Go um programa que solicita a escolha do cardápio.
<details>
  
<summary>cardapio</summary>
    
```go
// cardapio
package main

import "fmt"

func main() {
    var prato string
    fmt.Println("Digite seu prato preferido...")
    fmt.Println("P - Pizza")
    fmt.Println("H - Hambúrguer")
    fmt.Println("B - Bife com fritas")
    fmt.Println("S - Salada Caesar")
    fmt.Println("F - Salada de Frutas")
    fmt.Println("E - Estrogonofe")
    fmt.Println("O - Outros")
    fmt.Scan(&prato)

    switch prato {
    case "B":
        fmt.Println("Com batatas Palito ou Noisete?")
    case "H":
        fmt.Println("Com Queijo ou com Ovo?")
    case "P":
        fmt.Println("Calabresa ou Napolitana?")
    case "S":
        fmt.Println("Alface ou Rúcula?")
    case "F":
        fmt.Println("Kiwi ou Frango?")
    case "E":
        fmt.Println("Carne ou Frango?")
    case "O":
        fmt.Println("Não gostou de nosso cardápio?")
    default:
        fmt.Println("Não entendi seu paladar...")
    }
}
```

</details>

<details>
  
<summary>Código fonte com fluxo </summary>
    
```go
package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
)

func main() {
	nome := "Alex"
	cmd := exec.Command("bash", "-c", "go version")
	version, err := cmd.Output()
	if err != nil {
		fmt.Println((err.Error()))
		return
	}

	fmt.Println("Hola ", nome, "! ", nome, "es variable de tipo ", reflect.TypeOf(nome))
	fmt.Println("Voce esta usando a GO ", string(version))

	fmt.Println("+", strings.Repeat("-", 25), "+")
	// menu
	fmt.Println("| 1-Iniciar Monitoramento   |")
	fmt.Println("| 2-Exibir Logs             |")
	fmt.Println("| 0-Sair do Programa        |")

	fmt.Println("+", strings.Repeat("-", 25), "+")

	// aceitar variables do input
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Voce escolheu a opcao: ", comando, "-> com endereco: ", &comando)

	// com if
	if comando == 1 {
		fmt.Println("Monitorando ...")
	} else if comando == 2 {
		fmt.Println("verificando Logs ...")
	} else if comando == 0 {
		fmt.Println("Volte sempre !")
	} else {
		fmt.Println("Comando desconhecido ?, tente de novo")
	}

	// com sitch
	switch comando {
	case 1:
		fmt.Println("[1]")
	case 2:
		fmt.Println("[2]")
	case 0:
		fmt.Println("[0]")
	default:
		fmt.Println("[?]")
	}

}

```

</details>

<details>
	<summary>Código fonte com funcões</summary>

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {

	view_info()
	view_menu()
	comando := read_comando()

	switch comando {
	case 1:
		fmt.Println("[1] Monitorando ...")
	case 2:
		fmt.Println("[2] Exibindo Logs")
	case 0:
		fmt.Println("[0] Volte sempre !!!")
		os.Exit(0)
	default:
		fmt.Println("[?] Comando desconhecido ?")
		os.Exit(-1) // sair com err: exit status 255
	}

}

func view_info() {
	nome := "Alex"
	cmd := exec.Command("bash", "-c", "go version")
	version, err := cmd.Output()
	if err != nil {
		fmt.Println((err.Error()))
		return
	}

	fmt.Println("Hola ", nome, "! ", nome, "es variable de tipo ", reflect.TypeOf(nome))
	fmt.Println("Voce esta usando a GO ", string(version))
}

func view_menu() {
	// menu
	fmt.Println("+", strings.Repeat("-", 25), "+")
	fmt.Println("| 1-Iniciar Monitoramento   |")
	fmt.Println("| 2-Exibir Logs             |")
	fmt.Println("| 0-Sair do Programa        |")
	fmt.Println("+", strings.Repeat("-", 25), "+")
}

func read_comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Voce escolheu a opcao: ", comando, "-> com endereco: ", &comando)
	return comando
}
```

</details>