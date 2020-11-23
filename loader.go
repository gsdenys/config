package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//FromFile function to load configurantion from JSON file
func FromFile(path string) *Configuration {
	file, err := os.Open(path)
	failOnError(err, "Error opening file.")

	defer file.Close()

	data, errR := ioutil.ReadAll(file)
	failOnError(errR, "Error reading file.")

	return New(data)
}

func getQuery(taskName string, component Component) []byte {
	queryBase := `{"selector":{"$or": [{"name": "%s"}, {"type": "%s"}]}}`
	query := fmt.Sprintf(queryBase, taskName, string(component))

	return []byte(query)
}

//FromCouchDB function to load configurantion from couchdb database
func FromCouchDB(dbURL string, taskName string, component Component) *Configuration {
	query := getQuery(taskName, component)

	dbURL = fmt.Sprintf("%s/_find", dbURL)

	res, err := http.Post(dbURL, "application/json", bytes.NewBuffer(query))
	failOnError(err, "Error querying data")

	body, errR := ioutil.ReadAll(res.Body)
	failOnError(errR, "Error reading body")

	return New(body)
}
