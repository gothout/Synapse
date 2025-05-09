package enterprise

import enterprise "Synapse/internal/app/admin/enterprise/model"

type Repository interface {
	Create(enterprise *enterprise.AdminEnterprise) (*enterprise.AdminEnterprise, error)
}
