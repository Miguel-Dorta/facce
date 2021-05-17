package types

import (
	"encoding/json"
	"strconv"
	"strings"
)

func DetectType(s string) interface{} {
	switch s {
	case "", "undefined", "null", "NULL":
		return nil
	case "true", "TRUE":
		return true
	case "false", "FALSE":
		return false
	}

	switch s[0] {
	case '{':
		return getObject(s)
	case '[':
		return getArray(s)
	case '"', '\'':
		return s[1:len(s)-1]
	}

	number, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return number
	}
	return s
}

func getObject(s string) map[string]interface{} {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic("error parsing JSON object: " + err.Error())
	}
	return m
}

func getArray(s string) []interface{} {
	s = s[1:len(s)-1]
	elementStr := strings.Split(s, ",")
	elements := make([]interface{}, len(elementStr))
	for i, str := range elementStr {
		str = strings.TrimSpace(str)
		elements[i] = DetectType(s)
	}
	return elements
}
