# Variáveis e Struturas

Neste capitulo criamos nossa primeira struct chamada ContaCorrente;

```go
type ContaCorrente struct {
    // variaveis
}
```

Em seguida criamos um tipo com base na struct Conta corrente sem atribuir valor para os campos e descobrimos o Zero-initialization ou Inicialização zero;

```go
type ContaCorrente struct {
    titular    string
    numAgencia nil
    numConta   int
    saldo      float64
}
```

Go garante inicializar todas as variáveis, conforme a imagem abaixo:
|| Zero-Initialization |
| -- | -- |
| bool| false |
| int |  0 |
| float | "" |
| string | "" |
| struct | {} |
| Nulo | nil |

Para finalizar, vimos duas formas utilizar a struct criada: informando ou não os nomes dos campos.

```go
func main() {
    contaAlex := ContaCorrente{titular: "Alex"}
    fmt.Println(contaAlex)

    contaGonzalo := ContaCorrente{"Gonzalo", 543, 123457, 0}

    fmt.Println(contaGonzalo)
}
```
> por mais que o Go garanta a inicialização zero de diferentes tipos, devemos ficar atentos com os tipos que estamos trabalhando.

compilando o script

```shell
go run main.go
# {Alex 0 0 0}
# {Gonzalo 543 123457 0}
```

- show me the code

```go
package main

import "fmt"

// criamos uma structura, que contem nossa variaveis e tipos
type ContaCorrente struct {
    titular    string
    numAgencia nil
    numConta   int
    saldo      float64
}

func main() {
    contaAlex := ContaCorrente{titular: "Alex"}
    fmt.Println(contaAlex)

    contaGonzalo := ContaCorrente{"Gonzalo", 543, 123457, 0}

    fmt.Println(contaGonzalo)
}
```