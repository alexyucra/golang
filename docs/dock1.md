# Build and Push Docker Images with Go

Install SDK do docker [https://docs.docker.com/engine/api/sdk/](https://docs.docker.com/engine/api/sdk/)
Na ultima versao ficou como incompativel, siga os seguintes passos para instalar os modulos do docker

> $ go install github.com/docker/docker/client@v18.06.1-ce

> $ go env GOPATH 

> $ go mod tidy 

esto criará os arquivos `go.mod` e `go.sum`na mesma pasta do nosso projeto. [link da documentacão](https://pkg.go.dev/github.com/docker/docker/client#pkg-examples)

```shell
.
├── go.mod
├── go.sum
└── main.go
```

## Listando container con GO

agora no arquivo `main.go` isto listara todos os containers.

```go
package main

import (
    "context"
    "fmt"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func main() {
    cli, err := client.NewClientWithOpts(client.FromEnv)
    if err != nil {
        panic(err)
    }

    containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
    if err != nil {
        panic(err)
    }

    for _, container := range containers {
        fmt.Printf("%s %s\n", container.ID[:10], container.Image)
    }
}
```

## Criando um ambiente com Dockerfile

Primeiro, precisamos configurar o ambiente. Crie um projeto e inclua o aplicativo que queremos conteinerizar:

> mkdir docker-go-tutorial && cd docker-go-tutorial && mkdir node-hello

Adicionaremos um aplicativo Node.js simples:

```js
// node-hello/app.js
console.log("Hello From LoginRadius");
```

com o Dockerfile:

```dockerfile
# node-hello/Dockerfile
FROM node:12
WORKDIR /src
COPY . .
CMD [ "node", "app.js" ]
```

agora criaremos nosso arquivo `main.go`, nosso projeto debe ficar asim:

```shell
.
├── go.mod
├── go.sum
├── main.go
└── node-hello
    ├── Dockerfile
    └── app.js
```

Em seguida, instale o Go SDK . Estas são as importações relacionadas ao Docker que usaremos:

```go
"github.com/docker/docker/api/types"
"github.com/docker/docker/client"
"github.com/docker/docker/pkg/archive"
```

## Criando uma imagen Docker

Uma maneira de construir uma imagem do Docker a partir de nossos arquivos locais é compactar esses arquivos em um arquivo tar primeiro.
para isso usaremos o pacote de arquivo fornecido pelo Docker: `github.com/docker/docker/pkg/archive`

```go
func imageBuild(dockerClient *client.Client) error {
    tar, err := archive.TarWithOptions("node-hello/", &archive.TarOptions{})
    if err != nil {
        return err
    }
}
```

Agora, podemos chamar a função ImageBuild usando o Go SDK:

- Observe que a tag de imagem inclui nosso ID de usuário do registro do Docker, para que possamos enviar essa imagem para nosso registro posteriormente.

```go
func imageBuild(dockerClient *client.Client) error {
    tar, err := archive.TarWithOptions("node-hello/", &archive.TarOptions{})
    if err != nil {
        return err
    }
    // aqui a tag com ID Registro docker
    opts := types.ImageBuildOptions{
        Dockerfile:  "Dockerfile",
        Tags:        []string{dockerRegistryUserID + "/node-hello"},
        Remove:      true,
    }

    res, err := dockerClient.ImageBuild(ctx, tar, opts)
    if err != nil {
        return err
    }
}
```

- Para imprimir a resposta, usamos um scanner para percorrer linha por linha:

```go
func imageBuild(dockerClient *client.Client) error {
    tar, err := archive.TarWithOptions("node-hello/", &archive.TarOptions{})
    if err != nil {
        return err
    }
    // 
    opts := types.ImageBuildOptions{
        Dockerfile:  "Dockerfile",
        Tags:        []string{dockerRegistryUserID + "/node-hello"},
        Remove:      true,
    }

    res, err := dockerClient.ImageBuild(ctx, tar, opts)
    if err != nil {
        return err
    }
    // percorremos linha a linha
    scanner := bufio.NewScanner(res.Body)
    for scanner.Scan() {
        lastLine = scanner.Text()
        fmt.Println(scanner.Text())
    }
}
```