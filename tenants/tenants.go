package tenants

import (
	"main/database"
	"main/database/arango"
	"main/router"
	"main/tenants/application/services"
	"main/tenants/domain/factories"
	"main/tenants/infrastructure/persistence/repositories"
	"main/tenants/presenters/http/controllers"
	"main/uuid"
)

func GetService(dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) *services.TenantsService {
	return &services.TenantsService{
		TenantsFactory: &factories.TenantsFactory{
			UuidService: uuid.GetService(),
		},
		TenantsRepository: &repositories.ArangodbTenantsRepository{
			DbHandler: dbHandler,
		},
	}
}

func Init(mainRouter *router.MainRouter, dbHandler *database.DbHandler[arango.DriverDatabase, arango.DriverClient]) {
	tenantsController := &controllers.TenantsController{
		TenantsService: GetService(dbHandler),
	}

	mainRouter.Get("/tenants/{id}", tenantsController.GetById)
	mainRouter.Post("/tenants", tenantsController.Create)
}