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

var jsonBody string = `{
	"docs": [
		{
			"name": "test", 
			"type": "unit-test", 
			"data": {
				"someValue": "test", 
				"anotherValue": 123, 
				"anyValue": true
			}
		},
		{
			"name": "", 
			"type": "unit-test", 
			"data": {
				"someValue": "test1", 
				"anotherValue": 456, 
				"anyValue": false
			}
		}
	]
}`

var jsonBodyGlobal string = `{
	"docs": [
		{
			"name": "", 
			"type": "unit-test", 
			"data": {
				"someValue": "test1", 
				"anotherValue": 456, 
				"anyValue": false
			}
		}
	]
}`

func TestNew(t *testing.T) {
	assert := assert.New(t)

	config := New([]byte(jsonBody))

	assert.NotNil(config)
	assert.NotEmpty(config)
	assert.Len(*config, 2)

	assert.Equal(config.Get("test", unitTestComponent).Name, "test")
	assert.Equal(config.Get("test", unitTestComponent).Type, unitTestComponent)
}

func TestNewGlobal(t *testing.T) {
	assert := assert.New(t)

	config := New([]byte(jsonBodyGlobal))

	assert.NotNil(config)
	assert.NotEmpty(config)
	assert.Len(*config, 1)

	assert.Equal(config.Get("test", unitTestComponent).Name, "")
	assert.Equal(config.Get("test", unitTestComponent).Type, unitTestComponent)
}

func TestConfiguration_Add(t *testing.T) {
	assert := assert.New(t)

	jsn := `{
		"name": "test", 
		"type": "unit-test", 
		"data": {
			"someValue": "test", 
			"anotherValue": 123, 
			"anyValue": true
		}
	}`
	item := CreateItemFromJSON([]byte(jsn))

	config := New([]byte(jsonBodyGlobal))
	i := config.Get("test", unitTestComponent)

	assert.Len(*config, 1)
	assert.Equal(i.Name, "")
	assert.Equal(i.Type, unitTestComponent)

	config.Add(*item)
	i = config.Get("test", unitTestComponent)

	assert.Len(*config, 2)
	assert.Equal(i.Name, "test")
	assert.Equal(i.Type, unitTestComponent)
}
