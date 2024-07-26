package documents

import (
	"main/database"
	"main/database/arango"
	"main/documents/application/services"
	"main/documents/domain/factories"
	"main/documents/infrastructure/persistence/repositories"
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
	}
}

func Init(mainRouter *router.MainRouter, dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) {
	documentsController := &controllers.DocumentsController{
		DocumentsService: GetService(dbHandler),
	}

	mainRouter.Get("/documents/{schemaId}/{collectionName}", documentsController.GetAll)
	mainRouter.Get("/documents/{schemaId}/{collectionName}/{key}", documentsController.GetByKey)
	mainRouter.Post("/documents", documentsController.Create)
}