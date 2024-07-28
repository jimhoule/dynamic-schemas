package schemas

import (
	"log"
	"main/database"
	"main/database/arango"
	"main/queue"
	"main/queue/topics"
	"main/router"
	"main/schemas/application/services"
	"main/schemas/domain/factories"
	"main/schemas/infrastructure/persistence/arangodb/repositories"
	"main/schemas/presenters/consumer/handlers"
	"main/schemas/presenters/http/controllers"
	"os"
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
	// Sets http routes
	schemasController := &controllers.SchemasController{
		SchemasService: GetService(dbHandler),
	}

	mainRouter.Get("/schemas/{name}", schemasController.GetByName)

	schemasHandler := handlers.SchemasHandler{
		SchemasService: GetService(dbHandler),
	}

	// Sets queue consumer group
	queueConsumerGroupHandler, err := queue.NewConsumerGroupHandler(
		[]string{os.Getenv("QUEUE_ADDRESS")},
		"schemas_consumer_group",
	)
	if err != nil {
		log.Panic(err)
	}

	queueConsumerGroupHandler.Handlers = map[string]queue.Handler{
		topics.TenantCreated: schemasHandler.Create,
	}

	go func() {
		err = queueConsumerGroupHandler.Listen([]string{
			topics.TenantCreated,
		})
		if err != nil {
			log.Panic(err)
		}
	}()
}
