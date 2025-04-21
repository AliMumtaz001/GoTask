package resultprocess

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/AliMumtaz001/GoTask/utils"
)

func SaveResult(db *sql.DB, result utils.Multiples) error {
    query := `
        INSERT INTO result 
        (words, digits, special_char, lines, spaces, consonants, vowels, sentences, paragraphs, punctuation)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

    _, err := db.Exec(query,
        result.Words,
        result.Digits,
        result.SpecialChar,
        result.Lines,
        result.Spaces,
        result.Punctuation, // Fixed order to match the query
        result.Consonants,
        result.Vowels,
        result.Sentences,   // Fixed order to match the query
        result.Paragraphs,
    )
    if err != nil {
        log.Println("Error inserting result:", err)
        return err // Return the error instead of logging and ignoring it
    }

    fmt.Println("Result saved successfully.")
    return nil
}