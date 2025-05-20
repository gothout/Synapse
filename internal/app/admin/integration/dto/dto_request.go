package integration

import model "Synapse/internal/app/admin/integration/model"

// GetIntegracoesByMarcaIDRequest representa o filtro por marca_id via query ou JSON
type GetIntegracoesByMarcaIDRequest struct {
	MarcaID int64 `form:"marca_id" json:"marca_id" binding:"required"`
}

func (req GetIntegracoesByMarcaIDRequest) ToModelMarca() model.Marca {
	return model.Marca{
		Id: req.MarcaID,
	}
}
