package jsonx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

//通用json解析
//使用 interface{} 解析 json 内容，以便兼容各种复杂结构；
//interface{} 解析结果只有 string, float64（number）, bool 三种类型；
//此结构的方法仅限json内使用。
type JSON map[string]interface{}

type Array []JSON

//panic warning
func (m JSON) String(key string) string {
	return m[key].(string)
}

func (m JSON) SafeString(key string) (string, error) {
	if m[key] == nil {
		return "", fmt.Errorf("field %s not exist", key)
	}
	v, ok := m[key].(string)
	if ok {
		return v, nil
	}
	return "", fmt.Errorf("field '%s' is not string", key)
}

func (m JSON) MustString(key string) string {
	s, _ := m.SafeString(key)
	return s
}

//panic warning
func (m JSON) Int(key string) int {
	return int(m.Number(key))
}

func (m JSON) SafeInt(key string) (int, error) {
	i, err := m.SafeNumber(key)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func (m JSON) MustInt(key string) int {
	return int(m.MustNumber(key))
}

//panic warning
func (m JSON) Number(key string) float64 {
	return m[key].(float64)
}

func (m JSON) SafeNumber(key string) (float64, error) {
	if m[key] == nil {
		return 0, fmt.Errorf("field %s not exist", key)
	}
	v, ok := m[key].(float64)
	if ok {
		return v, nil
	}
	return 0, fmt.Errorf("field '%s' is not number", key)
}

func (m JSON) MustNumber(key string) float64 {
	f, _ := m.SafeNumber(key)
	return f
}

//panic warning
func (m JSON) Bool(key string) bool {
	return m[key].(bool)
}

func (m JSON) SafeBool(key string) (bool, error) {
	if m[key] == nil {
		return false, fmt.Errorf("field %s not exist", key)
	}
	v, ok := m[key].(bool)
	if ok {
		return v, nil
	}
	return false, fmt.Errorf("field '%s' is not bool", key)
}

func (m JSON) MustBool(key string) bool {
	b, _ := m.SafeBool(key)
	return b
}

//panic warning
func (m JSON) GetChild(key string) JSON {
	if m[key] == nil {
		return nil
	}
	return m[key].(map[string]interface{})
}

func (m JSON) SafeGetChild(key string) (JSON, error) {
	if m[key] == nil {
		return nil, fmt.Errorf("field %s not exist", key)
	}
	v, ok := m[key].(map[string]interface{})
	if ok {
		return v, nil
	}
	return nil, fmt.Errorf("field '%s' is not map[string]interface{}", key)
}

func (m JSON) Get(key string) interface{} {
	return m[key]
}

func DecodeJSON(b []byte) (interface{}, error) {
	b = bytes.TrimSpace(b)
	var v interface{}
	if b[0] == '[' {
		v = Array{}
	} else if b[0] == '{' {
		v = JSON{}
	} else {
		return nil, errors.New("not json")
	}
	err := json.Unmarshal(b, v)
	return v, err
}
