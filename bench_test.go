package main 

import "testing"
import "github.com/ryo-imai-bit/go_worker_pool/pool"
import  work "github.com/ryo-imai-bit/go_worker_pool/work"

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(WORKER_COUNT) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range work.CreateJobs(20) {
			collector.Work <-pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, job := range work.CreateJobs(20) {
			work.DoWork(job, 1)
		}
	}
}
