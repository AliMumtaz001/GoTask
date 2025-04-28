package repositories

import (
	"database/sql"
	"log"

	model "github.com/AliMumtaz001/GoTask/models"
)

type ResultsRepository struct {
	DB *sql.DB
}

// GetTotalRecords retrieves the total number of records for a user.
func (repo *ResultsRepository) GetTotalRecords(userID int) (int, error) {
	var totalRecords int
	query := `SELECT COUNT(*) FROM results WHERE user_id = $1`
	err := repo.DB.QueryRow(query, userID).Scan(&totalRecords)
	if err != nil {
		return 0, err
	}
	return totalRecords, nil
}

// GetResults retrieves paginated results for a user.
func (repo *ResultsRepository) GetResults(userID, pageSize, offset int) ([]model.Multiples, error) {
	query := `
        SELECT words, digits, special_char, lines, spaces, punctuation, consonants, vowels, sentences, paragraphs 
        FROM results 
        WHERE user_id = $1 
        LIMIT $2 OFFSET $3`
	rows, err := repo.DB.Query(query, userID, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []model.Multiples
	for rows.Next() {
		var result model.Multiples
		err := rows.Scan(
			&result.Words,
			&result.Digits,
			&result.SpecialChar,
			&result.Lines,
			&result.Spaces,
			&result.Punctuation,
			&result.Consonants,
			&result.Vowels,
			&result.Sentences,
			&result.Paragraphs,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// SaveResult inserts analysis results into the database.
func (repo *ResultsRepository) SaveResult(result model.Multiples, userID int) error {
	query := `
        INSERT INTO results 
        (words, digits, special_char, lines, spaces, consonants, vowels, sentences, paragraphs, punctuation, user_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := repo.DB.Exec(query,
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

	log.Println("Result saved successfully.")
	return nil
}

func (repo *ResultsRepository) GetUserIDByEmail(email string) (int, error) {
    var userID int
    query := `SELECT id FROM employee WHERE email = $1`
    err := repo.DB.QueryRow(query, email).Scan(&userID)
    return userID, err
}

func (repo *ResultsRepository) SaveUploadedResult(result string, userID int) error {
    query := `INSERT INTO results (user_id, result) VALUES ($1, $2)`
    _, err := repo.DB.Exec(query, userID, result)
    return err
}
