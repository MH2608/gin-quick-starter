package jsonx

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type JObj map[string]interface{}

func (j *JObj) GetStr(key string) string {
	if !j.isValid() {
		return ""
	}
	if getter, hit := (*j)[key]; hit {
		return fmt.Sprint(getter)
	}
	return ""
}
func (j *JObj) GetBool(key string) bool {
	if !j.isValid() {
		return false
	}
	if getter, hit := (*j)[key]; hit {
		if got, hit := getter.(bool); hit {
			return got
		}
	}
	return false
}
func (j *JObj) GetJObj(key string) *JObj {
	if !j.isValid() {
		return nil
	}
	if getter, hit := (*j)[key]; hit {
		if got, hit := getter.(map[string]interface{}); hit {
			return (*JObj)(&got)
		}

	}
	return nil
}
func (j *JObj) GetJArr(key string) *JArr {
	if !j.isValid() {
		return nil
	}
	if getter, hit := (*j)[key]; hit {
		if got, hit := getter.([]interface{}); hit {
			return (*JArr)(&got)
		}
	}
	return nil
}
func (j *JObj) GetStrArr(key string) []string {
	ret := make([]string, 0)
	src := j.GetJArr(key)
	if j.isValid() {
		src.Foreach(func(index int, value interface{}) bool {
			ret = append(ret, fmt.Sprint(value))
			return true
		})
	}
	return ret
}
func (j *JObj) GetStruct(key string, ret interface{}) {
	if mapper := j.GetJObj(key); mapper != nil {
		mapstructure.Decode(mapper, ret)
	}
}
func (j *JObj) GetInt(key string) int64 {
	if !j.isValid() {
		return 0
	}
	if getter, hit := (*j)[key]; hit {
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
	}
	return 0
}
func (j *JObj) GetFloat(key string) float64 {
	if !j.isValid() {
		return 0
	}
	if getter, hit := (*j)[key]; hit {
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
	}
	return 0
}
func (j *JObj) GetString(key string) string {
	if !j.isValid() {
		return ""
	}
	if getter, hit := (*j)[key]; hit {
		if got, hit := getter.(string); hit {
			return got
		}
	}
	return ""
}
func DecodeFromJson(infoRaw string) *JObj {
	ret := make(JObj)
	if err := json.Unmarshal([]byte(infoRaw), &ret); err != nil {
		return &ret
	}
	return &ret
}
func DecodeFromMap(src map[string]interface{}) JObj {
	return src
}
func (j *JObj) Foreach(looper func(key string, value interface{})) {
	if !j.isValid() {
		return
	}
	for k, v := range *j {
		looper(k, v)
	}
}
func (j *JObj) isValid() bool {
	if j == nil || *j == nil {
		return false
	}
	return true
}
func (j *JObj) CheckRequire(requireField ...string) bool {
	if !j.isValid() {
		return false
	}
	for _, field := range requireField {
		if _, ok := (*j)[field]; !ok {
			return false
		}
	}
	return true
}
