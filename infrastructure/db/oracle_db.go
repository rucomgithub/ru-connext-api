package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/godror/godror"
)

type OracleDB struct {
	db *sqlx.DB
}

func NewOracleDB(connectionString string) (*OracleDB, error) {
	db, err := sqlx.Connect("godror", connectionString)
	if err != nil {
		return nil, err
	}
	
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	return &OracleDB{db: db}, nil
}

func (o *OracleDB) Close() error {
	return o.db.Close()
}

func (o *OracleDB) GetDB() *sqlx.DB {
	return o.db
}