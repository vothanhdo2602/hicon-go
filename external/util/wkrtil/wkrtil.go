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

func (s *WorkerPool) Start() {
	defer func() {
		if r := recover(); r != nil {
			log.WithCtx(context.Background()).Error("Recover from panic", zap.Any("error", r))
			return
		}
	}()

	for i := 0; i < s.numWorkers; i++ {
		go func() {
			for {
				select {
				case job := <-s.jobs:
					job.Handler(job)
				}
			}
		}()
	}
}

func (s *WorkerPool) Submit(job *Job) {
	s.jobs <- job
}

func (s *WorkerPool) Stop() {
	close(s.jobs)
}

func GetWorkerPool() *WorkerPool {
	return wp
}

func (s *WorkerPool) SendJob(fn func(*Job), cb func(r chan *Result)) {
	if fn == nil {
		return
	}
	j := &Job{Handler: fn, ChResult: make(chan *Result, 1)}
	s.Submit(j)
	if cb != nil {
		cb(j.ChResult)
		close(j.ChResult)
	}
}
