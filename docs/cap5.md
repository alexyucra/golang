# Coleções GO

Neste exercício vamos melhorar o monitoramento de nosso programa, para que ele monitore mais de um site ao mesmo tempo.

1 - Nosso primeiro passo é substituir a nossa string `site` pelo slice `sites`, que conterá os endereços dos vários sites a serem testados. Vamos testar pelo menos 3 sites:
- https://alexyucra.vcard.repl.co
- https://m.vcard.repl.co
- https://google.com
  
<details>
    <summary>show me the code</summary>

```go
//hello.go

//restante do arquivo

func init_monitor() {
    fmt.Println("Monitorando...")

    sites := []string{
        "https://alexyucra.vcard.repl.co", 
        "https://m.vcard.repl.co", 
        "https://google.com"
    }
    //restante da função
}
```
</details>

2 - Como queremos que cada um destes sites seja testado uma vez, vamos um for para percorrer o slice inteiro, utilize o range para facilitar:

<details>
    <summary>show me the code</summary>

```go
//hello.go

//restante do arquivo

func init_monitor() {
    fmt.Println("Monitorando...")

    sites := []string{
        "https://alexyucra.vcard.repl.co", 
        "https://m.vcard.repl.co", 
        "https://google.com"
    }

    for i, site := range sites {
        // restante 
    }
}
```
</details>

3 - Vamos agora extrair o nosso código que testa um site para uma função externa, para facilitar na hora que estamos iterando sobre o slice, aonde cada site deve executar uma vez a função. Crie a função `ping_ur()` que deve receber uma string com o site a ser testado e mova o código go http.Get para lá:

<details>
    <summary>show me the code</summary>

```go
//hello.go 
//restante do arquivo

func ping_url(url string) {
	result, _ := http.Get(url)

	if result.StatusCode == 200 {
		fmt.Println("Site: ", url, "status: Run")
	} else {
		fmt.Println("Site: ", url, "status: [???]")
	}
}
```
</details>

4 - Vamos agora chamar a funcao `ping_url()` dentro do nosso for que está varrendo o slice:

<details>
    <summary>show me the code</summary>

```go
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://alexyucra.vcard.repl.co", "https://m.vcard.repl.co", "https://google.com"}

    for i, site := range sites {
        fmt.Println("Monitorando site: [", i, "] -> ", site)
        ping_url(site)
    }

    fmt.Println("")
}
```
</details>

> Adicione o `fmt.Println("")` vazio ao final, e o que está dentro do for também, para que o nosso script vá exibindo mensagens para o nosso usuário enquanto ele é executado.

Execute o nosso script e veja que agora conseguimos testar diversos sites, e ao final do teste o nosso Menu surge para escolhermos a opção de monitorar novamente.

___

## Delay e Aumentando o número de testes

Agora vamos fazer com que a cada vez que o usuário selecione a opção de monitoramento, cada site seja testado mais de uma vez, de acordo com o que o usuário setar nas constantes.

1 - Primeiro crie a constante `num_monitoramentos` que indicará quantas vezes o site será testado. Coloque o número que desejar, mas lembre-se que quanto mais você testar, mais o script vai demorar:

```go
//hello.go
package main

import (
    "fmt"
    "net/http"
    "os"    
)

const num_monitoramentos = 3

// restante do arquivo
```

2 - Agora vamos colocar um for dentro da função `init_monitor()` para que o slice seja percorrido o número de vezes que configuramos na constante:

```go
//hello.go

//restante do arquivo
func init_monitor() {
	fmt.Println("Monitorando...")

	// usando slice por array
	sites := []string{"https://alexyucra.vcard.repl.co", "https://m.vcard.repl.co", "https://google.com"}
	fmt.Println("Usando slices", sites)
	fmt.Println(reflect.TypeOf(sites))

	// 2 forma de percorrer slice com range
	for i := 0; i < num_monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Monitorando site: [", i, "] -> ", site)
			ping_url(site)
		}
	}
	fmt.Println("")
}
```

3 - Para que haja uma pausa entre cada um dos testes que faremos, vamos adicionar um pequeno `delay` entre cada vez cada iteração de monitoramento. Utilize a função `Sleep` do pacote `time` para dar uma pausa entre monitoramentos. Para que o tamanho desta pausa fique configurável, vamos criar uma constante que vai dizer o tamanho do delay:

```go
//hello.go
package main

import (
    "fmt"
    "net/http"
    "os"    
)

const monitoramentos = 3
const delay = 5

// restante do arquivo
```

4 - E a partir do valor da constante, vamos adicionar um `time.Sleep` na função `init_monitor()`. Vamos multiplicar a nossa contante (`delay`) pela constante que representa o número de segundos no Go(`time.Seconds`):

```go
//hello.go

//restante do arquivo
func init_monitor() {
	fmt.Println("Monitorando...")

	// usando slice por array
	sites := []string{"https://alexyucra.vcard.repl.co", "https://m.vcard.repl.co", "https://google.com"}
	fmt.Println("Usando slices", sites)
	fmt.Println(reflect.TypeOf(sites))

	// 2 forma de percorrer slice com range
	for i := 0; i < num_monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Monitorando site: [", i, "] -> ", site)
			ping_url(site)
		}
		// adição AQUI!
        time.Sleep(delay * time.Second)
        fmt.Println("")
	}
	fmt.Println("")
}
```

> Adicione o outro `fmt.Println("")` para que a exibição fique mais organizada para o seu usuário.

Agora o seu script deve estar testando os seus sites de acordo com o número de vezes que você setou na constante monitoramentos, e entre cada testada ele deve dar uma pausa de acordo com o número de segundos que você configurou na constante delay.


___

<details>
    <summary>$ go run hello.go </summary>

```go
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"time"
)

const num_monitoramentos = 3
const delay = 5

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

	// usando arrays
	var n_sites [4]string
	n_sites[0] = "https://alexyucra.vcard.repl.co"
	n_sites[1] = "https://m.vcard.repl.co"
	n_sites[2] = "https://google.com"
	fmt.Println("usando arrays", n_sites)
	fmt.Println(reflect.TypeOf(n_sites))

	// usando slice por array
	sites := []string{"https://alexyucra.vcard.repl.co", "https://m.vcard.repl.co", "https://google.com"}
	fmt.Println("Usando slices", sites)
	fmt.Println(reflect.TypeOf(sites))

	// percorrendo o slice
	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
	}
	fmt.Println("")

	// 2 forma de percorrer slice com range
	for i, site := range sites {
		fmt.Println("Monitorando site: [", i, "] -> ", site)
		ping_url(site)
	}
	fmt.Println("")

	// teste en loop
	//
	for i := 0; i < num_monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Monitorando site: [", i, "] -> ", site)
			ping_url(site)
		}
		// adicionando delay cada 10 seg
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")

	// site := "https://alexyucra.vcard.repl.co"
	// result, _ := http.Get(site)

	// if result.StatusCode == 200 {
	// 	fmt.Println("Site: ", site, "status: Run")
	// } else {
	// 	fmt.Println("Site: ", site, "status: ???")
	// }
}

func ping_url(url string) {
	result, _ := http.Get(url)

	if result.StatusCode == 200 {
		fmt.Println("Site: ", url, "status: Run")
	} else {
		fmt.Println("Site: ", url, "status: erro: ???")
	}
}

```
</details>
