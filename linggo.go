package linggo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type JFile string

func Set(file string) (jsonFile JFile, err error) {
	// add extension json to file name
	jsonFile = JFile(file + ".json")

	// Open the jsonFile
	_, err = os.Open(string(jsonFile))
	errorHandler(err)

	return jsonFile, err
}

func (f JFile) Ts(key string) (message interface{}) {
	// Open the jsonFile
	jsonFile, err := os.Open(string(f))
	errorHandler(err)

	// read the byte
	jsonByte, err := ioutil.ReadAll(jsonFile)
	errorHandler(err)

	// set the result interface
	var result map[string]interface{}

	// unmarshal the byte to interface
	err = json.Unmarshal(jsonByte, &result)

	// set the message result
	message = result[key]

	// return the result message
	return message
}

func Tr(file string, key string) (message interface{},  err error){
	// Open the jsonFile
	jsonFile, err := os.Open(file + ".json")
	errorHandler(err)

	// defer the closing of our jsonFile so that we can parse it later on
	defer func() {
		err = jsonFile.Close()
		errorHandler(err)
	}()

	// read the file
	message, err = readTheFile(jsonFile, key)

	// return the result message or error
	return message, err
}

// Read the file.
func readTheFile(jsonFile *os.File, key string)  (message interface{}, err error) {
	// read the byte
	jsonByte, err := ioutil.ReadAll(jsonFile)

	// set the result interface
	var result map[string]interface{}

	// unmarshal the byte to interface
	err = json.Unmarshal(jsonByte, &result)

	// set the message result
	message = result[key]

	// return the result message or error
	return message, err
}

// error handler
func errorHandler(err error) {
	// if error is not nil then print the error message and quit the application
	if err != nil { log.Fatal(err) }
}


