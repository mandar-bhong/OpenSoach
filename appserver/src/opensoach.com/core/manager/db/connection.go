package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func getConnectionEngine(connection string) (err error, engine *sqlx.DB) {

	value, exists := connectionDBEngine[connection]

	if exists == false {

		engine, err := sqlx.Connect("mysql", connection)

		if err != nil {
			return err, nil
		}

		connectionDBEngine[connection] = engine

		return nil, engine

	} else {
		return nil, value
	}

}
