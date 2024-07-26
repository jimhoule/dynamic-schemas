package factories

import (
	"main/schemas/domain/models"
)

type SchemasFactory struct{}

func (*SchemasFactory) Create(id string) *models.Schema {
	return &models.Schema{
		Id: id,
	}
}
