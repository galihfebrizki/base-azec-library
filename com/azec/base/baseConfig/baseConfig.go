package baseConfig

import (
	"base-azec-library/com/azec/base/baseConfig/mysql"
	"base-azec-library/com/azec/base/baseModel"
	"database/sql"
	"encoding/json"
	"log"
	"io/ioutil"
	"os"
)

func MysqlConnection(databaseModel baseModel.DatabaseModel) *sql.DB {
	return mysql.MysqlConnection(databaseModel)
}

func ReadDbModel(pathConfFile string) baseModel.DatabaseModel {
	var databaseModel baseModel.DatabaseModel

	jsonFile, err := os.Open(pathConfFile)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err.Error())
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &databaseModel)

	return databaseModel
}
