# Leitura de arquivos

Neste exercício vamos fazer com que os sites monitorados venham de um arquivo `.txt` para que fique fácil adicionar ou remover um site do monitoramento, além disto vamos colocar algumas detecções de erro em nossas funções.

1 - Primeiro, vamos colocar nosso aprendizado sobre erros deste capítulo para capturar os erros, o primeiro deles de uma função que já utilizamos, a `http.Get()` da funçao `ping_url()`:

```go
// restante do código omitido ...
func ping_url(site string) {

    //Capturando o segundo parâmetro
    resp, err := http.Get(site)
    //Verificando se houve algum erro
    if err != nil {
        fmt.Println("[error] (ping_url):", err)
    }

    if result.StatusCode == 200 {
        fmt.Println("Site: ", url, "status: Run")
    } else {
        fmt.Println("Site: ", url, "status: erro: ???")
    }
}
```

2 - Agora vamos começar a ler os nossos sites de um arquivo .txt. Crie um arquivo de texto chamado `sites.txt` e coloque os sites que você quer monitorar lá, um em cada linha:

```go
// neste exemplo voce pode criar o arquivo sites.txt manualmente
func read_file() []string {

    var sites []string

    return sites;
}
```

4 - Agora vamos abrir o arquivo de sites.txt e detectar caso algum erro aconteça na abertura. Vamos utilizar o pacote os para abrir com a função `Open:

```go
func read_file() []string {

    var sites []string
    file_names := "sites.txt"

    fmt.Println("")
    fmt.Println("------metodo 1-----")
    // metodo 001: os.Open() result: &{0xc000214060}
    file1, err := os.Open(file_names)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }
    // Não esqueça de fechar o arquivo no final com a função 
    file1.Close() 
    return sites
}
```

5 - Agora vamos criar um leitor com o pacote `bufio`, para que facilite o processo de percorrer o arquivo `sites.txt`:

```go
func read_file() []string {

    var sites []string
    file_names := "sites.txt"

    fmt.Println("")
    fmt.Println("------metodo 1-----")
    // metodo 001: os.Open() result: &{0xc000214060}
    file1, err := os.Open(file_names)

    if err != nil {
        fmt.Println("[error 1] (read_file):", err)
        log.Fatal(err)
        // si arquivo ñao é encontrado criamos
        f, err := os.Create(file_names)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()
    }
    fmt.Println(file1)

    fmt.Println("")
    fmt.Println("------metodo 2-----")
    // metodo 002: ioutil.ReadFile() result: array de bits use string() para converter
    file2, err := ioutil.ReadFile(file_names)
    if err != nil {
        fmt.Println("[error 2] (read_file):", err)
        log.Fatal(err)
    }
    fmt.Println(string(file2))

    // Não esqueça de fechar o arquivo no final com a função 
    file1.Close() 
    return sites
}
```

6 - Agora vamos utilizar um for e ler linha a linha até que o leitor encontre o EOF, o que nos dará um erro e utilizaremos a instrução break para sair do loop, indicando que chegamos ao final. Para ler cada linha utilizaremos a função ReadString do leitor, lendo até o caractere `\n` que indica o final da linha:

7 - Por último vamos dar um trim em cada linha lida para remover caracteres especiais, como \n , espaços e tabs. E claro, após isto vamos adicionar a linha ao slice de sites:

```go
func read_file() []string {

    var sites []string
    file_names := "sites.txt"

    fmt.Println("")
    fmt.Println("------metodo 1-----")
    // metodo 001: os.Open() result: &{0xc000214060}
    file1, err := os.Open(file_names)
    if err != nil {
        fmt.Println("[error 1] (read_file):", err)
        log.Fatal(err)
        // si arquivo ñao é encontrado criamos
        f, err := os.Create(file_names)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()

    }
    fmt.Println(file1)

    fmt.Println("")
    fmt.Println("------metodo 2-----")
    // metodo 002: ioutil.ReadFile() result: array de bits use string() para converter
    file2, err := ioutil.ReadFile(file_names)
    if err != nil {
        fmt.Println("[error 2] (read_file):", err)
        log.Fatal(err)
    }
    fmt.Println(string(file2))

    fmt.Println("")
    fmt.Println("------metodo 3-----")
    // metodo 003: leitura linha-linha result:
    file3, err := os.Open(file_names)
    if err != nil {
        fmt.Println("[error 3.1] (read_file):", err)
    }

    leitor := bufio.NewReader(file3)
    for {
        linha, err := leitor.ReadString('\n') // ler ate o delimitador da linha com '' simples
        linha = strings.TrimSpace(linha)
        // fmt.Println(linha)
        sites = append(sites, linha)
        if err == io.EOF {
            break // quando encontre o final sair do loop
        }
        // if err != nil {
        // fmt.Println("[error 3.2] (read_file):", err)
        // }
        // fmt.Println(linha)
    }
    file3.Close()

    return sites
}
```

8 - Agora vamos chamar a nossa função récem criada dentro função `init_monitor()`, para que ao invés de obtermos os sites de um slice fixo, obteremos do arquivo `sites.txt`:

```go
func init_monitor() {
    fmt.Println("Monitorando...")

    sites := read_file()

    // restante da função
}
```

___

<details>
    <summary>show me the code</summary>

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
    "reflect"
    "strings"
    "time"
)

const num_monitoramentos = 1
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

    // // usando slice por array
    // sites := []string{"https://alexyucra.vcard.repl.co", "https://m.vcard.repl.co", "https://google.com"}
    // fmt.Println("Usando slices", sites)
    // fmt.Println(reflect.TypeOf(sites))

    sites := read_file()

    // teste en loop
    for i := 0; i < num_monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Monitorando site: [", i, "] -> ", site)
            ping_url(site)
        }
        // adicionando delay cada 10 seg
        time.Sleep(delay * time.Second)
    }
    fmt.Println("")
}

func ping_url(url string) {
    result, err := http.Get(url)
    if err != nil {
        fmt.Println("[error] (ping_url):", err)
    }

    if result.StatusCode == 200 {
        fmt.Println("Site: ", url, "status: Run")
    } else {
        fmt.Println("Site: ", url, "status: erro: ???")
    }
    }

func read_file() []string {

    var sites []string
    file_names := "sites.txt"

    fmt.Println("")
    fmt.Println("------metodo 1-----")
    // metodo 001: os.Open() result: &{0xc000214060}
    file1, err := os.Open(file_names)
    if err != nil {
        fmt.Println("[error 1] (read_file):", err)
        log.Fatal(err)
        // si arquivo ñao é encontrado criamos
        f, err := os.Create(file_names)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()

    }
    fmt.Println(file1)

    fmt.Println("")
    fmt.Println("------metodo 2-----")
    // metodo 002: ioutil.ReadFile() result: array de bits use string() para converter
    file2, err := ioutil.ReadFile(file_names)
    if err != nil {
        fmt.Println("[error 2] (read_file):", err)
        log.Fatal(err)
    }
    fmt.Println(string(file2))

    fmt.Println("")
    fmt.Println("------metodo 3-----")
    // metodo 003: leitura linha-linha result:
    file3, err := os.Open(file_names)
    if err != nil {
        fmt.Println("[error 3.1] (read_file):", err)
    }

    leitor := bufio.NewReader(file3)
    for {
        linha, err := leitor.ReadString('\n') // ler ate o delimitador da linha com '' simples
        linha = strings.TrimSpace(linha)
        // fmt.Println(linha)
        sites = append(sites, linha)
        if err == io.EOF {
            break // quando encontre o final sair do loop
        }
        // if err != nil {
        // 	fmt.Println("[error 3.2] (read_file):", err)
        // }
        // fmt.Println(linha)
    }
    file3.Close()

    return sites
}
```

</details>

