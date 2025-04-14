package Utils

import (
	"fmt"
	"sync"
)

func Chunking(s string, start int, end int, wg *sync.WaitGroup, ch chan<- Multiples) {
	defer wg.Done()
	m := Multiples{}
	for i := start; i < end && i < len(s); i++ {
		if s[i] == '\n' || (s[i] == '.' && i+1 < len(s) && s[i+1] == '\n') {
			m.Paragraph++
		}
		if s[i] == ' ' || s[i] == '\n' {
			m.Word++
		}
		if s[i] == ' ' {
			m.Space++
		} else if s[i] == '.' {
			m.Line++
		} else if s[i] == 'b' || s[i] == 'c' || s[i] == 'd' || s[i] == 'f' || s[i] == 'g' ||
			s[i] == 'h' || s[i] == 'j' || s[i] == 'k' || s[i] == 'l' || s[i] == 'm' ||
			s[i] == 'n' || s[i] == 'p' || s[i] == 'q' || s[i] == 'r' || s[i] == 's' ||
			s[i] == 't' || s[i] == 'v' || s[i] == 'w' || s[i] == 'x' || s[i] == 'y' ||
			s[i] == 'z' || s[i] == 'B' || s[i] == 'C' || s[i] == 'D' || s[i] == 'F' ||
			s[i] == 'G' || s[i] == 'H' || s[i] == 'J' || s[i] == 'K' || s[i] == 'L' ||
			s[i] == 'M' || s[i] == 'N' || s[i] == 'P' || s[i] == 'Q' || s[i] == 'R' ||
			s[i] == 'S' || s[i] == 'T' || s[i] == 'V' || s[i] == 'W' || s[i] == 'X' ||
			s[i] == 'Y' || s[i] == 'Z' {
			m.Consient++
		} else if s[i] == '!' || s[i] == '?' || s[i] == ',' || s[i] == '.' {
			m.Punctuation += 1
		} else if s[i] == '%' || s[i] == '@' || s[i] == '#' || s[i] == '$' || s[i] == '^' || s[i] == '&' ||
			s[i] == '*' || s[i] == '(' || s[i] == ')' || s[i] == '_' || s[i] == '+' || s[i] == '-' ||
			s[i] == '=' || s[i] == '{' || s[i] == '}' || s[i] == '[' || s[i] == ']' || s[i] == '|' ||
			s[i] == '\\' || s[i] == ':' || s[i] == ';' || s[i] == '"' || s[i] == '\'' || s[i] == '<' ||
			s[i] == '>' || s[i] == '/' || s[i] == '~' || s[i] == '`' {
			m.Special += 1
		} else if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			m.Vowel += 1
		} else if s[i] == '0' || s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' ||
			s[i] == '5' || s[i] == '6' || s[i] == '7' || s[i] == '8' || s[i] == '9' {
			m.Digit += 1
		}
	}
	fmt.Println("..................Chunking..................\n")
	fmt.Println("Total Words", m.Word)
	fmt.Println("Total Lines", m.Line)
	fmt.Println("Total Paragraphs are", m.Paragraph)
	fmt.Println("Total Consinent are", m.Consient)
	fmt.Println("Total Vowels are", m.Vowel)
	fmt.Println("Total Space are", m.Space)
	fmt.Println("Total Digits are", m.Digit)
	fmt.Println("Total Punctuation", m.Punctuation)
	fmt.Println("Total Special Characters are", m.Special)
	ch <- m // Send result to channel
}
