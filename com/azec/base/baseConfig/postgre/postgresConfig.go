package postgre

import (
	"base-azec-library/com/azec/base/baseModel"
	"database/sql"
	"fmt"
	"github.com/ian-kent/go-log/log"
)

func PostgresConnection(databaseModel baseModel.DatabaseModel) *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		databaseModel.Host, databaseModel.Port, databaseModel.User, databaseModel.Password, databaseModel.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(err.Error())
	}
	return db
}