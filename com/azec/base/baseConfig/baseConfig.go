package baseConfig

import (
	"base-azec-library/com/azec/base/baseConfig/mysql"
	"base-azec-library/com/azec/base/baseModel"
	"database/sql"
)

func MysqlConnection(databaseModel baseModel.DatabaseModel) *sql.DB {
	return mysql.MysqlConnection(databaseModel)
}