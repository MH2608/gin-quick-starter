package bus

import (
	"sync"
	"time"
)

type Topic string
type Message interface{}
type HandlerFunc func(message Message) bool
type Handler struct {
	Async bool
	Pipe
	HandlerFunc
}
type asyncTask struct {
	Message
	HandlerFunc
}
type Pipe string
type PipeLine []Pipe
type AsyncFunc struct {
	sync.RWMutex
	Tasks []asyncTask
}

func (async *AsyncFunc) addFunc(message Message, handlerFunc HandlerFunc) {
	async.Lock()
	defer func() {
		async.Unlock()
	}()
	async.Tasks = append(async.Tasks, asyncTask{
		Message:     message,
		HandlerFunc: handlerFunc,
	})
}

func (async *AsyncFunc) runTasks() {
	async.RLocker()
	defer func() {
		async.RUnlock()
	}()
	for _, task := range async.Tasks {
		task.HandlerFunc(task.Message)
	}
	async.Tasks = async.Tasks[:0]
}

// NewBus
// return a Bus ptr
// if you want to make sure the order of handler calling,call NewBus with a pipeline sorted,please
// /*
func NewBus(pipes ...Pipe) *Bus {
	if len(pipes) == 0 {
		return &Bus{
			topicWithHandlers: make(map[Topic][]*Handler),
			AsyncFunc: AsyncFunc{
				Tasks: make([]asyncTask, 0),
			},
		}
	}
	return &Bus{
		topicWithHandlers: make(map[Topic][]*Handler),
		pip:               pipes,
		AsyncFunc: AsyncFunc{
			Tasks: make([]asyncTask, 0),
		},
	}
}

type Bus struct {
	topicWithHandlers map[Topic][]*Handler
	pip               PipeLine
	AsyncFunc
}

func (b *Bus) RegisterTopic(topic Topic) {
	if _, hit := b.topicWithHandlers[topic]; hit {
		return
	}
	b.topicWithHandlers[topic] = make([]*Handler, 0)
}

func (b *Bus) RegisterHandler(topic Topic, handler *Handler) {
	b.RegisterTopic(topic)
	b.topicWithHandlers[topic] = append(b.topicWithHandlers[topic], handler)
}
func (b *Bus) produceWithPipe(topic Topic, message Message) {
	for _, pipe := range b.pip {
		for _, handler := range b.topicWithHandlers[topic] {
			if pipe == handler.Pipe {
				if handler.Async {
					b.AsyncFunc.addFunc(message, handler.HandlerFunc)
				} else if !handler.HandlerFunc(message) {
					return
				}
			}
		}
	}
}

func (b *Bus) produceWithoutPipe(topic Topic, message Message) {
	for _, handler := range b.topicWithHandlers[topic] {
		if handler.Async {
			b.AsyncFunc.addFunc(message, handler.HandlerFunc)
		} else if !handler.HandlerFunc(message) {
			return
		}
	}
}

// Produce
// Product message to eventbus,if there is a pipeline in the bus,it would run with handlers order by pipelines
// /*
func (b *Bus) Produce(topic Topic, message Message) {
	if b.pip == nil {
		b.produceWithoutPipe(topic, message)
	} else {
		b.produceWithPipe(topic, message)
	}
}

func (b *Bus) RunAsync(Intervals int64) {
	clock := time.NewTicker(time.Duration(Intervals))
	for {
		select {
		case <-clock.C:
			b.runTasks()
		}
	}
}
