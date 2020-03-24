package mysql

import (
	"base-azec-library/com/azec/base/baseModel"
	"database/sql"
	"fmt"
	"github.com/ian-kent/go-log/log"
   _ "github.com/go-sql-driver/mysql"
)

func MysqlConnection(databaseModel baseModel.DatabaseModel) *sql.DB {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		databaseModel.User, databaseModel.Password, databaseModel.Host, databaseModel.Port, databaseModel.Dbname)
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Error(err.Error())
	}
	return db
}