package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id  int
	err error
}

func (wk *Worker) Work(workerChan chan *Worker) {
	defer func() { //捕获异常信息，防止panic直接退出
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				wk.err = err
			} else {
				wk.err = fmt.Errorf(`Panic happened with [%v] `, r)
			}
			workerChan <- wk
		}
	}()
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	panic("err")
}

type WorkerManager struct {
	workerChan chan *Worker
	nWorkers   int
}

func NewWorkerManager(nWorkers int) *WorkerManager {
	return &WorkerManager{nWorkers: nWorkers, workerChan: make(chan *Worker, nWorkers)}
}

func (wm *WorkerManager) StartWorkerPool() {
	for i := 0; i < wm.nWorkers; i++ { //开启工人
		wk := &Worker{id: i}
		go wk.Work(wm.workerChan)
	}
	wm.KeepLiveWorkers() //开启保活
}

func (wm *WorkerManager) KeepLiveWorkers() {
	for wk := range wm.workerChan {
		fmt.Printf("Worker %d Stopped with err : [%v] \n", wk.id, wk.err)
		wk.err = nil
		fmt.Printf("Start id = %d \n", wk.id)
		go wk.Work(wm.workerChan)
	}
}

func main() {
	wm := NewWorkerManager(10)
	wm.StartWorkerPool()
}
