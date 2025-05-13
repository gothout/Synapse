package enterprise

import "time"

type AdminEnterprise struct {
	ID        int64     `json:"id,omitempty"`
	Nome      string    `json:"nome"`
	Cnpj      string    `json:"cnpj"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
