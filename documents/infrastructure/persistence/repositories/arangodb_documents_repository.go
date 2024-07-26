package repositories

import (
	"context"
	"main/database"
	"main/database/arango"
	"main/documents/domain/models"
)

type ArangodbDocumentsRepository struct{
	DbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]
}

func (adr *ArangodbDocumentsRepository) GetAll(schemaId string, collectionName string) ([]*models.Document, error) {
	database, err := adr.DbHandler.Client.Database(context.Background(), schemaId)
	if err != nil {
		return nil, err
	}

	query := "FOR document in @collection return { key: document._key, body: document.body }"
	bindVars := map[string]any{
		"collection": collectionName,
	}

	cursor, err := database.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	documents := []*models.Document{}
	for cursor.HasMore() {
		document := &models.Document{}

		_, err := cursor.ReadDocument(context.Background(), &document)
		if err != nil {
			return nil, err
		}

		documents = append(documents, document)
	}

	return documents, nil
}

func (adr *ArangodbDocumentsRepository) GetByKey(schemaId string, collectionName string, key string) (*models.Document, error) {
	database, err := adr.DbHandler.Client.Database(context.Background(), schemaId)
	if err != nil {
		return nil, err
	}

	query := "FOR document in @collection FILTER document._key == @key return { key: document._key, body: document.body }"
	bindVars := map[string]any{
		"collection": collectionName,
		"key": key,
	}

	cursor, err := database.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	document := &models.Document{}

		_, err = cursor.ReadDocument(context.Background(), &document)
		if err != nil {
			return nil, err
		}


	return document, nil
}

func (adr *ArangodbDocumentsRepository) Create(schemaId string, collectionName string, document *models.Document) (*models.Document, error) {
	database, err := adr.DbHandler.Client.Database(context.Background(), schemaId)
	if err != nil {
		return nil, err
	}

	query := "INSERT { _key: @key, body: @body } INTO @collection"
	bindVars := map[string]any{
		"collection": collectionName,
		"key": document.Key,
		"body": document.Body,
	}

	_, err = database.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}
	
	return document, nil
}