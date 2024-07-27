package factories

import (
	"main/schemas/domain/models"
)

type SchemasFactory struct{}

func (*SchemasFactory) Create(name string) *models.Schema {
	return &models.Schema{
		Name: name,
	}
}
