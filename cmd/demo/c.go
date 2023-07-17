package main

import (
	"fmt"
	"sync"
)

func levelThreeWorker(level int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Level %d Worker starting\n", level)

	// 在这里执行第三层子线程的任务
	// ...

	fmt.Printf("Level %d Worker done\n", level)
}

func levelTwoWorker(level int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Level %d Worker starting\n", level)

	numWorkers := 3
	var wgLevelThree sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wgLevelThree.Add(1)
		go levelThreeWorker(i, &wgLevelThree)
	}

	wgLevelThree.Wait()

	fmt.Printf("Level %d Worker done\n", level)
}

func levelOneWorker(level int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Level %d Worker starting\n", level)

	numWorkers := 2
	var wgLevelTwo sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wgLevelTwo.Add(1)
		go levelTwoWorker(i, &wgLevelTwo)
	}

	wgLevelTwo.Wait()

	fmt.Printf("Level %d Worker done\n", level)
}

func main1() {
	numWorkers := 1
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go levelOneWorker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All workers done!")
}
