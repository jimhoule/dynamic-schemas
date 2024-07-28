package arango

import (
	"context"
	"fmt"
	"main/database"
	"os"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type DriverDatabase = driver.Database
type DriverClient = driver.Client
type DriverCollection = driver.Collection
type CreateCollectionOptions = driver.CreateCollectionOptions
type CreateCollectionPropertiesOptions = driver.CollectionSchemaOptions

type ArangoDb[TDatabase DriverDatabase, TClient DriverClient] struct{}

func NewArangoDb() *ArangoDb[DriverDatabase, DriverClient] {
	return &ArangoDb[DriverDatabase, DriverClient]{}
}

func (ab *ArangoDb[TDatabase, TClient]) New(
	name string,
	address string,
	username string,
	password string,
) *database.DbHandler[DriverDatabase, DriverClient] {
	connection, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{address},
	})
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	arangoClient, err := driver.NewClient(driver.ClientConfig{
		Connection:     connection,
		Authentication: driver.BasicAuthentication(username, password),
	})
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	dbhandler := &database.DbHandler[DriverDatabase, DriverClient]{
		Client: arangoClient,
	}

	// NOTE: Only adds Database if a name id provided
	if name != "" {
		arangoDatabase, err := arangoClient.Database(context.Background(), name)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		dbhandler.Database = arangoDatabase
	}

	return dbhandler
}
