package business_tool

type BusinessReq interface {
	IsValid() bool
	DoService() (interface{}, error)
}
