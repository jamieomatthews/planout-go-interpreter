/*
 * Copyright 2014 URX
 * 
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * 
 *     http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package goplanout

import (
	"fmt"
	"reflect"
)

func evaluate(code interface{}, params map[string]interface{}) interface{} {
	v := reflect.ValueOf(code)

	switch v.Kind() {
	case reflect.Map:
		js := code.(map[string]interface{})
		opstr, exists := isOperator(js)
		if exists {
			e := getOperator(opstr, params)
			return e.execute(js)
		} else {
			fmt.Printf("Not an operator: \n%v\n", js)
		}
	case reflect.Array, reflect.Slice:
		v := make([]interface{}, len(code.([]interface{})))
		for i, js := range code.([]interface{}) {
			v[i] = evaluate(js, params)
		}
		return v
	}
	return code
}

func Experiment(code interface{}, params map[string]interface{}) bool {

	defer func() bool {
		if r := recover(); r != nil {
			fmt.Println("Recovered ", r)
			return false
		}
		return true
	}()

	evaluate(code, params)

	return true
}
