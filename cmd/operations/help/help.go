package help

import "fmt"

type HelperCmd struct{}

func (HelperCmd) Run() {
	helpText := `
✅ COMANDOS DISPONÍVEIS:

  --help       Exibe esta mensagem de ajuda.
  --check-db   Verifica a disponibilidade do banco de dados de ambiente.
  --create-db  Cria as tabelas do banco de dados de ambiente.
  --drop-db    Apaga as tabelas do banco de dados de ambiente.

📝 Exemplo de uso:
Pelo arquivo principal: go run main.go --create-db
Pelo systemd: synapse --create-db
Pelo binario: ./synapse --create-db
`
	fmt.Println(helpText)
}
