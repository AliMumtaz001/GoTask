package main

import (
	"fmt"
	// "strings"
	"log"
	"os"
)

func main() {
	content, error := os.ReadFile("Testfile.txt")
	if error != nil {
		log.Fatal(error)
	}
	str := string(content)
	fmt.Println("Total Words", WordCount(str))
	fmt.Println("Total Lines", LineCount(str))
	fmt.Println("Total Punctuation", PuncCount(str))
	fmt.Println("Total Vowels are", VowelCount(str))
	fmt.Println("Total Space are", SpaceCount(str))
	fmt.Println("Total Digits are", DigitCount(str))
	fmt.Println("Total Characters are", len(str))
	fmt.Println("Total Paragraphs are", ParaCount(str))

}

func WordCount(s string) any {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' {
			word += 1
		}
	}
	return word + 1

}
func SpaceCount(s string) any {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			word += 1
		}
	}
	return word + 1

}
func LineCount(s string) any {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			word += 1
		}
	}
	return word + 1

}
func ParaCount(s string) any {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			if  s[i] == '\n' {
			word += 1
		}
		}
	}
	return word + 1

}

func PuncCount(s string) any {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == ',' || s[i] == '(' || s[i] == ')' || s[i] == ';' || s[i] == '!' || s[i] == '?' || s[i] == '\'' || s[i] == '%' || s[i] == '@'{
			word += 1
		}
	}
	return word + 1

}

func VowelCount(s string) any {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			word += 1
		}
	}
	return word + 1
}
func DigitCount(s string) any {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' || s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' || s[i] == '5' || s[i] == '6' || s[i] == '7' || s[i] == '8' || s[i] == '9' {
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
