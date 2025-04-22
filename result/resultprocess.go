package resultprocess

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/AliMumtaz001/GoTask/utils"
)

func SaveResult(db *sql.DB, result utils.Multiples, userID int) error {
    query := `
        INSERT INTO results 
        (words, digits, special_char, lines, spaces, consonants, vowels, sentences, paragraphs, punctuation, user_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

    _, err := db.Exec(query,
        result.Words,
        result.Digits,
        result.SpecialChar,
        result.Lines,
        result.Spaces,
        result.Consonants,
        result.Vowels,
        result.Sentences,
        result.Paragraphs,
        result.Punctuation,
        userID, 
    )
    if err != nil {
        log.Println("Error inserting result:", err)
        return err
    }

    fmt.Println("Result saved successfully.")
    return nil
}