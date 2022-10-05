package jsonx

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"os"
)

type JObj map[string]interface{}

func (j *JObj) GetStr(key string) string {
	if j == nil || *j == nil {
		return ""
	}
	if getter, hit := (*j)[key]; hit {
		return fmt.Sprint(getter)
	}
	return ""
}
func (j *JObj) GetBool(key string) bool {
	if j == nil || *j == nil {
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
	if j == nil || *j == nil {
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
	if j == nil || *j == nil {
		return nil
	}
	if getter, hit := (*j)[key]; hit {
		if got, hit := getter.([]interface{}); hit {
			return (*JArr)(&got)
		}
	}
	return nil
}
func (j *JObj) GetStruct(key string, ret interface{}) {
	if mapper := j.GetJObj(key); mapper != nil {
		mapstructure.Decode(mapper, ret)
	}
}
func (j *JObj) GetInt(key string) int64 {
	if j == nil || *j == nil {
		return 0
	}
	if getter, hit := (*j)[key]; hit {
		if got, hit := getter.(float64); hit {
			return int64(got)
		}
	}
	return 0
}
func (j *JObj) GetFloat(key string) float64 {
	if j == nil || *j == nil {
		return 0
	}
	if getter, hit := (*j)[key]; hit {
		if got, hit := getter.(float64); hit {
			return got
		}
	}
	return 0
}
func (j *JObj) GetString(key string) string {
	if j == nil || *j == nil {
		return ""
	}
	if getter, hit := (*j)[key]; hit {
		if got, hit := getter.(string); hit {
			return got
		}
	}
	return ""
}
func GetJObjFromInterface(src interface{}) (*JObj, bool) {
	if got, hit := (src).(map[string]interface{}); hit {
		ret := (*JObj)(&got)
		return ret, hit
	}
	return nil, false
}
func DecodeFromJson(infoRaw string) *JObj {
	ret := make(JObj)
	if err := json.Unmarshal([]byte(infoRaw), &ret); err != nil {
		return &ret
	}
	return &ret
}
func DecodeFromFile(path string) *JObj {
	ret := make(JObj)
	jsonRaw, err := os.ReadFile(path)
	if err != nil {
		return &ret
	}
	if err = json.Unmarshal(jsonRaw, &ret); err != nil {
		return &ret
	}
	return &ret

}
func DecodeFromMap(src map[string]interface{}) *JObj {
	return (*JObj)(&src)
}
func InterfaceToString(src interface{}) string {
	if got, hit := src.(string); hit {
		return got
	}
	return fmt.Sprintf("%v", src)
}
