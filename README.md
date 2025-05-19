# 📘 Guia de Desenvolvimento — Sistema Synapse

> Documentação oficial para onboard e manutenção de desenvolvedores no projeto **Synapse**, com foco em organização por **domínio/subdomínio**, controle de permissões, logging inteligente, retorno de erro padronizado e execução por comandos CLI.

---

## 🧱 Estrutura Geral

O projeto segue uma arquitetura **modular com DDD** (Domain-Driven Design), organizada por domínios (ex: `user`, `rule`, `enterprise`) e subdomínios. Cada domínio possui suas próprias camadas.

### 🌳 Árvore Base de Diretórios

```
internal/
└── app/
    └── admin/
        ├── user/
        │   ├── controller/
        │   ├── dto/
        │   ├── model/
        │   ├── repository/
        │   └── service/
        ├── rule/
        ├── enterprise/
        ├── binding/            # Validações manuais
        ├── handler/            # Roteamento por versão
        │   └── v1/
        ├── middleware/         # RBAC e outros
        └── pkg/                # Utils e segurança

configurations/
├── env/                       # Leitura e controle de variáveis .env
├── logger/                    # Gerenciador de logger por nível/ambiente

logger/
└── log_print/                 # Logger com níveis: DEBUG, INFO, etc.
    └── log.go

rest_err/
└── rest_err.go                # Tratamento padronizado de erros

/main.go                    # Entrada do sistema + comandos CLI
```

---

## 🔐 Middleware de Permissão (RBAC)

### 📌 Conceito

Middleware de controle de acesso baseado em regras RBAC:

- Verifica permissões com base no token JWT.
- Permissões são do tipo: `"admin.<subdomínio>", "<ação>"`.

**Exemplo real:**

```go
rbacMiddleware.RequirePermission("admin.user", "read")
```

### 🗂 Localização

```
internal/app/admin/middleware/service/
├── service.go
└── service.impl/
    └── permissions.go
```

---

## ➕ Criar Nova Funcionalidade

### Exemplo: novo módulo `enterprise`

**1. Crie as pastas:**

```
internal/app/admin/enterprise/
├── controller/
├── dto/
├── model/
├── repository/
└── service/
```

**2. Implemente por camada:**

- DTOs: entrada/saída
- Models: estrutura de banco
- Repository: acesso a dados
- Service: regra de negócio
- Controller: rotas e validação

**3. Crie as rotas:**

```go
group := router.Group("/enterprise")
{
	group.POST("/", rbacMiddleware.RequirePermission("admin.enterprise", "create"), ctrl.Create)
	group.GET("/:id", rbacMiddleware.RequirePermission("admin.enterprise", "read"), ctrl.ReadByID)
}
```

---

## 🐞 Logging com Níveis e Ambientes

### 🧠 Como funciona

Logger com controle de ambiente (`DEV`, `PROD`) e nível (`DEBUG`, `INFO`, `WARNING`, `ERROR`, `FATAL`).

**Local:**

```
logger/log_print/log.go
```

**Inicialização no sistema:**

```go
log_print.Init(os.Getenv("ENV"), os.Getenv("LOG"))
```

### ✅ Uso

```go
log_print.Debug("🔍 Executando verificação de usuário")
log_print.Info("✅ Usuário criado com sucesso")
log_print.Warn("⚠️ Tentativa com dados incompletos")
log_print.Error(errors.New("Erro ao consultar banco"))
log_print.Fatal(errors.New("Erro crítico ao inicializar servidor"))
```

**Ambiente PROD:**

- Salva logs no formato:

  ```
  logs/app_PROD_2025-05-19.log
  ```

---

## ❌ Tratamento de Erros com `rest_err`

Retornos consistentes e estruturados em todas as camadas.

**Exemplo de retorno JSON:**

```json
{
  "message": "E-mail inválido",
  "error": "bad_request",
  "code": 400
}
```

### ✅ Funções disponíveis

| Função                           | HTTP Code |
| -------------------------------- | --------- |
| `NewBadRequestError()`           | 400       |
| `NewBadRequestValidationError()` | 400       |
| `NewInternalServerError()`       | 500       |
| `NewNotFoundError()`             | 404       |
| `NewForbiddenError()`            | 403       |

---

## ⚙️ Execução por Comando (CLI via `cmd/`)

**Arquivo de entrada:** `main.go`

O sistema aceita comandos pela linha de terminal com flags:

| Flag          | Função                         |
| ------------- | ------------------------------ |
| `--create-db` | Cria o banco de dados          |
| `--drop-db`   | Apaga o banco atual            |
| `--check-db`  | Verifica a conexão com o banco |
| `--help`      | Lista comandos disponíveis     |

### 🚀 Exemplo prático:

```bash
go run main.go --drop-db --create-db --check-db
```

---

## 🧪 Documentação com Swagger

### 📝 Comentários em Controller:

```go
// @Summary Criação de usuário
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.AdminUserCreateDTO true "Usuário"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Router /admin/v1/user [post]
```

### 🔧 Gerar documentação:

```bash
swag init --parseDependency --exclude internal/test,migrations
```

### 🌐 Acessar no navegador:

```
http://localhost:8080/swagger/index.html
```

---

## 📂 Organização de Permissões

As permissões seguem o padrão:

```bash
"admin.<subdomínio>", "<ação>"
```

### Ações possíveis:

| Ação     | Descrição                   |
| -------- | --------------------------- |
| `create` | Criar recurso               |
| `read`   | Visualizar recurso          |
| `update` | Atualizar recurso existente |
| `remove` | Deletar recurso             |

> Ex: `"admin.user", "update"` representa a permissão de atualização de usuários administrativos.
