package oracle

import (
	"base-azec-library/com/azec/base/baseModel"
	"database/sql"
	"github.com/ian-kent/go-log/log"
	"strconv"
	_ "github.com/godror/godror"
)

func GoracleConnection(databaseModel baseModel.DatabaseModel) *sql.DB {

	db, err := sql.Open("godror", databaseModel.User+"/"+databaseModel.Password+"@"+databaseModel.Host+":"+strconv.FormatInt(int64(databaseModel.Port), 10)+"/"+databaseModel.Dbname)
	if err != nil {
		log.Error(err.Error())
	}

	return db
}