// package utils

// type Multiples struct {
// 	Word        int
// 	Line        int
// 	Paragraph   int
// 	Consient    int
// 	Vowel       int
// 	Space       int
// 	Digit       int
// 	Punctuation int
// 	Special     int
// 	Sentences   int
// }

package utils

type Multiples struct {
	Words       int `json:"words"`
	Lines       int `json:"lines"`
	Paragraphs  int `json:"paragraphs"`
	Consonants  int `json:"consonants"`
	Vowels      int `json:"vowels"`
	Spaces      int `json:"spaces"`
	Digits      int `json:"digits"`
	Punctuation int `json:"punctuation"`
	SpecialChar int `json:"special_char"`
	Sentences   int `json:"sentences"`
}
