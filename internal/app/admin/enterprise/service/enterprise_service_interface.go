package enterprise

import enterprise "Synapse/internal/app/admin/enterprise/model"

type Service interface {
	Create(enterprise *enterprise.AdminEnterprise) (*enterprise.AdminEnterprise, error)
	Read()
	Update()
	Delete()
}
