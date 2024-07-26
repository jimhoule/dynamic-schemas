package schemas

import (
	"main/database"
	"main/database/arango"
	"main/router"
	"main/schemas/application/services"
	"main/schemas/domain/factories"
	"main/schemas/infrastructure/persistence/repositories"
	"main/schemas/presenters/http/controllers"
)

func GetService(dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) *services.SchemasService {
	return &services.SchemasService{
		SchemasFactory: &factories.SchemasFactory{},
		SchemasRepository: &repositories.ArangodbSchemasRepository{
			DbHandler: dbHandler,
		},
	}
}

func Init(mainRouter *router.MainRouter, dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) {
	schemasController := &controllers.SchemasController{
		SchemasService: GetService(dbHandler),
	}

	mainRouter.Get("/schemas/{id}", schemasController.GetById)
	mainRouter.Post("/schemas", schemasController.Create)
}