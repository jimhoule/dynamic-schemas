package repositories

import (
	"context"
	"fmt"
	"main/database"
	"main/database/arango"
	"main/documents/domain/models"
)

type ArangodbDocumentsRepository struct {
	DbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]
}

func (adr *ArangodbDocumentsRepository) GetAll(schemaId string, collectionName string) ([]*models.Document, error) {
	database, err := adr.DbHandler.Client.Database(context.Background(), schemaId)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(
		"FOR document IN %s RETURN { key: document._key, body: document.body }",
		collectionName,
	)
	bindVars := map[string]any{}

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

	query := fmt.Sprintf(
		"FOR document IN %s FILTER document._key == @key RETURN { key: document._key, body: document.body }",
		collectionName,
	)
	bindVars := map[string]any{
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

	query := fmt.Sprintf("INSERT { _key: @key, body: @body } INTO %s", collectionName)
	bindVars := map[string]any{
		"key":  document.Key,
		"body": document.Body,
	}

	_, err = database.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}

	return document, nil
}
