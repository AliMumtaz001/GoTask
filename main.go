package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"task_program/Utils"
	"time"
)

func main() {
	start := time.Now()                           //for executionn time
	content, error := os.ReadFile("Testfile.txt") // to Read file
	if error != nil {                             // if err then log
		log.Fatal(error)
	}

	str := string(content) //coonvert it into string
	// Chunking methoddology

	chanel := make(chan Utils.Multiples, 5)
	start0 := time.Now()
	length := len(str)

	chunks := [5]int{}
	for i := 0; i < 5; i++ {
		if i == 0 {

			chunks[i] = length / 5
			fmt.Println(length)
		} else {
			chunks[i] = chunks[i-1] + chunks[0]
		}
	}
	fmt.Println(chunks)
	var wg sync.WaitGroup
	b := chunks[0]
	b2 := chunks[1]
	b3 := chunks[2]
	b4 := chunks[3]
	b5 := chunks[4]

	fmt.Println(".......Chunking......\n")
	wg.Add(5)
	go func() {
		defer wg.Done()
		Utils.Chunking(str, 0, b, &wg)
	}()

	go func() {
		defer wg.Done()
		Utils.Chunking(str, b, b2, &wg)
	}()

	// fmt.Println(".......Chunk2......\n")
	go func() {
		defer wg.Done()
		Utils.Chunking(str, b2, b3, &wg)
	}()

	// fmt.Println("\n.......Chunk3......\n")
	go func() {
		defer wg.Done()
		Utils.Chunking(str, b3, b4, &wg)
	}()

	// fmt.Println("\n.......Chunk4......\n")
	go func() {
		defer wg.Done()
		Utils.Chunking(str, b4, b5, &wg)
	}()
	// fmt.Println("\n.......Chunk5......\n")
	data := <-chanel
	fmt.Println("......................Chunking data......................", data)
	wg.Wait()
	close(chanel)

	elapse := time.Since(start0)
	fmt.Printf("page took %s \n", elapse)

	fmt.Println("\n.......Separate functions......")
	fmt.Println("Total Characters:  ", len(str))
	fmt.Println("Total Words", Utils.WordCount(str))
	fmt.Println("Total Lines", Utils.LineCount(str))
	fmt.Println("Total Paragraphs are", Utils.ParaCount(str))
	fmt.Println("Total Consinent are", Utils.ConsinentCount(str))
	fmt.Println("Total Vowels are", Utils.VowelCount(str))
	fmt.Println("Total Space are", Utils.SpaceCount(str))
	fmt.Println("Total Digits are", Utils.DigitCount(str))
	fmt.Println("Total Punctuation", Utils.PuncCount(str))
	fmt.Println("Total Special Characters are", Utils.SpaceCount(str))
	elapse0 := time.Since(start)
	fmt.Printf("page took %s \n", elapse0)

	start1 := time.Now()
	fmt.Print("................Combined function.............\n")
	multiples := Utils.CombineFunc(str)
	fmt.Println("Total Characters:  ", len(str))
	fmt.Println("Total Words", multiples.Word)
	fmt.Println("Total Lines", multiples.Line)
	fmt.Println("Total Paragraphs are", multiples.Paragraph)
	fmt.Println("Total Consinent are", multiples.Consient)
	fmt.Println("Total Vowels are", multiples.Vowel)
	fmt.Println("Total Space are", multiples.Space)
	fmt.Println("Total Digits are", multiples.Digit)
	fmt.Println("Total Punctuation", multiples.Punctuation)
	fmt.Println("Total Special Characters are", multiples.Special)
	elapse1 := time.Since(start1)
	fmt.Printf("page took %s \n", elapse1)
}
