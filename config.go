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

import "encoding/json"

//Configuration type to store all configuration itens
type Configuration map[string]*Item

//New create a new Configuration based on a json.
//
//	Example:
//		var jsonBody string = `{
//			"docs": [
//				{
//					"name": "test",
//					"type": "unit-test",
//					"data": {
//						"someValue": "test",
//						"anotherValue": 123,
//						"anyValue": true
//					}
//				},
//				{
//					"name": "test1",
//					"type": "unit-test",
//					"data": {
//						"someValue": "test1",
//						"anotherValue": 456,
//						"anyValue": false
//					}
//				}
//			]
//		}`
//
//		conf := New(jsonBody)
func New(jsonBody []byte) *Configuration {
	var d map[string]interface{}
	err := json.Unmarshal(jsonBody, &d)
	failOnError(err, "Fail on parser jsonBody")

	conf := make(Configuration)
	for _, i := range d["docs"].([]interface{}) {
		v := CreateItemFromInterface(i)
		conf.Add(*v)
	}

	return &conf
}

//Add function to add new Item to configuration. This module takes for
//itself the decision of which item will be used, the global or local.
//
//By default the local configuration ovewrite the global. Merge the
//global and local configuration is not possible.
//
//	Example:
//		conf := make(Configuration)
//		conf.add(item)
func (conf *Configuration) Add(item Item) {
	if item.Name != "" {
		(*conf)[item.Name] = &item
		return
	}

	cName := string(item.Type)
	(*conf)[cName] = &item
}

//Get function to obtain the Item based on it key
func (conf *Configuration) Get(cName string, cType Component) *Item {
	val := (*conf)[cName]

	if val == nil {
		val = (*conf)[string(cType)]
	}

	return val
}
