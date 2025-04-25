package utils

import (
	"os"
	"sync"
	"github.com/AliMumtaz001/GoTask/models"
)

func Chunking(s string, start int, end int) model.Multiples {
	m := model.Multiples{}
	for i := start; i < end && i < len(s); i++ {
		if s[i] == '\n' || (s[i] == '.' && i+1 < len(s) && s[i+1] == '\n') {
			m.Paragraphs++
		}
		if s[i] == ' ' || s[i] == '\n' {
			m.Words++
		}
		if s[i] == ' ' {
			m.Spaces++
		} else if s[i] == '.' {
			m.Lines++
			m.Sentences++
		} else if s[i] == 'b' || s[i] == 'c' || s[i] == 'd' || s[i] == 'f' || s[i] == 'g' ||
			s[i] == 'h' || s[i] == 'j' || s[i] == 'k' || s[i] == 'l' || s[i] == 'm' ||
			s[i] == 'n' || s[i] == 'p' || s[i] == 'q' || s[i] == 'r' || s[i] == 's' ||
			s[i] == 't' || s[i] == 'v' || s[i] == 'w' || s[i] == 'x' || s[i] == 'y' ||
			s[i] == 'z' || s[i] == 'B' || s[i] == 'C' || s[i] == 'D' || s[i] == 'F' ||
			s[i] == 'G' || s[i] == 'H' || s[i] == 'J' || s[i] == 'K' || s[i] == 'L' ||
			s[i] == 'M' || s[i] == 'N' || s[i] == 'P' || s[i] == 'Q' || s[i] == 'R' ||
			s[i] == 'S' || s[i] == 'T' || s[i] == 'V' || s[i] == 'W' || s[i] == 'X' ||
			s[i] == 'Y' || s[i] == 'Z' {
			m.Consonants++
		} else if s[i] == '!' || s[i] == '?' || s[i] == ',' || s[i] == '.' {
			m.Punctuation += 1
		} else if s[i] == '%' || s[i] == '@' || s[i] == '#' || s[i] == '$' || s[i] == '^' || s[i] == '&' ||
			s[i] == '*' || s[i] == '(' || s[i] == ')' || s[i] == '_' || s[i] == '+' || s[i] == '-' ||
			s[i] == '=' || s[i] == '{' || s[i] == '}' || s[i] == '[' || s[i] == ']' || s[i] == '|' ||
			s[i] == '\\' || s[i] == ':' || s[i] == ';' || s[i] == '"' || s[i] == '\'' || s[i] == '<' ||
			s[i] == '>' || s[i] == '/' || s[i] == '~' || s[i] == '`' {
			m.SpecialChar += 1
		} else if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			m.Vowels += 1
		} else if s[i] == '0' || s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' ||
			s[i] == '5' || s[i] == '6' || s[i] == '7' || s[i] == '8' || s[i] == '9' {
			m.Digits += 1
		}
	}
	return m
}

func Analyzer(filepath string) (model.Multiples, error) { // Change return type to Multiples
	data, err := os.ReadFile(filepath)
	if err != nil {
		return model.Multiples{}, err
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
	ch := make(chan model.Multiples, 5)

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

	var finalResult model.Multiples
	for part := range ch {
		finalResult.Words += part.Words
		finalResult.Lines += part.Lines
		finalResult.Paragraphs += part.Paragraphs
		finalResult.Consonants += part.Consonants
		finalResult.Vowels += part.Vowels
		finalResult.Spaces += part.Spaces
		finalResult.Digits += part.Digits
		finalResult.Punctuation += part.Punctuation
		finalResult.SpecialChar += part.SpecialChar
		finalResult.Sentences += part.Sentences
	}

	return finalResult, nil // Return the Multiples struct
}
