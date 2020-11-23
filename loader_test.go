package config

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const couchDB string = "http://admin:admin@localhost:5984/test"

func TestFromFile(t *testing.T) {
	assert := assert.New(t)

	config := FromFile("./loader_test.json")

	assert.NotNil(config)
	assert.NotEmpty(config)
	assert.Len(*config, 2)

	assert.Equal(config.Get("test", unitTestComponent).Name, "test")
	assert.Equal(config.Get("test", unitTestComponent).Type, unitTestComponent)
}

func createOrUpdateDB() {
	client := &http.Client{}

	reqD, _ := http.NewRequest(http.MethodDelete, couchDB, nil)
	client.Do(reqD)

	reqC, _ := http.NewRequest(http.MethodPut, couchDB, nil)
	client.Do(reqC)
}

func createEntry() {
	first := bytes.NewBuffer([]byte(`{
		"name": "test", 
		"type": "unit-test", 
		"data": {
			"someValue": "test", 
			"anotherValue": 123, 
			"anyValue": true
		}
	}`))

	second := bytes.NewBuffer([]byte(`{
		"name": "", 
		"type": "unit-test", 
		"data": {
			"someValue": "test1", 
			"anotherValue": 456, 
			"anyValue": false
		}
	}`))

	client := &http.Client{}

	reqF, _ := http.NewRequest(http.MethodPost, couchDB, first)
	reqF.Header.Set("Content-Type", "application/json")
	client.Do(reqF)

	reqS, _ := http.NewRequest(http.MethodPost, couchDB, second)
	reqS.Header.Set("Content-Type", "application/json")
	client.Do(reqS)

}

func TestFromCouchDB(t *testing.T) {
	createOrUpdateDB()
	createEntry()

	assert := assert.New(t)

	config := FromCouchDB(couchDB, "test", "unit-test")

	assert.NotNil(config)
	assert.NotEmpty(config)
	assert.Len(*config, 2)

	assert.Equal(config.Get("test", unitTestComponent).Name, "test")
	assert.Equal(config.Get("test", unitTestComponent).Type, unitTestComponent)
}
