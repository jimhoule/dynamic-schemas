package database

type DbHandler[TDatabase any, TClient any] struct {
	Database TDatabase
	Client   TClient
}

type Db[TDatabase any, TClient any] interface {
	New(name string, address string, username string, password string) *DbHandler[TDatabase, TClient]
}

func NewDbHandler[TDatabase any, TClient any](db Db[TDatabase, TClient], name string, address string, username string, password string) *DbHandler[TDatabase, TClient] {
	return db.New(name, address, username, password)
}
