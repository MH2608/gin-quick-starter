package jsonx

type JsonX interface {
	ToString() string
}

func FilterStrArr(arrSrc []string, filter func(src string) bool) []string {
	ret := make([]string, 0)
	for _, str := range arrSrc {
		if filter(str) {
			ret = append(ret, str)
		}
	}
	return ret
}

func TernaryWithRet(condition bool, True func(paras ...interface{}) interface{}, False func(paras ...interface{}) interface{}, paras ...interface{}) interface{} {
	if condition {
		return True(paras)
	}
	return False(paras)
}
func EmptyTernary(condition bool, True func(), False func()) {
	if condition {
		True()
	} else {
		False()
	}
}
func Ternary(condition bool, True interface{}, False interface{}) interface{} {
	if condition {
		return True
	} else {
		return False
	}
}
