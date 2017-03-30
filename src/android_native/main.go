package main

import (
	"sync"
	"time"
	"fmt"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"log"
	"golang.org/x/mobile/event/touch"
)

func worker(id int, tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		d := time.Duration(4) * time.Millisecond
		time.Sleep(d)
		fmt.Println("worker", id, "processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(i, tasksCh, wg)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}

func dothat() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1000)
	go pool(&wg, 1000, 3000)
	wg.Wait()
	fmt.Println("\nDuration: ", time.Since(start), "\n")
}

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				log.Print("Lifecycle: "+e.String())
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
