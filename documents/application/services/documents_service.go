package services

import (
	"main/documents/application/payloads"
	"main/documents/application/ports"
	"main/documents/domain/factories"
	"main/documents/domain/models"
)

type DocumentsService struct{
	DocumentsFactory *factories.DocumentsFactory
	DocumentsRepository ports.DocumentsRepositoryPort
}

func (ds *DocumentsService) GetAll(schemaId string, collectionName string) ([]*models.Document, error) {
	return ds.DocumentsRepository.GetAll(schemaId, collectionName)
}

func (ds *DocumentsService) GetByKey(schemaId string, collectionName string, key string) (*models.Document, error) {
	return ds.DocumentsRepository.GetByKey(schemaId, collectionName, key)
}

func (ds *DocumentsService) Create(createDocumentPayload *payloads.CreateDocumentPayload) (*models.Document, error) {
	document := ds.DocumentsFactory.Create(createDocumentPayload.Body)

	return ds.DocumentsRepository.Create(
		createDocumentPayload.SchemaId,
		createDocumentPayload.CollectionName,
		document,
	)
}