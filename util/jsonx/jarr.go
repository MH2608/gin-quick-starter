package jsonx

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type JArr []interface{}

func (j *JArr) ToString() string {
	if !j.isValid() {
		return ""
	}
	return fmt.Sprint(j)
}
func (j *JArr) GetObj(index int) *JObj {
	if !j.isValid() || index >= len(*j) {
		return nil
	}
	getter := (*j)[index]
	if got, hit := getter.(map[string]interface{}); hit {
		return (*JObj)(&got)
	}
	return nil
}
func (j *JArr) GetInt(index int) int64 {
	if !j.isValid() || index >= len(*j) {
		return 0
	}
	getter := (*j)[index]
	switch got := getter.(type) {
	case float64:
		return int64(got)
	case float32:
		return int64(got)
	case int:
		return int64(got)
	case int64:
		return got
	case int32:
		return int64(got)
	case int16:
		return int64(got)
	case int8:
		return int64(got)
	default:
		return 0
	}
	return 0
}
func (j *JArr) GetFloat(index int) float64 {
	if !j.isValid() || index >= len(*j) {
		return 0
	}
	getter := (*j)[index]
	switch got := getter.(type) {
	case float64:
		return got
	case float32:
		return float64(got)
	case int:
		return float64(got)
	case int64:
		return float64(got)
	case int32:
		return float64(got)
	case int16:
		return float64(got)
	case int8:
		return float64(got)
	default:
		return 0
	}
	return 0
}
func (j *JArr) GetString(index int) string {
	if !j.isValid() || index >= len(*j) {
		return ""
	}
	getter := (*j)[index]
	if got, hit := getter.(string); hit {
		return got
	}
	return ""
}
func (j *JArr) GetArr(index int) *JArr {
	if !j.isValid() || index >= len(*j) {
		return nil
	}
	getter := (*j)[index]
	if got, hit := getter.([]interface{}); hit {
		return (*JArr)(&got)
	}
	return nil
}
func (j *JArr) GetStruct(index int, ret interface{}) {
	mapper := j.GetObj(index)
	if mapper != nil {
		mapstructure.Decode(mapper, ret)
	}
}
func (j *JArr) Foreach(looper func(index int, value interface{}) bool) bool {
	if !j.isValid() {
		return false
	}
	for i, v := range *j {
		if !looper(i, v) {
			return false
		}
	}
	return true
}
func (j *JArr) isValid() bool {
	if j == nil || *j == nil {
		return false
	}
	return true
}
func (j *JArr) ForeachObj(objLooper func(obj *JObj) bool) bool { //false会透传，但true不会
	if !j.isValid() {
		return false
	}
	for i, _ := range *j {
		jObj := j.GetObj(i)
		if jObj.isValid() {
			if !objLooper(jObj) {
				return false
			}
		}
	}
	return true
}
