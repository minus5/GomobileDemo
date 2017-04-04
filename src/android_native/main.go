package main

import (
	"fmt"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/touch"
	"log"
	"sync"
	"time"
)

func worker(id int, tasksCh <-chan int) {
	for task := range tasksCh {
		time.Sleep(time.Duration(4) * time.Millisecond)
		fmt.Println("worker", id, "processing task", task)
	}
}

func pool(workers int, tasks int) {
	tasksCh := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerId int) {
			worker(workerId, tasksCh)
			wg.Done()
		}(i)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
	wg.Wait()
}

func dothat() {
	start := time.Now()
	pool(1000, 3000)
	duration := time.Since(start)
	time.Sleep(time.Duration(100) * time.Millisecond)
	fmt.Println("Duration: ", duration)
}

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				log.Print("Lifecycle: " + e.String())
			case paint.Event:
				log.Print("Call OpenGL here.")
				a.Publish()
			case touch.Event:
				if e.Type == touch.TypeEnd {
					go dothat()
				}
			}
		}
	})
}
