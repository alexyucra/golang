# Requisicões Web

Neste exercício vamos iniciar o monitoramento de nossos sites e colocar o nosso programa para funcionar até que solicitemos o contrário.

1 - Vamos implementar a primeira funcionalidade do nosso script, que é a de monitoramento de um site. Crie a função `init_monitor()`, que fará uso do pacote `net/http` para fazer uma requisição para um site da web. Depois faça a chamada dela no primeiro case do switch.

<details>
    <summary>show me the code</summary>

```go
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main(){
    //restante da função 

    switch comando {
    case 1:
        init_monitor()
    //restante do switch
    }

}

// outras funções

func init_monitor() {
    fmt.Println("Monitorando...")
    site := "alexyucra.vcard.repl.co"
    result, _ := http.Get(site)
}
```
</details>

2 - Vamos fazer uma verificação através do status code da requisição para averiguar se o site está online ou sofreu alguma queda. Teste para ver se o status code é `200` e exiba mensagens de acordo. 
Utilize o operador **`_`** para não ter lidar com o segundo parâmetro de erro agora.

```go
func init_monitor() {
    fmt.Println("Monitorando...")
    site := "alexyucra.vcard.repl.co"
    result, _ := http.Get(site)

    if result.StatusCode == 200 {
		fmt.Println("Site: ", site, "status: Run")
	} else {
		fmt.Println("Site: ", site, "status: ???")
	}
}
```

3 - Agora com a função de monitoramento pronta, vamos colocar um loop infinito na função main, para que o usuário possa acessar o menu repetidas vezes até que ele escolha sair. Para isto, utilize a instrução `for` pura:

```go
func main() {

    view_info()

    for {
        view_menu()
        comando := read_comando()

        switch comando {
        case 1:
            init_monitor()
        // restante do switch
        }

    }
}
```

Pronto, agora o nosso usuário já consegue monitorar pelo menos um site até quando ele desejar sair do programa!

___

<details>
    <summary>$ go run hello.go</summary>

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func main() {

	view_info()

	// loop infinito sem parametros
	for {
		view_menu()
		comando := read_comando()

		switch comando {
		case 1:
			fmt.Println("[1] Monitorando ...")
			init_monitor()
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

func init_monitor() {
	fmt.Println("Monitorando...")
	site := "https://alexyucra.vcard.repl.co"
	result, _ := http.Get(site)

	if result.StatusCode == 200 {
		fmt.Println("Site: ", site, "status: Run")
	} else {
		fmt.Println("Site: ", site, "status: ???")
	}
}
```

</details>
