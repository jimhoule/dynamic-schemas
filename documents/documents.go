package documents

import (
	"main/collections"
	"main/database"
	"main/database/arango"
	"main/documents/application/services"
	"main/documents/domain/factories"
	"main/documents/infrastructure/persistence/arangodb/repositories"
	"main/documents/presenters/http/controllers"
	"main/router"
	"main/uuid"
)

func GetService(dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) *services.DocumentsService {
	return &services.DocumentsService{
		DocumentsFactory: &factories.DocumentsFactory{
			UuidService: uuid.GetService(),
		},
		DocumentsRepository: &repositories.ArangodbDocumentsRepository{
			DbHandler: dbHandler,
		},
		CollectionsService: collections.GetService(dbHandler),
	}
}

func Init(mainRouter *router.MainRouter, dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) {
	documentsController := &controllers.DocumentsController{
		DocumentsService: GetService(dbHandler),
	}

	mainRouter.Get("/documents/{schemaName}/{collectionName}", documentsController.GetAll)
	mainRouter.Get("/documents/{schemaName}/{collectionName}/{key}", documentsController.GetByKey)
	mainRouter.Post("/documents", documentsController.Create)
}
