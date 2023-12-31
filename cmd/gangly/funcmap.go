// Copyright © 2018 Heptio
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

package main

import (
	"encoding/base64"
	"html/template"
)

var genericMap = map[string]interface{}{
	"base64enc": base64encode,
}

func FuncMap() template.FuncMap {
	return template.FuncMap(genericMap)
}

func base64encode(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}
