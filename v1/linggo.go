package v1

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func Tr(file string, key string) (message interface{},  err error){
	// Open our jsonFile
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


