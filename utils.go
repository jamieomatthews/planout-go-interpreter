package main

import (
	"fmt"
	"reflect"
)

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
