package repositories

import (
	"context"
	"main/database"
	"main/database/arango"
	"main/tenants/domain/models"
)

type ArangodbTenantsRepository struct{
	DbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]
}

func (atr *ArangodbTenantsRepository) GetById(id string) (*models.Tenant, error) {
	query := "For tenant IN tenants Filter tenant._key == @key RETURN tenant"
	bindVars := map[string]any{
		"key": id,
	}

	cursor, err := atr.DbHandler.Database.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}
	cursor.Close()

	tenant := &models.Tenant{}
	_, err = cursor.ReadDocument(context.Background(), &tenant)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}

func (atr *ArangodbTenantsRepository) Create(tenant *models.Tenant) (*models.Tenant, error) {
	query := "INSERT { _key: @key, id: @id, name: @name } INTO tenants"
	bindVars := map[string]any{
		"key": tenant.Id,
		"id": tenant.Id,
		"name": tenant.Name,
	}

	_, err := atr.DbHandler.Database.Query(context.Background(), query, bindVars)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}