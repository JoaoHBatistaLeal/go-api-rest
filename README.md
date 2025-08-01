
# 🍕 Pizzaria API - Meu Primeiro Projeto em Go 🚀

API REST desenvolvida em Go, como parte da minha jornada de aprendizado em backend. Projeto focado em boas práticas, organização em camadas e separação de responsabilidades, ideal para demonstrar domínio técnico em um ambiente profissional.

## 📚 Descrição
Essa API é uma aplicação backend desenvolvida em Golang, utilizando o framework Gin para gerenciamento de pizzas e avaliações. A persistência é feita via arquivo JSON, simulando uma base de dados local.

Esse projeto foi desenvolvido como parte do curso APIs REST com Go e tem como objetivo aplicar conceitos fundamentais do backend, como:

- Criação de endpoints RESTful

- Organização modular do código

- Validação e manipulação de dados

- Preparação para escalabilidade futuro
## 🧱 Estrutura do projeto

```bash
.
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── dados/
│   └── pizza.json              # "Base de dados" local em JSON
├── internal/
│   ├── data/                   # Lógica de leitura/escrita dos dados
│   ├── handler/                # Handlers e rotas HTTP
│   ├── models/                 # Structs das entidades (Pizza e Avaliação)
│   └── service/                # Regras de negócio

 ```

## 📡 Endpoints da API

| Método | Rota                            | Descrição                         |
|--------|----------------------------------|------------------------------------|
| `GET`  | `/pizzas`                        | Retorna todas as pizzas cadastradas |
| `POST` | `/pizzas`                        | Cria uma nova pizza                |
| `GET`  | `/pizzas/:id`                    | Busca uma pizza específica por ID  |
| `PUT`  | `/pizzas/:id`                    | Atualiza os dados de uma pizza     |
| `DELETE` | `/pizzas/:id`                  | Remove uma pizza existente         |
| `POST` | `/pizzas/:id/avaliacoes`         | Adiciona avaliação à pizza         |


## ⚙️ Tecnologias e ferramentas

- Go 1.22.5
- Gin Framework
- JSON

## 🚀 Como executar localmente

1. Clone o repositório na máquina:
```bash 
git clone https://github.com/seu-usuario/pizzaria-api
cd pizzaria-api

```

2. Excecute a aplicação:
 ```bash 
go run main.go
 ```

3. Acesse os endpoints com Postman ou Navegador.

## 📈 Próximos passos
- 🚧 Validação de dados
- 🚧 Dockerização da aplicação
- 🚧 Utilização de Banco de Dados para persistência
- 🚧 Criação de documentação Swagger

## 👨‍🎓 Sobre o Autor

API desenvolvida por João durante o curso de API Rest em Go pela Alura, sendo meu primeiro projeto prático backend utilizando a linguagem Go

### 📘 Cursos 

- Go: criando uma API Rest - Alura, Certificado disponível em: https://cursos.alura.com.br/user/jhbatistaleal/course/go-criando-api-rest/certificate

 - Go: gerenciando e otimizando sua API - Alura, Certificado disponível em: https://cursos.alura.com.br/user/jhbatistaleal/course/go-gerenciando-otimizando-api/certificate

 ## 🔍 Para recrutadores e colegas desenvolvedores

 Se chegou até aqui, te convido a conferir o código e deixar seu feedback. Este é um projeto acadêmico com propósito de aprendizado, mas seguindo as boas práticas da linguagem

Conecte-se comigo em: https://www.linkedin.com/in/jo%C3%A3o-leal-8ba769307/






