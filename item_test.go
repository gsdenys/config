// Copyright 2020 intorch.org. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const unitTestComponent Component = "unit-test"

type unitTestStruct struct {
	SomeValue    string
	AnotherValue int
	AnyValue     bool
}

func createInterface() map[string]interface{} {
	data := make(map[string]interface{})

	data["someValue"] = "value of field"
	data["anotherValue"] = 123
	data["anyValue"] = true

	return data
}

func TestItem_Decode(t *testing.T) {
	assert := assert.New(t)

	data := createInterface()
	item := &Item{
		Name: "test",
		Type: unitTestComponent,
		Data: data,
	}

	assert.NotNil(item, "The structure should not be nil")

	var unitTest unitTestStruct
	item.Decode(&unitTest)

	assert.NotNil(unitTest, "The unitTest variable should not be nil")
	assert.Equalf(unitTest.SomeValue, data["someValue"], "Expected '%s', but found '%s'", data["someValue"], unitTest.SomeValue)
}

func TestCreateItemFromJSON(t *testing.T) {
	assert := assert.New(t)

	json := `{
		"name": "test", 
		"scope": "local", 
		"type": "unit-test", 
		"data": {
			"someValue": "test", 
			"anotherValue": 123, 
			"anyValue": true
		}
	}`

	item := CreateItemFromJSON([]byte(json))

	assert.NotNil(item)
	assert.Equalf(item.Name, "test", "Expected 'test', but found %s", item.Name)
}

func TestCreateItemFromInterface(t *testing.T) {
	assert := assert.New(t)

	data := createInterface()
	dt := make(map[string]interface{})

	dt["name"] = "test"
	dt["scope"] = "local"
	dt["type"] = "unit-test"
	dt["data"] = data

	item := CreateItemFromInterface(dt)
	assert.NotNil(item, "The item should not be nil")
	assert.Equalf(item.Name, "test", "Expected 'test', but found %s.", item.Name)
	assert.Equalf(item.Type, unitTestComponent, "Expected '%s', but found '%s'.", unitTestComponent, item.Type)
}
