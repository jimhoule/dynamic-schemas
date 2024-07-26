package factories

import (
	"main/collections/domain/models"
)

type CollectionsFactory struct{}

func (cf *CollectionsFactory) Create(name string, properties []*models.Property) *models.Collection {
	return &models.Collection{
		Name:       name,
		Properties: properties,
	}
}
