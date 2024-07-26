package ports

import "main/collections/domain/models"

type CollectionsRepositoryPort interface {
	GetAll(schemaId string) ([]*models.Collection, error)
	GetByName(schemaId string, name string) (*models.Collection, error)
	Create(schemaId string, collection *models.Collection) (*models.Collection, error)
}