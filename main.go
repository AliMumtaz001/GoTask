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
	fmt.Println(WordCount(str))
	// fmt.Println(m)
}

func WordCount(s string) any{
	word := 0
	for i:=0; i<len(s);i++{
		if s[i] == ' '|| s[i] == '\n' {
			word +=1 
		}
	}
	return word+1

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