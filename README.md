<div align="center">

# 🛍️ Gerenciador de Produtos

### Sistema CRUD Full-Stack para Cadastro de Produtos

*Desafio Técnico - Desenvolvedor(a) Júnior | LAIOB*

![React](https://img.shields.io/badge/React-18.x-61DAFB?style=flat-square&logo=react)
![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat-square&logo=go)
![SQLite](https://img.shields.io/badge/SQLite-3-003B57?style=flat-square&logo=sqlite)
![Docker](https://img.shields.io/badge/Docker-Enabled-2496ED?style=flat-square&logo=docker)

</div>

## 📋 Sobre o Projeto

Sistema completo de gerenciamento de produtos desenvolvido com **React** no frontend e **Go** no backend. A aplicação permite realizar operações CRUD (Create, Read, Update, Delete) de produtos de forma intuitiva e responsiva.

### ✨ Funcionalidades

- ✅ **Listar produtos** com interface em cards elegantes
- ✅ **Cadastrar novos produtos** via formulário responsivo
- ✅ **Editar produtos existentes** com preenchimento automático
- ✅ **Excluir produtos** com confirmação de segurança
- ✅ **Interface responsiva** que funciona em desktop e mobile
- ✅ **Tratamento de erros** com alertas informativos
- ✅ **Banco de dados** com dados de exemplo pré-carregados

🎥 [Assista ao vídeo da aplicação](https://youtu.be/Hec_Ekce_TU)
- [Apresentação no Canva (com vídeos)](https://www.canva.com/design/DAGzh_49k2Y/eSdN2_t8M6svVA1vY021sg/edit?utm_content=DAGzh_49k2Y&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton)

## 🚀 Tecnologias Utilizadas

### Frontend
- **React** 18.x - Biblioteca para interfaces de usuário
- **Axios** 1.x - Cliente HTTP para requisições
- **CSS3** - Estilização com Grid e Flexbox
- **HTML5** - Estruturação semântica

### Backend
- **Go** 1.24 - Linguagem de programação
- **Gin** 1.10.0 - Framework web para Go
- **SQLite** 1.14.x - Banco de dados relacional
- **CORS** 1.7.0 - Habilitação de requisições cross-origin

### DevOps
- **Docker** - Containerização da aplicação
- **Docker Compose** - Orquestração de containers
- **Multi-stage builds** - Otimização das imagens

## 🏗️ Estrutura do Projeto

desafio-desenvolvedora-laiob/

├── 📁 controller/ # Controladores da API

├── 📁 database/ # Configuração do banco de dados

├── 📁 model/ # Modelos de dados (structs)

├── 📁 frontend/ # Aplicação React

│ ├── 📁 src/ # Código-fonte React

│ ├── 📁 public/ # Arquivos estáticos

│ └── 📄 package.json # Dependências do frontend

├── 📄 main.go # Arquivo principal da aplicação

├── 📄 produto.go # Lógica de negócio dos produtos

├── 📄 go.mod # Dependências do Go

├── 📄 docker-compose.yml # Orquestração dos containers

├── 📄 Dockerfile # Container do backend

├── 📄 produtos.db # Banco de dados SQLite

└── 📄 README.md # Documentação do projeto


## ⚙️ Como Executar o Projeto

### 🐳 Método 1: Docker (Recomendado)

**Pré-requisitos:**
- Docker Desktop instalado
- Git para clonar o repositório

**Passos:**

1. **Clone o repositório:**
```
git clone https://github.com/belzebells/desafio-desenvolvedora-laiob.git
cd desafio-desenvolvedora-laiob
```

2. **Execute com Docker:**
````
docker-compose up --build
````

3. **Acesse a aplicação:**
- **Frontend:** http://localhost:3000
- **Backend API:** 
  http://localhost:8080/produtos
  
  ✅ http://localhost:8080/produtos/3
  
  ❌ http://localhost:8080 (404 é normal - use os endpoints acima)

### 💻 Método 2: Execução Local

**Pré-requisitos:**
- Go 1.24+ instalado
- Node.js 18+ instalado
- SQLite3 instalado

**Backend (Terminal 1):**
Na pasta raiz do projeto

```
go mod download
go run main.go
```

**Frontend (Terminal 2):**
Na pasta frontend

````
cd frontend
npm install
npm start
````

## 🗄️ Banco de Dados

### Configuração Automática

O banco SQLite é criado automaticamente na primeira execução com a seguinte estrutura:
CREATE TABLE produtos (
id INTEGER PRIMARY KEY AUTOINCREMENT,
nome TEXT NOT NULL,
preco REAL NOT NULL,
quantidade INTEGER NOT NULL,
descricao TEXT NOT NULL
);



### Dados de Exemplo

A aplicação já vem com produtos pré-cadastrados:
- Mouse Gamer - R$ 199,90
- PC Gamer - R$ 1.199,90  
- Notebook - R$ 4.500,99
- Mouse Gamer Pro - R$ 249,90

## 🌐 API Endpoints

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| GET | `/produtos` | Lista todos os produtos |
| GET | `/produtos/:id` | Retorna um produto específico |
| POST | `/produtos` | Cria um novo produto |
| PUT | `/produtos/:id` | Atualiza um produto existente |
| DELETE | `/produtos/:id` | Remove um produto |

### Exemplo de Requisição:

POST /produtos
{
"nome": "Teclado Mecânico",
"preco": 299.99,
"quantidade": 15,
"descricao": "Teclado mecânico com switches Cherry MX"
}


## 🔧 Scripts Disponíveis

### Frontend
- `npm start` - Executa em modo de desenvolvimento
- `npm build` - Cria build de produção
- `npm test` - Executa testes

### Backend
- `go run main.go` - Executa o servidor
- `go build` - Compila o binário
- `go test ./...` - Executa testes

### Docker
- `docker-compose up` - Inicia os containers
- `docker-compose down` - Para os containers
- `docker-compose build --no-cache` - Reconstrói sem cache

## 🎯 Diferenciais Implementados

- ✅ **Docker** para backend e frontend
- ✅ **Layout responsivo** adaptável
- ✅ **Tratamento de erros** completo
- ✅ **Interface moderna** e intuitiva
- ✅ **Documentação detalhada**
- ✅ **Código organizado** e comentado

## 👨‍💻 Autor

Desenvolvido com 💙 por **belzebells**

- GitHub: [@belzebells](https://github.com/belzebells)
- Projeto: [desafio-desenvolvedora-laiob](https://github.com/belzebells/desafio-desenvolvedora-laiob)

---

<div align="center">

**⭐ Se este projeto te ajudou, deixe uma estrela!**

</div>


