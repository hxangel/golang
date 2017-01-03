package main

import (
	"fmt"
	"runtime"
	"time"
)

type Job struct {
	filename string
	results  chan <- Result
}

type Result struct {
	filename string
	line     string
	lino     int
}

var worker = runtime.NumCPU()

func main() {
	// config cpu number
	runtime.GOMAXPROCS(worker)

	files := []string{"f1", "f2", "f3"}
	// 任务列表, 并发数目为CPU个数
	jobs := make(chan Job, worker)
	// 结果
	results := make(chan Result, minimum(1000, len(files)))
	// 标记完成
	dones := make(chan struct{}, worker)

	go addJob(files, jobs, results)
	for i := 0; i < worker; i++ {
		go doJob(jobs, dones)
	}
	awaitForCloseResult(dones, results)

}

func addJob(files []string, jobs chan <- Job, results chan <- Result) {
	for _, filename := range files {
		jobs <- Job{filename, results}
	}
	close(jobs)
}

func doJob(jobs <-chan Job, dones chan <- struct{}) {
	for job := range jobs {
		job.Do()
	}
	dones <- struct{}{}
}

func awaitForCloseResult(dones <-chan struct{}, results chan Result) {
	working := worker
	done := false
	for {
		select {
		case result := <-results:
			println(result)
		case <-dones:
			working -= 1
			if working <= 0 {
				done = true
			}
		default:
			if done {
				return
			}
		}
	}
}

func (j *Job) Do() {
	fmt.Println("doing ", j.filename)
	time.Sleep(time.Second)
}

func minimum(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func println(o ...interface{}) {
	fmt.Println(o...)
}

