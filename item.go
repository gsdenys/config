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

//Component definition of component. Each system component needs to define a
//constant of this type to access the remote configuration
type Component string

//Item data structure that define a remote configuration item. Each item can
//be stored at a separated doocument inside te cauchdb noSQL database
type Item struct {
	//This information has relationship  with the running containter,
	//where the container name is the same value of item Name.
	Name string
	//This is the Type of component. e.g. Health, Channel
	Type Component
	//The configuration parameters
	Data map[string]interface{}
}

//CreateItemFromJSON create new Item based on json body. Is very important to know that
//the json body needs match with the struct
//
// Example:
//
//		json := `{
//			"name": "test",
//			"type": "unit-test",
//			"data": {
//				"someValue": "test",
//				"anotherValue": 123,
//				"anyValue": true
//			}
//		}`
//
//		item = CreateItemFromJSON(json)
func CreateItemFromJSON(jsonBody []byte) *Item {
	var item Item

	err := json.Unmarshal(jsonBody, &item)
	failOnError(err, "Fail on decode jsonBody")

	return &item
}

//CreateItemFromInterface create a new Item base on a map[string]interface{} parameter.
//Is very important to know that the interface body needs match with the struct.
//
//	Example:
//		m := make(map[string]interface{})
//		//populate interface m with data
//		item := CreateItemFromInterface(m)
func CreateItemFromInterface(v interface{}) *Item {
	item := &Item{}

	jsonBody, err := json.Marshal(v)
	failOnError(err, "Fail on json encode.")

	err = json.Unmarshal(jsonBody, &item)
	failOnError(err, "Fail on Struct decode.")

	return item
}

//Decode function to decode a map[string]interface{} and store the Data
//at struct passed by parameter.
//
// Example:
//
//		type UnitTestStruct struct {
//			SomeValue    string
//			AnotherValue int
//			AnyValue     bool
//		}
//
//		item := New(jsonBody)
//		var st UnitTestStruct
//		item.Decode(&item)
//
//Remember to pass the struct object by reference!
func (item Item) Decode(v interface{}) {
	jsonBody, err := json.Marshal(item.Data)
	failOnError(err, "Fail on json encode.")

	err = json.Unmarshal(jsonBody, v)
	failOnError(err, "Fail on Struct decode.")
}
