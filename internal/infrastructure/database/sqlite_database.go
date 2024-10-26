package database

import (
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ppp3ppj/track-my-day-gotth/internal/config"
)

type sqliteDatabase struct {
	*sqlx.DB
}

var (
	sqliteDatabaseInstance *sqliteDatabase
	once                   sync.Once
)

func NewSqliteDatabase(conf *config.Database) *sqliteDatabase {
	once.Do(func() {
		currentDBPath := conf.CurrentDBPath

		conn, err := sqlx.Connect("sqlite3", currentDBPath)
		if err != nil {
			panic(err)
		}

		log.Print("Connected to postgres database %s successfully", currentDBPath)

		sqliteDatabaseInstance = &sqliteDatabase{conn}
	})

	return sqliteDatabaseInstance
}

func (db *sqliteDatabase) Connect() *sqlx.DB {
	return sqliteDatabaseInstance.DB
}

func (db *sqliteDatabase) Close() error {
	if sqliteDatabaseInstance != nil && sqliteDatabaseInstance.DB != nil {
		if err := sqliteDatabaseInstance.DB.Close(); err != nil {
			return err
		}

		log.Println("Closed sqlite database connection")
	}
	return nil
}
