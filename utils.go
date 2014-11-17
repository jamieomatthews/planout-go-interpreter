package goplanout

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
)

func getOrElse(m map[string]interface{}, key string, def interface{}) interface{} {
	v, exists := m[key]
	if !exists {
		return def
	}
	return v
}

func assertKind(lhs, rhs interface{}, opstr string) {
	lhstype, rhstype := reflect.ValueOf(lhs), reflect.ValueOf(rhs)
	if lhstype.Kind() != rhstype.Kind() {
		panic(fmt.Sprintf("%v: Type mismatch between LHS %v and RHS %v\n", opstr, lhs, rhs))
	}
}

func compare(lhs, rhs interface{}) int {
	lval, rval := reflect.ValueOf(lhs), reflect.ValueOf(rhs)
	switch lval.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cmpInt(lval.Int(), rval.Int())
	case reflect.Float32, reflect.Float64:
		return cmpFloat(lval.Float(), rval.Float())
	case reflect.String:
		return cmpString(lval.String(), rval.String())
	}
	panic(fmt.Sprintln("Compare: Unsupported type"))
}

func isTrue(v interface{}) bool {
	lval := reflect.ValueOf(v)
	switch lval.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cmpInt(lval.Int(), 1) == 0
	case reflect.Float32, reflect.Float64:
		return cmpFloat(lval.Float(), 1.0) == 0
	case reflect.Bool:
		return lval.Bool()
	case reflect.String:
		return len(lval.String()) > 0
	}
	panic(fmt.Sprintln("IsTrue: Unsupported type"))
}

func isFalse(v interface{}) bool {
	lval := reflect.ValueOf(v)
	switch lval.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cmpInt(lval.Int(), 0) == 0
	case reflect.Float32, reflect.Float64:
		return cmpFloat(lval.Float(), 0.0) == 0
	case reflect.Bool:
		return !lval.Bool()
	case reflect.String:
		return len(lval.String()) <= 0
	}
	panic(fmt.Sprintln("IsFalse: Unsupported type"))
}

func cmpFloat(lhs, rhs float64) int {
	ret := 0
	if lhs == rhs {
		ret = 0
	} else if lhs < rhs {
		ret = -1
	} else {
		ret = 1
	}
	return ret
}

func cmpInt(lhs, rhs int64) int {
	ret := 0
	if lhs == rhs {
		ret = 0
	} else if lhs < rhs {
		ret = -1
	} else {
		ret = 1
	}
	return ret
}

func cmpString(lhs, rhs string) int {
	ret := 0
	if lhs == rhs {
		ret = 0
	} else if lhs < rhs {
		ret = -1
	} else {
		ret = 1
	}
	return ret
}

func add(x, y interface{}) interface{} {

	assertKind(x, y, "Addition")

	a, b := reflect.ValueOf(x), reflect.ValueOf(y)

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() + b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return a.Uint() + b.Uint()
	case reflect.Float32, reflect.Float64:
		return a.Float() + b.Float()
	case reflect.String:
		return a.String() + b.String()
	}
	panic("Addition: Unsupported type")
}

func addSlice(x []interface{}) interface{} {
	ret := x[0]
	for i := range x {
		if i != 0 {
			ret = add(ret, x[i])
		}
	}
	return ret
}

func multiply(x, y interface{}) interface{} {

	assertKind(x, y, "Multiplication")

	a, b := reflect.ValueOf(x), reflect.ValueOf(y)

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return a.Int() * b.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return a.Uint() * b.Uint()
	case reflect.Float32, reflect.Float64:
		return a.Float() * b.Float()
	}
	panic("Multiplication: Unsupported type")
}

func multiplySlice(x []interface{}) interface{} {
	ret := x[0]
	for i := range x {
		if i != 0 {
			ret = multiply(ret, x[i])
		}
	}
	return ret
}

func generateUnitStr(units interface{}) string {
	unitval := reflect.ValueOf(units)
	switch unitval.Kind() {
	case reflect.Array, reflect.Slice:
		v := units.([]interface{})
		n := len(v)
		var buffer bytes.Buffer
		buffer.WriteString(v[0].(string))
		for i := 0; i < n; i++ {
			if i != 0 {
				buffer.WriteString(".")
				buffer.WriteString(v[i].(string))
			}
		}
		return buffer.String()
	case reflect.String:
		return unitval.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(unitval.Int(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(unitval.Float(), 'f', -1, 64)
	}
	return ""
}

func getCummulativeWeights(weights []interface{}) (float64, []float64) {
	nweights := len(weights)
	cweights := make([]float64, nweights)
	sum := 0.0
	for i := range weights {
		sum = sum + weights[i].(float64)
		cweights[i] = sum
	}
	return sum, cweights
}

func generateString() string {
	s := make([]byte, 10)
	for j := 0; j < 10; j++ {
		s[j] = 'a' + byte(rand.Int()%26)
	}
	return string(s)
}

func toString(unit interface{}) string {
	unitval := reflect.ValueOf(unit)
	switch unitval.Kind() {
	case reflect.String:
		return unitval.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(unitval.Int(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(unitval.Float(), 'f', -1, 64)
	}
	return ""
}
