package baseConfig

import (
	"base-azec-library/com/azec/base/baseModel"
	"encoding/json"
	"github.com/ian-kent/go-log/log"
	"io/ioutil"
	"os"
)

func ReadDatabaseModel(pathConfFile string) baseModel.DatabaseModel {
	var databaseModel baseModel.DatabaseModel

	jsonFile, err := os.Open(pathConfFile)

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Error(err.Error())
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &databaseModel)

	return databaseModel
}
