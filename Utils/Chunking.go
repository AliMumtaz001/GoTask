package utils

import (
	"fmt"
	"os"
	"sync"
)

func Chunking(s string, start int, end int) Multiples {
	
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
	
	return m // Send result to channel
}

func Analyzer(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	str := string(data)
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

	var wg sync.WaitGroup
	ch := make(chan Multiples, 5)

	for i := 0; i < len(chunks)-1; i++ {
		wg.Add(1)
		
		go func(start, end int) {
			defer wg.Done()
			result := Chunking(str, start, end)
			ch <- result
		}(chunks[i], chunks[i+1])
	}

	wg.Wait()
	close(ch)

	var finalResult Multiples
for part := range ch {
    finalResult.Word += part.Word
    finalResult.Line += part.Line
    finalResult.Paragraph += part.Paragraph
    finalResult.Consient += part.Consient
    finalResult.Vowel += part.Vowel
    finalResult.Space += part.Space
    finalResult.Digit += part.Digit
    finalResult.Punctuation += part.Punctuation
    finalResult.Special += part.Special
}
resultString := fmt.Sprintf("Words: %d, Lines: %d, Paragraphs: %d, Consonants: %d, Vowels: %d, Spaces: %d, Digits: %d, Punctuation: %d, Special: %d",
    finalResult.Word, finalResult.Line, finalResult.Paragraph, finalResult.Consient, finalResult.Vowel,
    finalResult.Space, finalResult.Digit, finalResult.Punctuation, finalResult.Special)

return resultString, nil
}