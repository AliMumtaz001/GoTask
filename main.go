package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type multiples struct {
	word, line, paragraph, consient, vowel, space, digit, punctuation, Special int
}

func combinefunc(s string) multiples {
	m := multiples{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' {
			m.word++
		}
		if s[i] == ' ' {
			m.space++
		}
		if s[i] == '.' {
			m.line++
		}
		if s[i] == '\n' {
			m.paragraph++
		}
		if s[i] == 'b' || s[i] == 'c' || s[i] == 'd' || s[i] == 'f' || s[i] == 'g' ||
			s[i] == 'h' || s[i] == 'j' || s[i] == 'k' || s[i] == 'l' || s[i] == 'm' ||
			s[i] == 'n' || s[i] == 'p' || s[i] == 'q' || s[i] == 'r' || s[i] == 's' ||
			s[i] == 't' || s[i] == 'v' || s[i] == 'w' || s[i] == 'x' || s[i] == 'y' ||
			s[i] == 'z' || s[i] == 'B' || s[i] == 'C' || s[i] == 'D' || s[i] == 'F' ||
			s[i] == 'G' || s[i] == 'H' || s[i] == 'J' || s[i] == 'K' || s[i] == 'L' ||
			s[i] == 'M' || s[i] == 'N' || s[i] == 'P' || s[i] == 'Q' || s[i] == 'R' ||
			s[i] == 'S' || s[i] == 'T' || s[i] == 'V' || s[i] == 'W' || s[i] == 'X' ||
			s[i] == 'Y' || s[i] == 'Z' {
			m.consient++
		}
		if s[i] == '!' || s[i] == '?' || s[i] == ',' || s[i] == '.' {
			m.punctuation += 1
		}
		if s[i] == '%' || s[i] == '@' || s[i] == '#' || s[i] == '$' || s[i] == '^' || s[i] == '&' ||
			s[i] == '*' || s[i] == '(' || s[i] == ')' || s[i] == '_' || s[i] == '+' || s[i] == '-' ||
			s[i] == '=' || s[i] == '{' || s[i] == '}' || s[i] == '[' || s[i] == ']' || s[i] == '|' ||
			s[i] == '\\' || s[i] == ':' || s[i] == ';' || s[i] == '"' || s[i] == '\'' || s[i] == '<' ||
			s[i] == '>' || s[i] == '/' || s[i] == '~' || s[i] == '`' {
			m.Special += 1
		}
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			m.vowel += 1
		}
		if s[i] == '0' || s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' ||
			s[i] == '5' || s[i] == '6' || s[i] == '7' || s[i] == '8' || s[i] == '9' {
			m.digit += 1
		}
	}
	// fmt.Println("Total Characters:  ", len(str))
	fmt.Println("Total Words", m.word)
	fmt.Println("Total Lines", m.line)
	fmt.Println("Total Paragraphs are", m.paragraph)
	fmt.Println("Total Consinent are", m.consient)
	fmt.Println("Total Vowels are", m.vowel)
	fmt.Println("Total Space are", m.space)
	fmt.Println("Total Digits are", m.digit)
	fmt.Println("Total Punctuation", m.punctuation)
	fmt.Println("Total Special Characters are", m.Special)
	return m
}

// function for Words count
func WordCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' {
			word += 1
		}
	}
	return word
}

// function for space count
func SpaceCount(s string) int {
	space := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			space += 1
		}
	}
	return space

}

// function for line count
func LineCount(s string) int {
	line := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			line += 1
		}
	}
	return line

}

// function for Paragraph count
func ParaCount(s string) int {
	paragraph := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			paragraph += 1
		}
	}
	return paragraph + 1

}

// function for Consient count
func ConsinentCount(s string) int {
	cons := 0
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
			cons++
		}
	}
	return cons
}

// function for Punctuation count
func PuncCount(s string) int {
	punct := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '!' || s[i] == '?' || s[i] == ',' || s[i] == '.' {
			punct += 1
		}
	}
	return punct

}

// function for Special Char count
func SpecialCount(s string) int {
	special := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '%' || s[i] == '@' || s[i] == '#' || s[i] == '$' || s[i] == '^' || s[i] == '&' ||
			s[i] == '*' || s[i] == '(' || s[i] == ')' || s[i] == '_' || s[i] == '+' || s[i] == '-' ||
			s[i] == '=' || s[i] == '{' || s[i] == '}' || s[i] == '[' || s[i] == ']' || s[i] == '|' ||
			s[i] == '\\' || s[i] == ':' || s[i] == ';' || s[i] == '"' || s[i] == '\'' || s[i] == '<' ||
			s[i] == '>' || s[i] == '/' || s[i] == '~' || s[i] == '`' {
			special += 1
		}
	}
	return special

}

// function for Vowels count
func VowelCount(s string) int {
	vowel := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			vowel += 1
		}
	}
	return vowel
}

