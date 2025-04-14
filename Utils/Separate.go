package utils

// function for Words count
func WordCount(s string) int {
	word := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '.' {
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
