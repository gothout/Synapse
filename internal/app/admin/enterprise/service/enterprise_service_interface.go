package enterprise

import enterprise "Synapse/internal/app/admin/enterprise/model"

type Service interface {
	Create(enterprise *enterprise.AdminEnterprise) (*enterprise.AdminEnterprise, error)
	ReadByCNPJ(cnpj string) (*enterprise.AdminEnterprise, error)
	ReadByNome(nome string) (*enterprise.AdminEnterprise, error)
	ReadByID(id int64) (*enterprise.AdminEnterprise, error)
	ReadAll(page int) ([]enterprise.AdminEnterprise, error)
	UpdateByCNPJ(cnpj, newCNPJ string, enterprise *enterprise.AdminEnterprise) (*enterprise.AdminEnterprise, error)
	//Update()
	//Delete()
}
