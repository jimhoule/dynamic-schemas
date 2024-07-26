package factories

import (
	collectionModels "main/collections/domain/models"
	"main/schemas/domain/models"
)

type SchemasFactory struct{}

func (*SchemasFactory) Create(id string) *models.Schema {
	return &models.Schema{
		Id: id,
		Collections: map[string]*collectionModels.Collection{},
	}
}