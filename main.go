package main

import (
	"fmt"
	"golang/Folder01"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()                           //for executionn time
	content, error := os.ReadFile("Testfile.txt") // to Read file
	if error != nil {                             // if err then log
		log.Fatal(error)
	}

	str := string(content) //coonvert it into string
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

	b := chunks[0]
	b2 := chunks[1]
	b3 := chunks[2]
	b4 := chunks[3]
	b5 := chunks[4]
	fmt.Println(".......Chunking......\n")
	fmt.Println(".......Chunk1......\n")
	Folder01.Chunking(str, 0, b)
	fmt.Println(".......Chunk2......\n")
	Folder01.Chunking(str, b, b2)
	fmt.Println("\n.......Chunk3......\n")
	Folder01.Chunking(str, b2, b3)
	fmt.Println("\n.......Chunk4......\n")
	Folder01.Chunking(str, b3, b4)
	fmt.Println("\n.......Chunk5......\n")
	Folder01.Chunking(str, b4, b5)
	elapse := time.Since(start0)
	fmt.Printf("page took %s \n", elapse)

	fmt.Println("\n.......Separate functions......")
	fmt.Println("Total Characters:  ", len(str))
	fmt.Println("Total Words", Folder01.WordCount(str))
	fmt.Println("Total Lines", Folder01.LineCount(str))
	fmt.Println("Total Paragraphs are", Folder01.ParaCount(str))
	fmt.Println("Total Consinent are", Folder01.ConsinentCount(str))
	fmt.Println("Total Vowels are", Folder01.VowelCount(str))
	fmt.Println("Total Space are", Folder01.SpaceCount(str))
	fmt.Println("Total Digits are", Folder01.DigitCount(str))
	fmt.Println("Total Punctuation", Folder01.PuncCount(str))
	fmt.Println("Total Special Characters are", Folder01.SpaceCount(str))
	elapse0 := time.Since(start)
	fmt.Printf("page took %s \n", elapse0)

	start1 := time.Now()
	fmt.Print("................Combined function.............\n")
	multiples := Folder01.CombineFunc(str)
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

//	words, lines, paragraphs, consonants, vowels, spaces, digits, punctuations, specials
