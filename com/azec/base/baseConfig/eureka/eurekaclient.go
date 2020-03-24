package eureka

import (
	"base-azec-library/com/azec/base/baseModel"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func Register(pathFile string,urlReg string) {
	//dir, _ := os.Getwd()
	data, _ := ioutil.ReadFile(pathFile)

	tpl := string(data)

	// Register.
	registerAction := baseModel.HttpAction {
		Url : urlReg,
		Method: "POST",
		ContentType: "application/json",
		Body: tpl,
	}
	var result bool
	for {
		result = DoHttpRequest(registerAction)
		if result {
			break
		} else {
			time.Sleep(time.Second * 5)
		}
	}
}


func StartHeartbeat(urlHearbeat string) {
	for {
		time.Sleep(time.Second * 5)
		heartbeat(urlHearbeat)
	}
}

func heartbeat(urlHearbeat string) {
	heartbeatAction := baseModel.HttpAction{
		Url : urlHearbeat,
		Method: "PUT",
		ContentType: "application/json",
		Body: "",
	}
	DoHttpRequest(heartbeatAction)
}

func Deregister(urlDeregister string) {
	log.Println("Trying to deregister application...")
	// Deregister
	deregisterAction := baseModel.HttpAction {
		Url : urlDeregister,
		Method: "DELETE",
		ContentType: "application/json",
		Body: "",
	}
	DoHttpRequest(deregisterAction)
	log.Println("Deregistered application, exiting. Check Eureka...")
}

func ReadStatusEureka(pathStatusEureka string)baseModel.EurekaConfig  {
	var databaseModel baseModel.EurekaConfig

	jsonFile, err := os.Open(pathStatusEureka)

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