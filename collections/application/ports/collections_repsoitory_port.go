package ports

import "main/collections/domain/models"

type CollectionsRepositoryPort interface {
	GetAll(schemaName string) ([]*models.Collection, error)
	GetByName(schemaName string, name string) (*models.Collection, error)
	Create(schemaName string, collection *models.Collection) (*models.Collection, error)
}
