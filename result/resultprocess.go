package resultprocess

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AliMumtaz001/GoTask/utils"
)

var dB *sql.DB

func SaveResult(dB, result utils.Multiples) error {
	query := `
		INSERT INTO result 
		(words, digits, special_char, lines, spaces, punctuation, consonants, vowels, sentences, paragraphs)
		VALUES ($2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := dB.Exec(query,
		result.Words,
		result.Digits,
		result.SpecialChar,
		result.Lines,
		result.Spaces,
		result.Sentences,
		result.Punctuation,
		result.Consonants,
		result.Vowels,
		result.Paragraphs,
	)

	if err != nil {
		log.Println("Error inserting result:", err)
		return err
	}
	fmt.Println("Result saved successfully.")

	return nil
}
