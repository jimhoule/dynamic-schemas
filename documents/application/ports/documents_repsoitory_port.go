package ports

import (
	"main/documents/domain/models"
)

type DocumentsRepositoryPort interface {
	GetAll(schemaName string, collectionName string) ([]*models.Document, error)
	GetByKey(schemaName string, collectionName string, key string) (*models.Document, error)
	Create(schemaName string, collectionName string, document *models.Document) (*models.Document, error)
}
