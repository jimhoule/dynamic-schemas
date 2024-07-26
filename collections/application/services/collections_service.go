package services

import (
	"main/collections/application/payloads"
	"main/collections/application/ports"
	"main/collections/domain/factories"
	"main/collections/domain/models"
)

type CollectionsService struct {
	CollectionsFactory    *factories.CollectionsFactory
	PropertiesFactory     *factories.PropertiesFactory
	CollectionsRepository ports.CollectionsRepositoryPort
}

func (cs *CollectionsService) GetAll(schemaId string) ([]*models.Collection, error) {
	return cs.CollectionsRepository.GetAll(schemaId)
}

func (cs *CollectionsService) GetByName(schemaId string, name string) (*models.Collection, error) {
	return cs.CollectionsRepository.GetByName(schemaId, name)
}

func (cs *CollectionsService) Create(createCollectionPayload *payloads.CreateCollectionPayload) (*models.Collection, error) {
	properties := []*models.Property{}
	for _, createPropertyPayload := range createCollectionPayload.CreatePropertyPayloads {
		property := cs.PropertiesFactory.Create(
			createCollectionPayload.Name,
			createPropertyPayload.IsRequired,
			createPropertyPayload.Type,
		)

		properties = append(properties, property)
	}

	collection := cs.CollectionsFactory.Create(createCollectionPayload.Name, properties)

	return cs.CollectionsRepository.Create(createCollectionPayload.SchemaId, collection)
}
