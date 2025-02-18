package wkrtil

import (
	"context"
	"github.com/vothanhdo2602/hicon/external/util/log"
	"go.uber.org/zap"
	"runtime"
	"sync"
)

type Job struct {
	ID       string
	Handler  func(*Job)
	ChResult chan *Result
}

type Result struct {
	JobID  string
	Output interface{}
	Error  error
}

var (
	wp *WorkerPool
)

type WorkerPool struct {
	numWorkers int
	jobs       chan *Job
	wg         sync.WaitGroup
	done       chan bool
}

func NewWorkerPool() *WorkerPool {
	wp = &WorkerPool{
		numWorkers: runtime.GOMAXPROCS(0),
		jobs:       make(chan *Job),
		done:       make(chan bool),
	}
	wp.Start()
	return wp
}

func (wp *WorkerPool) Start() {
	defer func() {
		if r := recover(); r != nil {
			log.WithCtx(context.Background()).Error("Recover from panic", zap.Any("error", r))
			return
		}
	}()

	for i := 0; i < wp.numWorkers; i++ {
		go func() {
			for {
				select {
				case job := <-wp.jobs:
					job.Handler(job)
				}
			}
		}()
	}
}

func (wp *WorkerPool) Submit(job *Job) {
	wp.jobs <- job
}

func (wp *WorkerPool) Stop() {
	close(wp.jobs)
}

func SendJob(fn func(*Job), cb func(r chan *Result)) {
	if fn == nil {
		return
	}
	j := &Job{Handler: fn, ChResult: make(chan *Result, 1)}
	wp.Submit(j)
	if cb != nil {
		cb(j.ChResult)
		close(j.ChResult)
	}
}
