package reqCommon

type BaseReq interface {
	IsValid() bool
	DoService() (interface{}, error)
}
