# Introducaõ a GO

Neste primeiro capítulo vamos nos focar em criar o nosso Hello World inicial com a linguagem.

O primeiro passo é criar a pasta `hello`, que irá conter o nosso primeiro script em Go. Crie a `pasta/hello` dentro da pasta src do seu Go Workspace, que criamos anteriormente no exercício de Preparando o Ambiente.

```shell
pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
        └── hello
```

Todos os projetos que você for desenvolver na linguagem Go ficarão em suas próprias pastas no diretório src.

Agora vamos criar o primeiro arquivo , o hello.go dentro pasta /hello:

```shell
pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
        └── hello
           └── hello.go
```

E neste arquivo vamos criar o nosso primeiro programa.

Como todo projeto em Go, precisamos definir qual será o pacote inicial com a instrução:

```go
//hello.go
package main
```

E também precisamos definir a função principal do programa,a nossa função main:

```go
//hello.go
package main

func main(){

}
```

Como desejamos exibir uma mensagem, precisamos importar o pacote fmt, que contêm as funções de formatação da linguagem inclusive a função que imprime uma mensagem, a função fmt.Println():

```go
//hello.go
package main

import "fmt"

func main(){

}
```

Para não dar azar e começar nossa jornada no mundo do Go com o pé direito, vamos começar com o Hello World clássico de todas as linguagens:

```go
//hello.go
package main

import "fmt"

func main(){
    fmt.Println("Hello World com Go!")
}
```

Para executar nosso código, basta utilizarmos o comando go run hello.go no terminal dentro da pasta que contêm nosso arquivo com o código fonte do programa que o executável será automaticamente criado e executado :

```shell
// Terminal
go run hello.go
Hello World com Go!
```

Nosso primeiro programa está completo!
