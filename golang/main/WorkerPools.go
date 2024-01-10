// Worker pools are great example of use of buffered channels .
// ex: find the sum of digits of numbers : for 500 random numbers ie 500 tasks
// you do this job with creating  worker pool, having 20-30 workers & giving them workers 500 tasks with division
// you may create channels for 10-20 buffer values
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	jobid int
	num   int
}

type result struct {
	job
	sum int
}

// create BUFFERED channel for providig numbers
var ch_jobs = make(chan job, 10)

// create channel foir writing results
var ch_results = make(chan result, 10)

func main() {
	startTime := time.Now()

	// Create 500 tasks
	noOfTasks := 500
	go AllocateTasks(noOfTasks, ch_jobs)

	//CreateWorkerPool
	noOfWorkers := 40
	go CreateWorkerPool(noOfWorkers)

	// finally result
	done := make(chan int)
	go Printresult(done)

	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("\ntotal time taken ", diff.Seconds(), "seconds")

}

func AllocateTasks(noOfTasks int, ch_jobs chan<- job) {
	j := job{}
	for x := 1; x <= noOfTasks; x++ {
		j = job{jobid: x, num: rand.Intn(9999)}
		ch_jobs <- j
	}
	close(ch_jobs)
}

func CreateWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for x := 1; x <= noOfWorkers; x++ {
		wg.Add(1)
		go Worker(x, &wg)

	}
	wg.Wait()
	close(ch_results)
}

// who does the main work
func Worker(x int, wg *sync.WaitGroup) {
	for job := range ch_jobs {
		fmt.Printf("\nThe worker is %d and job is %v and the number is %d", x, job.jobid, job.num)
		sum := CalculateSum(job.num)
		// write to result channel
		r := result{job, sum}
		ch_results <- r

	}
	wg.Done()

}

func CalculateSum(n int) int {
	time.Sleep(1 * time.Second)

	sum := 0
	for n > 0 {
		lg := n % 10
		sum = sum + lg
		n = n / 10
	}

	return sum
}

func Printresult(done chan<- int) {
	for r := range ch_results {
		fmt.Printf("\nthe RESULT Job id  %d number %d  sum %d", r.jobid, r.num, r.sum)
	}
	done <- 1
}
