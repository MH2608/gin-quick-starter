package qps

type API interface {
	WriteRecord(time int64, qps int64)
}
