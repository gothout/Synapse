# 🧠 Synapse - Arquitetura Escalável em Go (DDD + Clean Architecture)

Este projeto é a base do sistema **Synapse**, uma plataforma modular e expansível, estruturada com foco em **Clean Architecture**, **Domain-Driven Design (DDD)** e boas práticas de organização em Go.

---

## 📦 Estrutura do Projeto

```

internal/
├── app/
│   └── admin/
│       ├── controller/                # Orquestrador geral (ex: admin\_controller.go)
│       ├── handler/                   # Adaptadores HTTP (define rotas por domínio)
│       │   ├── handler.go             # Registro de rotas de admin
│       │   └── user/user.go           # Rotas específicas de usuários
│       ├── binding/                   # Validações e binds
│       │   └── user/user.go
│       ├── user/                      # Domínio: Usuário
│       │   ├── controller/
│       │   ├── dto/
│       │   ├── handler/
│       │   ├── model/
│       │   │   └── user.go
│       │   ├── repository/
│       │   │   ├── user\_repo.go
│       │   │   └── user\_repo\_interface.go
│       │   └── service/
│       │       ├── user\_service.go
│       │       └── user\_service\_interface.go
│       ├── enterprise/               # Domínio: Empresa
│       └── rules/                    # Domínio: Regras/permissões
├── configuration/
│   ├── env/                          # Configurações de ambiente
│   └── rest\_err/                     # Estrutura padrão de retorno de erro
│       ├── causes.go
│       ├── rest\_error.go
│       └── types.go
├── middleware/                       # Middlewares globais (ex: Auth, Log, etc)
├── database/
│   ├── db/                           # Conexão e instanciação do banco
│   └── migrations/                   # Migrations versionadas (golang-migrate)

```

---

## 🧱 Padrões utilizados

- **DDD (Domain-Driven Design)** para organizar por domínios (`user`, `enterprise`, etc.).
- **Clean Architecture** com separação clara entre:
  - `handler` (transporte)
  - `controller` (aplicação)
  - `service` (regra de negócio)
  - `repository` (persistência)
  - `model` (entidade)
  - `dto` (transporte de dados)
- **Go idiomático**, respeitando simplicidade, legibilidade e escalabilidade.

---

## 🛠 Tecnologias

- Linguagem: **Go**
- Banco de dados: **PostgreSQL**
- Migrations: [`golang-migrate`](https://github.com/golang-migrate/migrate)
- Framework HTTP: **Gin**
- Controle de erros: `rest_err`

---

## 📌 Status

🚧 Projeto em desenvolvimento

---
