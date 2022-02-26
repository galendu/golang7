// 实现json序列化/反序列化对map的支持
package main

/*
func convertBasic(val interface{}) string {
	switch v := val.(type) {
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.Itoa(int(v))
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.Itoa(int(v))
	case uint:
		return strconv.Itoa(int(v))
	case uint8:
		return strconv.Itoa(int(v))
	case uint16:
		return strconv.Itoa(int(v))
	case uint32:
		return strconv.Itoa(int(v))
	case uint64:
		return strconv.Itoa(int(v))
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	case string:
		return "\"" + string(v) + "\""
	default:
		return "wrong"
	}
}

func convertMap(keyType, elemType reflect.Type, value map[interface{}]interface{}) string {
	sb := strings.Builder{}
	sb.WriteString("{")
	for key, elem := range value {
		keyStr := convertBasic(key)
		sb.WriteString("\"" + keyStr + "\":" + convertBasic(elem))
	}
	sb.WriteString("}")
	fmt.Println(sb.String())
	return sb.String()
}

func Marshal(value interface{}) string {
	fmt.Println(value)
	valueType := reflect.TypeOf(value)
	valueTypeKind := valueType.Kind()
	fmt.Printf("typeof: %s, kind: %s\n", valueType, valueTypeKind)
	switch valueTypeKind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			 reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			 reflect.Float32, reflect.Float64,
			 reflect.Bool,
			 reflect.String:
		return convertBasic(value)
	case reflect.Map:
		keyType := valueType.Key()
		elemType := valueType.Elem()
		fmt.Printf("case: map, key: %s, elem: %s \n", keyType.Kind(), elemType.Kind())
		m := value.(map[interface{}]interface{})
		return convertMap(keyType, elemType, m)
	}
	return ""
}

func UnMarshal(str string) interface{} {
	fmt.Println("unmarshal")
	return nil
}
*/
func main() {
	// var basic int8 = 1
	m := make(map[string]string, 8)
	m["a"] = "a"
	// m[2] = "b"
	// m[3] = "c"
	// var ui uint32 = 244
	// result := Marshal(m)
	// fmt.Println("result: ", result)
}