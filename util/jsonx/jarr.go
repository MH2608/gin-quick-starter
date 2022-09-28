package jsonx

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type JArr []interface{}

func (j *JArr) ToString() string {
	if j == nil || *j == nil {
		return ""
	}
	return fmt.Sprint(j)
}
func (j *JArr) GetObj(index int) *JObj {
	if j == nil || *j == nil || index >= len(*j) {
		return nil
	}
	getter := (*j)[index]
	if got, hit := getter.(map[string]interface{}); hit {
		return (*JObj)(&got)
	}
	return nil
}
func (j *JArr) GetInt(index int) int64 {
	if j == nil || *j == nil || index >= len(*j) {
		return 0
	}
	getter := (*j)[index]
	if got, hit := getter.(float64); hit {
		return int64(got)
	}
	return 0
}
func (j *JArr) GetFloat(index int) float64 {
	if j == nil || *j == nil || index >= len(*j) {
		return 0
	}
	getter := (*j)[index]
	if got, hit := getter.(float64); hit {
		return got
	}
	return 0
}
func (j *JArr) GetString(index int) string {
	if j == nil || *j == nil || index >= len(*j) {
		return ""
	}
	getter := (*j)[index]
	if got, hit := getter.(string); hit {
		return got
	}
	return ""
}
func (j *JArr) GetArr(index int) *JArr {
	if j == nil || *j == nil || index >= len(*j) {
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
