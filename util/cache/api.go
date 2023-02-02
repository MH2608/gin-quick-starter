package cache

type SetStatusCode int8

const (
	Upward SetStatusCode = iota
	Insert
	Obsolete
)

// Setter set the cache map,and return a code for status of this action
type Setter func(key string, value interface{}) SetStatusCode
type Api interface {
	Get(key string, setter Setter) interface{}
	Set(key string, value interface{}) SetStatusCode
}

func NewCache(ObsoleteType string, MaxSize int64) Api {
	//todo full it
	panic("have not filled yet")
}
