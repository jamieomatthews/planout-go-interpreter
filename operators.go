package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var ops map[string]bool

func init() {
	ops = map[string]bool{
		"set":            true,
		"seq":            true,
		"get":            true,
		"array":          true,
		"length":         true,
		"cond":           true,
		">":              true,
		">=":             true,
		"<":              true,
		"<=":             true,
		"equals":         true,
		"and":            true,
		"or":             true,
		"not":            true,
		"min":            true,
		"max":            true,
		"sum":            true,
		"product":        true,
		"%":              true,
		"/":              true,
		"uniformChoice":  true,
		"bernoulliTrial": true,
		"weightedChoice": true,
		"randomInteger":  true,
		"randomFloat":    true,
	}

	rand.Seed(time.Now().UTC().UnixNano())
}

type operator interface {
	execute(map[string]interface{}) interface{}
}

func isOperator(expr interface{}) (string, bool) {
	v := reflect.ValueOf(expr)
	if v.Kind() != reflect.Map {
		return "", false
	}

	js := expr.(map[string]interface{})
	opstr, exists := js["op"]
	if exists {
		_, exists = ops[opstr.(string)]
		if exists {
			return opstr.(string), true
		}
	}
	return "", false
}

func getOperator(opstr string, p map[string]interface{}) operator {
	var op operator
	switch {
	case opstr == "seq":
		op = &seq{p}
	case opstr == "set":
		op = &set{p}
	case opstr == "get":
		op = &get{p}
	case opstr == "array":
		op = &array{p}
	case opstr == "length":
		op = &length{p}
	case opstr == "cond":
		op = &cond{p}
	case opstr == "<":
		op = &lt{p}
	case opstr == "<=":
		op = &lte{p}
	case opstr == ">":
		op = &gt{p}
	case opstr == ">=":
		op = &gte{p}
	case opstr == "equals":
		op = &eq{p}
	case opstr == "and":
		op = &and{p}
	case opstr == "or":
		op = &or{p}
	case opstr == "not":
		op = &not{p}
	case opstr == "max":
		op = &max{p}
	case opstr == "min":
		op = &min{p}
	case opstr == "sum":
		op = &sum{p}
	case opstr == "product":
		op = &mul{p}
	case opstr == "%":
		op = &mod{p}
	case opstr == "/":
		op = &div{p}
	case opstr == "uniformChoice":
		op = &uniformChoice{p}
	case opstr == "bernoulliTrial":
		op = &bernoulliTrial{p}
	case opstr == "weightedChoice":
		op = &weightedChoice{p}
	case opstr == "randomFloat":
		op = &randomFloat{p}
	case opstr == "randomInteger":
		op = &randomInteger{p}
	}
	return op
}

type seq struct{ params map[string]interface{} }

func (s *seq) execute(m map[string]interface{}) interface{} {
	return evaluate(m["seq"], s.params)
}

type set struct{ params map[string]interface{} }

func (s *set) execute(m map[string]interface{}) interface{} {
	s.params["salt"] = m["var"].(string)
	value := evaluate(m["value"], s.params)
	s.params[m["var"].(string)] = value
	return true
}

type get struct{ params map[string]interface{} }

func (s *get) execute(m map[string]interface{}) interface{} {
	value, exists := s.params[m["var"].(string)]
	if !exists {
		panic(fmt.Sprintf("No input for key %v\n", m["var"]))
	}
	return value
}

type array struct{ params map[string]interface{} }

func (s *array) execute(m map[string]interface{}) interface{} {
	return evaluate(m["values"], s.params)
}

type length struct{ params map[string]interface{} }

func (s *length) execute(m map[string]interface{}) interface{} {
	values := evaluate(m["values"], s.params).([]interface{})
	l := make([]int, len(values))
	for i, value := range values {
		l[i] = len(value.([]interface{}))
	}
	return l[0]
}

type and struct{ params map[string]interface{} }

func (s *and) execute(m map[string]interface{}) interface{} {
	values := evaluate(m["values"], s.params).([]interface{})
	if len(values) == 0 {
		return false
	}

	for _, value := range values {
		if isFalse(value) {
			return false
		}
	}
	return true
}

