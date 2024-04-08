package main

import (
	"context"
	"fmt"
	"math/rand"
	_ "math/rand"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	RateLimit()
}
func RateLimit() {
	const maxJobsPerSecond = 5
	const numJobs = 22
	var wg sync.WaitGroup

	// 计数器
	var runningJobs int32  // 当前正在执行的任务数量
	var startedJobs int32  // 启动后的任务数量
	var finishedJobs int32 // 刚完成的任务数量

	limiter := rate.NewLimiter(rate.Every(time.Second/time.Duration(maxJobsPerSecond)), maxJobsPerSecond)
	semaphore := make(chan struct{}, maxJobsPerSecond)

	for i := 1; i <= numJobs; i++ {
		wg.Add(1)
		go func(jobID int) {
			defer wg.Done()
			limiter.Wait(context.Background()) // 等待限流器允许进行下一个任务

			semaphore <- struct{}{} // 获取信号量
			atomic.AddInt32(&startedJobs, 1)
			atomic.AddInt32(&runningJobs, 1)

			executeJob(jobID) // 执行任务
			atomic.AddInt32(&finishedJobs, 1)
			atomic.AddInt32(&runningJobs, -1)

			<-time.After(time.Second) // 等待一秒钟后释放信号量
			<-semaphore

			// 打印当前状态
			printStatus(&runningJobs, &startedJobs, &finishedJobs)
		}(i)
	}

	wg.Wait()
	fmt.Println("所有工作完成")
}

func executeJob(jobID int) {
	startTime := time.Now() // 记录任务开始时间

	// 模拟任务执行时间
	fmt.Printf("%v Job %d started\n", time.Now().Format("2006-01-02 15:04:05.000"), jobID)
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 随机生成一个时间间隔（例如，1到5000毫秒之间）
	min := 1
	max := 5000
	duration := time.Duration(rand.Intn(max-min+1)+min) * time.Millisecond
	time.Sleep(duration)

	durationCost := time.Since(startTime) // 计算任务耗时

	fmt.Printf("%v Job %d finished Cost:%v\n", time.Now().Format("2006-01-02 15:04:05.000"), jobID, durationCost)
}

func printStatus(runningJobs, startedJobs, finishedJobs *int32) {
	fmt.Printf("Current status - Running: %d, Started: %d, Finished: %d\n",
		atomic.LoadInt32(runningJobs),
		atomic.LoadInt32(startedJobs),
		atomic.LoadInt32(finishedJobs))
}
