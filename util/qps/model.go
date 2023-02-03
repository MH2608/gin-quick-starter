package qps

import (
	"io"
	"sync"
	"time"
)

const dayMS = 86400

type DRecorder struct {
	date string
	qps  [dayMS]int64
}

type DataWriter struct {
	w      io.Writer
	actSig chan int
}

func (d *DataWriter) asyncWrite() {
	d.actSig <- 1
	defer func() { <-d.actSig }()
	//todo write to file && clickhouse
}

type Recorder struct {
	writing    *DRecorder
	memorizing *DRecorder
	writer     *DataWriter
	sync.Mutex
}

func (r *Recorder) exchange() {
	r.TryLock()
	temp := r.writing
	r.writing = r.memorizing
	r.writing.date = time.Now().Add(-time.Hour).Format("2006-01-02")
	r.memorizing = temp
	r.Unlock()
	go r.writer.asyncWrite()

}
