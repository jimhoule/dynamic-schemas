package tenants

import (
	"fmt"
	"log"
	"main/database"
	"main/database/arango"
	"main/queue"
	"main/router"
	"main/tenants/application/services"
	"main/tenants/domain/factories"
	"main/tenants/infrastructure/persistence/arangodb/repositories"
	"main/tenants/presenters/http/controllers"
	"main/uuid"
	"os"
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
	queueProducerHandler, err := queue.NewProducerHandler(
		[]string{
			fmt.Sprintf("%s:%s", os.Getenv("QUEUE_URL"), os.Getenv("QUEUE_PORT")),
		},
	)
	if err != nil {
		log.Panic(err)
	}

	tenantsController := &controllers.TenantsController{
		TenantsService:       GetService(dbHandler),
		QueueProducerHandler: &queueProducerHandler,
	}

	mainRouter.Get("/tenants/{id}", tenantsController.GetById)
	mainRouter.Post("/tenants", tenantsController.Create)
}
