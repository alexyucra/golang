# Ponteros e Métodos

neste capitulo criamos um nova conta corrente utilizando a palavra new;

Em seguida, comparamos os tipos criados comparando suas referências e entendemos o que são ponteiros na prática;

```go
func main() {
    // metodo 1
    contaAlex := ContaCorrente{
        titular: "Alex",
        saldo:   0,
    }
    fmt.Println(contaAlex)

    // metodo 2
    contaGonzalo := ContaCorrente{"Gonzalo", 543, 123457, 0}
    fmt.Println(contaGonzalo)

    // metodo 3, usando pontero
    var newAcount *ContaCorrente
    newAcount = new(ContaCorrente)
    newAcount.titular = "Alex2"
    newAcount.saldo = 500
    fmt.Println(newAcount)
    fmt.Println(&newAcount) // endereco

    var newAcount2 *ContaCorrente
    newAcount2 = new(ContaCorrente)
    newAcount2.titular = "Alex2"
    newAcount2.saldo = 500
    fmt.Println(newAcount2)
    fmt.Println(&newAcount2) // endereco
    // comparando enderecos = false
    fmt.Println(newAcount == newAcount2)

    // comparando conteudo dos enderecos com apontamento = true
    fmt.Println(*newAcount == *newAcount2)
}
```

Para finalizar, desenvolvemos o método sacar que verifica se o valor do saque é maior do que zero e se a conta possui saldo.

```go
// referente ao this or self en outras lenguagens (c *ContaCorrente)
func (c *ContaCorrente) Sacar(valor float64) string {
    // verificar o sando antes de fazer saque
    // verificar se o valor do saque > 0
    saldo := valor > 0 && valor <= c.saldo
    if saldo {
        c.saldo -= valor
        return "Saque Realizado com sucesso"
    } else {
        return "Saldo insuficiente..."
    }
}
```

<details>
    <summary>show me the code</summary>

```go
package main

import "fmt"

// criamos uma structura, que contem nossa variaveis criadas no main
type ContaCorrente struct {
    titular    string
    numAgencia int
    numConta   int
    saldo      float64
}

// referente ao this or self en outras lenguagens (c *ContaCorrente)
func (c *ContaCorrente) Sacar(valor float64) string {
    // verificar o sando antes de fazer saque
    // verificar se o valor do saque > 0
    saldo := valor > 0 && valor <= c.saldo
    if saldo {
        c.saldo -= valor
        return "Saque Realizado com sucesso"
    } else {
        return "Saldo insuficiente..."
    }
}

func main() {
    // metodo 1
    contaAlex := ContaCorrente{titular: "Alex", saldo: 0}
    fmt.Println(contaAlex)

    // metodo 2
    contaGonzalo := ContaCorrente{"Gonzalo", 543, 123457, 0}
    fmt.Println(contaGonzalo)

    // metodo 3, usando pontero
    var newAcount *ContaCorrente
    newAcount = new(ContaCorrente)
    newAcount.titular = "Alex2"
    newAcount.saldo = 500
    fmt.Println(newAcount)
    fmt.Println(&newAcount) // endereco

    var newAcount2 *ContaCorrente
    newAcount2 = new(ContaCorrente)
    newAcount2.titular = "Alex2"
    newAcount2.saldo = 500
    fmt.Println(newAcount2)
    fmt.Println(&newAcount2) // endereco
    // comparando enderecos = false
    fmt.Println(newAcount == newAcount2)

    // comparando conteudo dos enderecos com apontamento = true
    fmt.Println(*newAcount == *newAcount2)

    // -------------------------------------
    novaConta := ContaCorrente{}
    novaConta.titular = "Silvia"
    novaConta.saldo = 500

    fmt.Println(novaConta.saldo)

    fmt.Println(novaConta.Sacar(-100))
}

```

</details>

___

## Excercicios

- Com base no seguinte Código, , marque as alternativas verdadeiras.

```go
package main

import (
    "fmt"
)

type Conta struct {
    saldo float64
}

func (c *Conta) depositarDezReais() float64 {
    return c.saldo + 10
}

func main() {
    contaTeste := Conta{saldo: 10}

    contaTeste.depositarDezReais()
    contaTeste.depositarDezReais()

    fmt.Println(contaTeste)
}
```

- [ ] Ao executar o código, uma mensagem de erro informará que o saldo está incorreto.
- [ ] Ao executar o código acima, teremos o saldo de 20 como resultado.
- [ ] Ao executar o código acima, teremos o saldo de 30 como resultado.
- [x] Certo! Como não atribuímos o código no saldo da conta( c.saldo = c.saldo + 10), nosso saldo continuará sendo 10.


## Função com quantidade de parâmetros indeterminados

```go
package main

import (
    "fmt"
)

func SemParametro() string {
    return "Exemplo de função sem parâmetro!"
}

func UmParametro(texto string) string {
    return texto
}

func DoisParametros(texto string, numero int) (string, int) {
    return texto, numero
}

// devemos preceder o tipo do do argumento com reticências
func nParametros(numeros ...int) int {
    // soma de n numeros
    resultadoDaSoma := 0
    for _, numero := range numeros {
        resultadoDaSoma += numero
    }
    return resultadoDaSoma
}

func main() {
    fmt.Println(SemParametro())
    fmt.Println(UmParametro("Exemplo de função com um parâmetro"))
    fmt.Println(DoisParametros("Passando dois parâmetros: essa string e o número", 10))

    //
    fmt.Println(nParametros(1))
    fmt.Println(nParametros(1,1))
    fmt.Println(nParametros(1,1,1))
    fmt.Println(nParametros(1,2,3,4))
}

```

Note o uso das reticências na declaração do parâmetro número: numeros ...int. Portanto, podemos criar uma função sem parâmetro, com um, dois, três, ou uma quantidade indeterminada de parâmetros com Go.

[Neste link, você pode executar o código acima.](https://go.dev/play/p/EVwLVFMDB3j)

