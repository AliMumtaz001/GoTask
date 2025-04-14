package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/AliMumtaz001/GoTask/Utils"
)

func main() {
	start := time.Now()
	content, err := os.ReadFile("Test1.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := string(content)
	chanel := make(chan Utils.Multiples, 5)
	start0 := time.Now()
	length := len(str)

	chunks := [5]int{}
	chunkSize := length / 5
	for i := 0; i < 5; i++ {
		if i == 0 {
			chunks[i] = chunkSize
		} else {
			chunks[i] = chunks[i-1] + chunkSize
		}
	}
	chunks[4] = length
	fmt.Println(chunks)

	var wg sync.WaitGroup
	fmt.Println(".......Chunking......\n")
	wg.Add(5)

	go Utils.Chunking(str, 0, chunks[0], &wg, chanel)
	go Utils.Chunking(str, chunks[0], chunks[1], &wg, chanel)
	go Utils.Chunking(str, chunks[1], chunks[2], &wg, chanel)
	go Utils.Chunking(str, chunks[2], chunks[3], &wg, chanel)
	go Utils.Chunking(str, chunks[3], chunks[4], &wg, chanel)

	// goroutine 
	go func() {
		for i := 0; i < 5; i++ {
			data := <-chanel
			fmt.Println("......................Chunking data......................")
			fmt.Println("Chunk Result:", data)
		}
	}()

	wg.Wait()
	close(chanel)

	elapse := time.Since(start0)
	fmt.Printf("Chunking took %s\n", elapse)

	elapse0 := time.Since(start)
	fmt.Printf("Total execution took %s\n", elapse0)
}
