# 

```go
package main

import (
    "net/http"
    "text/template"
)

var tpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
    http.HandleFunc("/", index)
    http.ListenAndServe(":6080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    tpl.ExecuteTemplate(w, "Index", nil)
}
```


<details>
    <summary>`cat templates/index.htm`</summary>

```html
{{define "Index"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"
        crossorigin="anonymous">
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
        crossorigin="anonymous"></script>
    <title>Minha Loja</title>
</head>
<nav class="navbar navbar-light bg-light mb-4">
    <a class="navbar-brand" href="/">Minha Loja</a>
</nav>

<body>
    <div class="container">
        <section class="card">
            <div>
                <table class="table table-striped table-hover mb-0">
                    <thead>
                        <tr>
                            <th>Nome</th>
                            <th>Descrição</th>
                            <th>Preço</th>
                            <th>Quantidade</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>item 1</td>
                            <td>descrião 1</td>
                            <td>29</td>
                            <td>10</td>
                        </tr>
                        <tr>
                            <td>item 2</td>
                            <td>descrião 2</td>
                            <td>1999</td>
                            <td>2</td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </section>

    </div>
</body>

</html>
{{end}}
```

</details>

## criando uma struc - estrutura de produto

```go
type Produto struct {
    Nome string
    Descricao string
    Preco float64
    Quantidade int
}
```

Dentro da funcão index criamos uma lista de produtos

```go
func index(w http.ResponseWriter, r *http.Request) {
    // criamos uma slide de produtos
    produtos := []Produto{
        {Nome:"Item 1", Descricao:"descricao 1", Preco:"", Quantidade:""},
        {"Item 2", "descricao 2", 59, 2},
    }
    // passamos o array de produtos no template
    tpl.ExecuteTemplate(w, "Index", produtos)
}
```

Dentro do template modificamos o array de produtos dinamicamente.

<details>
    <summary>`cat templates/index.htm`</summary>

```html
{{define "Index"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"
        crossorigin="anonymous">
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
        crossorigin="anonymous"></script>
    <title>Minha Loja</title>
</head>
<nav class="navbar navbar-light bg-light mb-4">
    <a class="navbar-brand" href="/">Minha Loja</a>
</nav>

<body>
    <div class="container">
        <section class="card">
            <div>
                <table class="table table-striped table-hover mb-0">
                    <thead>
                        <tr>
                            <th>Nome</th>
                            <th>Descrição</th>
                            <th>Preço</th>
                            <th>Quantidade</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .}}
                        <tr>
                            <td>{{.Nome}}</td>
                            <td>{{.Descricao}}</td>
                            <td>{{.preco}}</td>
                            <td>{{.Quantidade}}</td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </section>

    </div>
</body>

</html>
{{end}}
```

</details>

___