type or struct{ params map[string]interface{} }

func (s *or) execute(m map[string]interface{}) interface{} {
	values := evaluate(m["values"], s.params).([]interface{})
	if len(values) == 0 {
		return false
	}

	for _, value := range values {
		if isTrue(value) {
			return true
		}
	}
	return false
}

type not struct{ params map[string]interface{} }

func (s *not) execute(m map[string]interface{}) interface{} {
	value := evaluate(m["value"], s.params)
	return !isTrue(value)
}

type cond struct{ params map[string]interface{} }

func (s *cond) execute(m map[string]interface{}) interface{} {
	conditions := m["cond"].([]interface{})
	for i := range conditions {
		c := conditions[i].(map[string]interface{})
		if evaluate(c["if"], s.params).(bool) == true {
			return evaluate(c["then"], s.params)
		}
	}
	return true
}

type lt struct{ params map[string]interface{} }

func (s *lt) execute(m map[string]interface{}) interface{} {
	lhs, rhs := evaluate(m["left"], s.params), evaluate(m["right"], s.params)
	assertKind(lhs, rhs, "LessThan")
	return compare(lhs, rhs) < 0
}

type lte struct{ params map[string]interface{} }

func (s *lte) execute(m map[string]interface{}) interface{} {
	lhs, rhs := evaluate(m["left"], s.params), evaluate(m["right"], s.params)
	assertKind(lhs, rhs, "LessThanEqual")
	return compare(lhs, rhs) <= 0
}

type gt struct{ params map[string]interface{} }

func (s *gt) execute(m map[string]interface{}) interface{} {
	lhs, rhs := evaluate(m["left"], s.params), evaluate(m["right"], s.params)
	assertKind(lhs, rhs, "GreaterThan")
	return compare(lhs, rhs) > 0
}

type gte struct{ params map[string]interface{} }

func (s *gte) execute(m map[string]interface{}) interface{} {
	lhs, rhs := evaluate(m["left"], s.params), evaluate(m["right"], s.params)
	assertKind(lhs, rhs, "GreaterThanEqual")
	return compare(lhs, rhs) >= 0
}

type eq struct{ params map[string]interface{} }

func (s *eq) execute(m map[string]interface{}) interface{} {
	lhs, rhs := evaluate(m["left"], s.params), evaluate(m["right"], s.params)
	assertKind(lhs, rhs, "Equal")
	return compare(lhs, rhs) == 0
}

type min struct{ params map[string]interface{} }

func (s *min) execute(m map[string]interface{}) interface{} {
	values := evaluate(m["values"], s.params).([]interface{})
	if len(values) == 0 {
		return false
	}
	return values[0]
}

type max struct{ params map[string]interface{} }

func (s *max) execute(m map[string]interface{}) interface{} {
	values := evaluate(m["values"], s.params).([]interface{})
	if len(values) == 0 {
		return false
	}
	return values[len(values)-1]
}

type sum struct{ params map[string]interface{} }

func (s *sum) execute(m map[string]interface{}) interface{} {
	values := evaluate(m["values"], s.params).([]interface{})
	return addSlice(values)
}

type mul struct{ params map[string]interface{} }

func (s *mul) execute(m map[string]interface{}) interface{} {
	values := evaluate(m["values"], s.params).([]interface{})
	return multiplySlice(values)
}

type mod struct{ params map[string]interface{} }

func (s *mod) execute(m map[string]interface{}) interface{} {
	var ret int64 = 0
	lhs := evaluate(m["left"], s.params).(float64)
	rhs := evaluate(m["right"], s.params).(float64)
	ret = int64(lhs) % int64(rhs)
	return float64(ret)
}

type div struct{ params map[string]interface{} }

func (s *div) execute(m map[string]interface{}) interface{} {
	var ret float64 = 0
	lhs := evaluate(m["left"], s.params).(float64)
	rhs := evaluate(m["right"], s.params).(float64)
	ret = lhs / rhs
	return ret
}
