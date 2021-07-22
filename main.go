package main 

import (
	"log"
	//"./limiter"
	"github.com/ryo-imai-bit/go_worker_pool/pool"
	"github.com/ryo-imai-bit/go_worker_pool/work"
)

const WORKER_COUNT = 5
const JOB_COUNT = 100

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for i, job := range work.CreateJobs(JOB_COUNT) {
		collector.Work <-pool.Work{Job: job, ID: i}
	}
}	
