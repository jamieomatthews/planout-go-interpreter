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
