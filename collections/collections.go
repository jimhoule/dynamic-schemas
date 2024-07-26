package collections

import (
	"main/collections/application/services"
	"main/collections/domain/factories"
	"main/collections/infrastructure/persistence/repositories"
	"main/collections/presenters/http/controllers"
	"main/database"
	"main/database/arango"
	"main/router"
)

func GetService(dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) *services.CollectionsService {
	return &services.CollectionsService{
		CollectionsFactory: &factories.CollectionsFactory{},
		PropertiesFactory: &factories.PropertiesFactory{},
		CollectionsRepository: &repositories.ArangodbCollectionsRepository{
			DbHandler: dbHandler,
		},
	}
}

func Init(mainRouter *router.MainRouter, dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) {
	collectionsController := &controllers.CollectionsController{
		CollectionsService: GetService(dbHandler),
	}

	mainRouter.Get("/collections/{schemaId}", collectionsController.GetAll)
	mainRouter.Get("/collections/{schemaId}/{name}", collectionsController.GetByName)
	mainRouter.Post("/collections", collectionsController.Create)
}