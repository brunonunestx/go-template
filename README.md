# Go Template API

Template de API REST em Go seguindo princípios de Clean Architecture e boas práticas de desenvolvimento.

## 📋 Índice

- [Arquitetura](#arquitetura)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Tecnologias](#tecnologias)
- [Pré-requisitos](#pré-requisitos)
- [Como Rodar o Projeto](#como-rodar-o-projeto)
- [Endpoints](#endpoints)

## 🏗️ Arquitetura

Este projeto segue os princípios da **Clean Architecture**, organizando o código em camadas bem definidas com responsabilidades específicas:

### Camadas da Aplicação

```
┌─────────────────────────────────────┐
│         HTTP Layer (Handlers)       │  ← Entrypoint da aplicação
├─────────────────────────────────────┤
│         Services Layer              │  ← Lógica de negócio
├─────────────────────────────────────┤
│         Repository Layer            │  ← Acesso a dados
├─────────────────────────────────────┤
│         Domain Layer                │  ← Entidades e regras
└─────────────────────────────────────┘
```

#### 1. **Domain Layer** (`internal/domain/`)
- Contém as entidades de negócio e interfaces
- Define contratos (interfaces) que serão implementados pelas outras camadas
- Independente de frameworks e bibliotecas externas
- Exemplo: `User`, `Product`, erros customizados

#### 2. **Repository Layer** (`internal/repository/`)
- Implementa a persistência de dados
- Responsável pela comunicação com banco de dados
- Implementa interfaces definidas no Domain
- Abstrai detalhes de implementação de storage

#### 3. **Services Layer** (`internal/services/`)
- Contém a lógica de negócio da aplicação
- Orquestra operações entre repositories
- Implementa regras de negócio complexas
- Validações e transformações de dados

#### 4. **HTTP Layer** (`internal/http/`)
- Handlers HTTP que recebem as requisições
- Middlewares para cross-cutting concerns (autenticação, logging, etc.)
- Responses padronizadas
- Roteamento de requisições

### Benefícios da Arquitetura

- ✅ **Testabilidade**: Camadas isoladas facilitam testes unitários
- ✅ **Manutenibilidade**: Código organizado e fácil de localizar
- ✅ **Escalabilidade**: Fácil adicionar novas funcionalidades
- ✅ **Independência**: Mudanças em uma camada não afetam outras
- ✅ **Flexibilidade**: Fácil trocar implementações (ex: mudar de banco de dados)

## 📁 Estrutura do Projeto

```
.
├── cmd/
│   └── api/
│       └── main.go              # Entry point da aplicação
├── internal/
│   ├── config/
│   │   └── config.go            # Configurações da aplicação
│   ├── db/
│   │   └── db.md                # Documentação do banco de dados
│   ├── domain/
│   │   ├── errors.go            # Erros customizados
│   │   └── user.go              # Entidade User
│   ├── http/
│   │   ├── router.go            # Configuração de rotas
│   │   ├── handlers/            # HTTP handlers
│   │   │   └── health_handler.go
│   │   ├── middlewares/         # Middlewares HTTP
│   │   └── responses/           # Padronização de respostas
│   ├── repository/              # Camada de acesso a dados
│   │   └── repository.md
│   ├── services/                # Lógica de negócio
│   │   └── services.md
│   └── utils/                   # Utilitários gerais
├── pkg/                         # Pacotes reutilizáveis
├── scripts/                     # Scripts auxiliares
├── api/                         # Documentação da API (Swagger, etc)
├── docker-compose.yml           # Orquestração de containers
├── Dockerfile                   # Imagem Docker da aplicação
├── go.mod                       # Dependências Go
└── README.md                    # Este arquivo
```

### Convenções de Nomenclatura

- **internal/**: Código privado da aplicação (não exportável)
- **pkg/**: Código que pode ser importado por outras aplicações
- **cmd/**: Entry points da aplicação
- **api/**: Documentação de API (OpenAPI/Swagger)

## 🚀 Tecnologias

- **Go 1.25.4**: Linguagem de programação
- **net/http**: Biblioteca padrão para HTTP server
- **Docker**: Containerização da aplicação
- **Docker Compose**: Orquestração de containers

## 📦 Pré-requisitos

Antes de começar, você precisa ter instalado:

- [Go 1.25.4+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## 🎯 Como Rodar o Projeto

### Opção 1: Rodando Localmente (sem Docker)

1. **Clone o repositório**
```bash
git clone <url-do-repositorio>
cd go-template
```

2. **Instale as dependências**
```bash
go mod download
```

3. **Execute a aplicação**
```bash
go run cmd/api/main.go
```

4. **Acesse a aplicação**
```
http://localhost:8080
```

### Opção 2: Rodando com Docker

1. **Clone o repositório**
```bash
git clone <url-do-repositorio>
cd go-template
```

2. **Construa e execute com Docker Compose**
```bash
docker-compose up --build
```

3. **Acesse a aplicação**
```
http://localhost:8080
```

### Opção 3: Rodando apenas com Docker (sem Compose)

1. **Build da imagem**
```bash
docker build -t go-api .
```

2. **Execute o container**
```bash
docker run -p 8080:8080 --name go-api-container go-api
```

3. **Acesse a aplicação**
```
http://localhost:8080
```

## 🔗 Endpoints

### Health Check

Verifica se a aplicação está rodando corretamente.

```bash
GET /health
```

**Resposta de sucesso:**
```
ok
```

**Exemplo usando curl:**
```bash
curl http://localhost:8080/health
```

### Próximos endpoints

Os seguintes endpoints estão planejados:
- `GET /users` - Listar usuários
- `POST /users` - Criar usuário
- `GET /users/:id` - Buscar usuário por ID
- `PUT /users/:id` - Atualizar usuário
- `DELETE /users/:id` - Deletar usuário

## 🛠️ Desenvolvimento

### Rodando em modo desenvolvimento

```bash
# Com hot reload (usando air)
go install github.com/cosmtrek/air@latest
air
```

### Build para produção

```bash
go build -o server ./cmd/api
./server
```

### Rodando testes

```bash
# Todos os testes
go test ./...

# Com coverage
go test -cover ./...

# Coverage detalhado
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📝 Variáveis de Ambiente

| Variável | Descrição | Padrão |
|----------|-----------|--------|
| `APP_ENV` | Ambiente da aplicação (development, production) | development |
| `PORT` | Porta onde a aplicação irá rodar | 8080 |

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT.

## 👤 Autor

Bruno Nunes - [@brunonunestx](https://github.com/brunonunestx)