// function for Digits count
func DigitCount(s string) int {
	digit := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' || s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' ||
			s[i] == '5' || s[i] == '6' || s[i] == '7' || s[i] == '8' || s[i] == '9' {
			digit += 1
		}
	}
	return digit
}

func main() {
	start := time.Now() //for executionn time

	content, error := os.ReadFile("Testfile.txt") // to Read file
	if error != nil {                             // if err then log
		log.Fatal(error)
	}

	str := string(content) //coonvert it into string

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
	fmt.Println(".......Chunking......")
	chunking(str,0,b)
	chunking(str,b,b2)
	chunking(str,b2,b3)
	chunking(str,b3,b4)


	fmt.Println("\n.......Separate functions......")
	fmt.Println("Total Characters:  ", len(str))
	fmt.Println("Total Words", WordCount(str))
	fmt.Println("Total Lines", LineCount(str))
	fmt.Println("Total Paragraphs are", ParaCount(str))
	fmt.Println("Total Consinent are", ConsinentCount(str))
	fmt.Println("Total Vowels are", VowelCount(str))
	fmt.Println("Total Space are", SpaceCount(str))
	fmt.Println("Total Digits are", DigitCount(str))
	fmt.Println("Total Punctuation", PuncCount(str))
	fmt.Println("Total Special Characters are", SpaceCount(str))
	// elapse := time.Since(start)
	// fmt.Printf("page took %s \n", elapse)

	fmt.Println("Total Characters:  ", len(str))
	fmt.Println("Total Characters:  ", combinefunc(str))
	elapse := time.Since(start)
	fmt.Printf("page took %s \n", elapse)
	// }

}
func chunking(s string, start int, end int) multiples {
	m := multiples{}
	for i := start; i < end; i++ {
		if s[i] == ' ' || s[i] == '\n' {
			m.word++
		}
		if s[i] == ' ' {
			m.space++
		}
		if s[i] == '.' {
			m.line++
		}
		if s[i] == '\n' {
			m.paragraph++
		}
		if s[i] == 'b' || s[i] == 'c' || s[i] == 'd' || s[i] == 'f' || s[i] == 'g' ||
			s[i] == 'h' || s[i] == 'j' || s[i] == 'k' || s[i] == 'l' || s[i] == 'm' ||
			s[i] == 'n' || s[i] == 'p' || s[i] == 'q' || s[i] == 'r' || s[i] == 's' ||
			s[i] == 't' || s[i] == 'v' || s[i] == 'w' || s[i] == 'x' || s[i] == 'y' ||
			s[i] == 'z' || s[i] == 'B' || s[i] == 'C' || s[i] == 'D' || s[i] == 'F' ||
			s[i] == 'G' || s[i] == 'H' || s[i] == 'J' || s[i] == 'K' || s[i] == 'L' ||
			s[i] == 'M' || s[i] == 'N' || s[i] == 'P' || s[i] == 'Q' || s[i] == 'R' ||
			s[i] == 'S' || s[i] == 'T' || s[i] == 'V' || s[i] == 'W' || s[i] == 'X' ||
			s[i] == 'Y' || s[i] == 'Z' {
			m.consient++
		}
		if s[i] == '!' || s[i] == '?' || s[i] == ',' || s[i] == '.' {
			m.punctuation += 1
		}
		if s[i] == '%' || s[i] == '@' || s[i] == '#' || s[i] == '$' || s[i] == '^' || s[i] == '&' ||
			s[i] == '*' || s[i] == '(' || s[i] == ')' || s[i] == '_' || s[i] == '+' || s[i] == '-' ||
			s[i] == '=' || s[i] == '{' || s[i] == '}' || s[i] == '[' || s[i] == ']' || s[i] == '|' ||
			s[i] == '\\' || s[i] == ':' || s[i] == ';' || s[i] == '"' || s[i] == '\'' || s[i] == '<' ||
			s[i] == '>' || s[i] == '/' || s[i] == '~' || s[i] == '`' {
			m.Special += 1
		}
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' ||
			s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			m.vowel += 1
		}
		if s[i] == '0' || s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' ||
			s[i] == '5' || s[i] == '6' || s[i] == '7' || s[i] == '8' || s[i] == '9' {
			m.digit += 1
		}
	}
	// fmt.Println("Total Characters:  ", len(str))
	
	fmt.Println("Total Words", m.word)
	fmt.Println("Total Lines", m.line)
	fmt.Println("Total Paragraphs are", m.paragraph)
	fmt.Println("Total Consinent are", m.consient)
	fmt.Println("Total Vowels are", m.vowel)
	fmt.Println("Total Space are", m.space)
	fmt.Println("Total Digits are", m.digit)
	fmt.Println("Total Punctuation", m.punctuation)
	fmt.Println("Total Special Characters are", m.Special)
	return m
}
