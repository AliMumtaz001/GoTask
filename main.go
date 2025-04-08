package main

import (
	"time"
	"fmt"
	"log"
	"os"
	"sync"
)
func main() {
	start := time.Now() //for executionn time
	time.Sleep(time.Second * 2)

	var wg sync.WaitGroup //Now you initialized one WaitGroup (wg)
    Chans := make(chan int,10)
	
	content, error := os.ReadFile("Testfile.txt") // to Read file
	if error != nil {	// if err then log  
		log.Fatal(error)
	}
	str := string(content) //coonvert it into string
	wg.Add(10)			//creates 10 Channels (Chans)
	go func ()  {
		Chans <- len(str)  //value Stre in chanel
		wg.Done()	//job is ok
	}()
	TotalChar := <- Chans //then again value recieve from Chans and same this process for remaining 9 chans
	go func ()  {
		Chans <- WordCount(str)
		wg.Done()
	}()
	Totalword := <- Chans
	go func ()  {
		Chans <- LineCount(str)
		wg.Done()
	}()
	TotalLine := <- Chans
	go func ()  {
		Chans <- ParaCount(str)
		wg.Done()
	}()
	TotalPara := <- Chans
	go func ()  {
		Chans <- ConsinentCount(str)
		wg.Done()
	}()
	TotalCons := <- Chans
	go func ()  {
		Chans <- VowelCount(str)
		wg.Done()
	}()
	TotalVowel := <- Chans
	go func ()  {
		Chans <- SpaceCount(str)
		wg.Done()
	}()
	TotalSpace := <- Chans
	go func ()  {
		Chans <- DigitCount(str)
		wg.Done()
	}()
	TotalDigit := <- Chans
	go func ()  {
		Chans <- PuncCount(str)
		wg.Done()
	}()
	TotalPun := <- Chans
	go func ()  {
		Chans <- SpecialCount(str)
		wg.Done()
	}()
	TotalSpec := <- Chans
	go func () {
		wg.Wait()
		close(Chans)
	}()
	// arr :=[10]string{"Total Characters are :","Total Words are :","Total Lines are :","Total Paragraphs are :",
	// 				"Total Consinent are : ","Total Vowels are : ","Total Space are : ","Total Digits are :",
	// 				"Total Punctuation are : ", "Total Special Characters are : "}
	// for a := range Chans{
	// 	for i:=0;i<9;i++{
	// 		final:=arr[i]
	// 		fmt.Println(final,a)
	// 	}	
	// }
		
	fmt.Println("Total Characters:  ", TotalChar)
	fmt.Println("Total Words", Totalword)
	fmt.Println("Total Lines", TotalLine)
	fmt.Println("Total Paragraphs are", TotalPara)
	fmt.Println("Total Consinent are", TotalCons)
	fmt.Println("Total Vowels are", TotalVowel)
	fmt.Println("Total Space are", TotalSpace)
	fmt.Println("Total Digits are", TotalDigit)
	fmt.Println("Total Punctuation", TotalPun)
	fmt.Println("Total Special Characters are", TotalSpec)
	
	elapse :=time.Since(start)
	fmt.Printf("page took %s \n",elapse)
}

func WordCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' {
			word += 1
		}
	}
	return word + 1

}
func SpaceCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			word += 1
		}
	}
	return word + 1

}
func LineCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			word += 1
		}
	}
	return word + 1

}
func ParaCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if  s[i] == '\n'{ 
				word += 1
		}
	}
	return word + 1

}
func ConsinentCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' || s[i] == 'c' || s[i] == 'd' || s[i] == 'f' || s[i] == 'g' ||
			s[i] == 'h' || s[i] == 'j' || s[i] == 'k' || s[i] == 'l' || s[i] == 'm' ||
			s[i] == 'n' || s[i] == 'p' || s[i] == 'q' || s[i] == 'r' || s[i] == 's' ||
			s[i] == 't' || s[i] == 'v' || s[i] == 'w' || s[i] == 'x' || s[i] == 'y' ||
			s[i] == 'z' || s[i] == 'B' || s[i] == 'C' || s[i] == 'D' || s[i] == 'F' || s[i] == 'G' ||
			s[i] == 'H' || s[i] == 'J' || s[i] == 'K' || s[i] == 'L' || s[i] == 'M' ||
			s[i] == 'N' || s[i] == 'P' || s[i] == 'Q' || s[i] == 'R' || s[i] == 'S' ||
			s[i] == 'T' || s[i] == 'V' || s[i] == 'W' || s[i] == 'X' || s[i] == 'Y' ||
			s[i] == 'Z' {
				word++
		}
	}
	return word + 1

}

func PuncCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if   s[i] == '!' || s[i] == '?' || s[i] == ',' || s[i] == '.' {
			word += 1
		}
	}
	return word + 1

}

func SpecialCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '%' || s[i] == '@' || s[i] == '#' || s[i] == '$' || s[i] == '^' || s[i] == '&' || 
		 s[i] == '*' || s[i] == '(' || s[i] == ')' || s[i] == '_' || s[i] == '+' || s[i] == '-' || 
		 s[i] == '=' || s[i] == '{' || s[i] == '}'|| s[i] == '[' || s[i] == ']' || s[i] == '|' || 
		 s[i] == '\\' || s[i] == ':' || s[i] == ';' || s[i] == '"' || s[i] == '\'' || s[i] == '<' || 
		 s[i] == '>'   || s[i] == '/' || s[i] == '~' || s[i] == '`'  {
			word += 1
		}
	} //# $ % ^ & * ( ) _ + - = { } [ ] | \ : ; " ' < > , . ? / ~ `
	return word + 1

}

func VowelCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
		 s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			word += 1
		}
	}
	return word + 1
}
func DigitCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' || s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' || 
		s[i] == '5' || s[i] == '6' || s[i] == '7' || s[i] == '8' || s[i] == '9' {
			word += 1
		}
	}
	return word + 1
}

// func WordCount(s string) map[string]int{
// 	words := strings.Fields(s)
// 	m := make(map[string]int)
// 	for _, words := range words{
// 		m[words] += 1
// 	}
// 	fmt.Println(m)
// 	return m
// }
