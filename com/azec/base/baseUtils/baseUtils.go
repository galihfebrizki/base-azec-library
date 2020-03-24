package baseUtils

import (
	"io/ioutil"
	"os"
	"time"
	"base-azec-library/com/azec/base/baseConstanta"
)

func CreateDirectoryBaseOnDate(basePath string) (string, int64) {
	var OutMessage string
	var OutError int64

	currentTime := time.Now()
	err := os.Mkdir(basePath+currentTime.Format(baseConstanta.YYYYMMDD), 0777)
	if err != nil {
		OutMessage = err.Error()
		OutError = baseConstanta.DefaultOutErrorFailed
	} else {
		OutMessage = baseConstanta.DefaultOutMessageSucces
		OutError = baseConstanta.DefaultOutErrorSuccess
	}
	return OutMessage, OutError
}

func CreateDirectory(basePath string) (string, int64) {
	var OutMessage string
	var OutError int64

	err := os.Mkdir(basePath, 0777)
	if err != nil {
		OutMessage = err.Error()
		OutError = baseConstanta.DefaultOutErrorFailed
	} else {
		OutMessage = baseConstanta.DefaultOutMessageSucces
		OutError = baseConstanta.DefaultOutErrorSuccess
	}
	return OutMessage, OutError
}

func CheckDirectoryOrFileExists(path string) (string, int64) {
	var OutMessage string
	var OutError int64

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		OutMessage = baseConstanta.DefaultNotExistMessage
		OutError = baseConstanta.DefaultOutErrorSuccess
	} else {
		OutMessage = baseConstanta.DefaultExistsMessage
		OutError = baseConstanta.DefaultOutErrorFailed
	}
	return OutMessage, OutError
}

func CreateFile(path string, file []byte) (string, int64) {
	var OutMessage string
	var OutError int64

	err := ioutil.WriteFile(path, file, 0777)
	if err != nil {
		OutMessage = err.Error()
		OutError = baseConstanta.DefaultOutErrorFailed
	} else {
		OutMessage = baseConstanta.DefaultOutMessageSucces
		OutError = baseConstanta.DefaultOutErrorSuccess		
	}
	return OutMessage, OutError
}