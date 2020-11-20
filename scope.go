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

//Scope type to define wich kind of configuration we'd obtained.
//by default there are 2 defined Scopes (Global and Local)
type Scope string

const (
	//Global scope that define configuration as global
	Global Scope = "global"
	//Local scope that define configuration as local. By default
	//local configuration replaces global
	Local Scope = "local"
)
