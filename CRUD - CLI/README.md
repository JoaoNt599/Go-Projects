# CRUD CLI em Go


Este projeto é uma aplicação de linha de comando (CLI) para gerenciar itens utilizando a linguagem Go.
O Objetivo é realizar um aprendizado contínuo a partir de uma aplicação simples.


## 🛠️ Tecnologias

- Go
- Testify (para testes)
- JSON para armazenamento
- Clean Architecture
- Docker

## Estrutura do Projeto

### CRUD-CLI/
### ├── cmd/                  # Interface CLI
### │   └── main.go
### ├── docker/               # Docker Configuration
### │   ├── Dockerfile        # Dockerfile para criação da imagem
### │   └── .dockerignore     # Arquivos a serem ignorados pelo Docker
### ├── internal/             # Camada de Domínio e Casos de Uso
### │   ├── domain/           # Entidades
### │   │   └── item.go
### │   ├── usecase/          # Casos de Uso
### │   │   ├── item_service.go
### │   │   ├── interfaces.go
### │   │   └── item_service_test.go
### ├── pkg/                  # Infraestrutura
### │   └── local/
### │       ├── repository.go
### │       ├── storage.go
### │       └── repository_test.go
### └── items.json            # Arquivo de armazenamento
### └── go.mod                # Dependências


## Funcionalidades 

- Criar Item
- Listar Item
- Atualizar Item (ainda não implementado)
- deletar item (ainda não implementado)


## Execute o Projeto

1. Clone o repositório:

        git clone <url>
        cd crud-cli

2. Instale as dependências:

    go mod tidy

3. Execute o programa:

    go run cmd/main.go

## Docker

1. Build da imagem:

    docker build -t crud-cli -f docker/Dockerfile .

2. Execute o container:

    docker run --rm crud-cli

## Testes

Executando os testes:

    go test ./...