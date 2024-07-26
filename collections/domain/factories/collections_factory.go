package factories

import (
	"main/collections/domain/models"
	documentModels "main/documents/domain/models"
)

type CollectionsFactory struct{}

func (cf *CollectionsFactory) Create(name string, properties map[string]*models.Property) *models.Collection {
	return &models.Collection{
		Name: name,
		Properties: properties,
		Documents: map[string]*documentModels.Document{},
	}
}