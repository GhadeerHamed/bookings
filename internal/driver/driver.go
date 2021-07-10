package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

//DB holds the database connection
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

//ConnectSQL creates database pol for postgres
func ConnectSQL(dsn string) (*DB, error) {
	database, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	database.SetMaxOpenConns(maxOpenDbConn)
	database.SetMaxIdleConns(maxIdleDbConn)
	database.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = database
	err = testDB(database)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

//testDB tries to pind the database
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

//NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
