package goplanout

type TypedMap struct{
	data	map[string]interface{}
}

func NewTypedMap(data map[string]interface{}) *TypedMap {
	return &TypedMap{data:data}
}

func (t *TypedMap) get(key string) (interface{}, bool) {
	value, exists := t.data[key]
	return value, exists
}

func (t *TypedMap) getString(key string) (string, bool) {
	value, exists := t.get(key)
	if exists {
		return value.(string), true
	}

	return "", false
}

func (t *TypedMap) getBool(key string) (bool, bool) {
	value, exists := t.get(key)
	if exists {
		return value.(bool), true
	}

	return false, false
}

func (t *TypedMap) getInt(key string) (int, bool) {
	value, exists := t.get(key)
	if exists {
		return value.(int), true
	}

	return 0, false
}

func (t *TypedMap) getInt64(key string) (int64, bool) {
	value, exists := t.get(key)
	if exists {
		return value.(int64), true
	}

	return 0, false
}

func (t *TypedMap) getFloat32(key string) (float32, bool) {
	value, exists := t.get(key)
	if exists {
		return value.(float32), true
	}

	return 0.0, false
}

func (t *TypedMap) getFloat64(key string) (float64, bool) {
	value, exists := t.get(key)
	if exists {
		return value.(float64), true
	}

	return 0.0, false
}

func (t *TypedMap) getMap(key string) (map[string]interface{}, bool) {
	value, exists := t.get(key)
	if exists {
		return value.(map[string]interface{}), true
	}

	return nil, false
}

func (t *TypedMap) getArray(key string) ([]interface{}, bool) {
	value, exists := t.get(key)
	if exists {
		return value.([]interface{}), true
	}

	return nil, false
}
