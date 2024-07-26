package ports

import "main/documents/domain/models"

type DocumentsRepositoryPort interface {
	GetAll(schemaId string, collectionName string) ([]*models.Document, error)
	GetByKey(schemaId string, collectionName string, key string) (*models.Document, error)
	Create(schemaId string, collectionName string, document *models.Document) (*models.Document, error)
}