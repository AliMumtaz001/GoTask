package authentication

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AliMumtaz001/GoTask/utils"
	"github.com/gin-gonic/gin"
)

func getdata(db *sql.DB, id string) (string, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "", err
	}

	var result utils.Multiples

	query := `
        SELECT words, digits, special_char, lines, spaces, punctuation, consonants, vowels, sentences, paragraphs 
        FROM result 
        WHERE id = $1`
	err = db.QueryRow(query, idInt).Scan(
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
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func GetResultHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing id query parameter"})
			return
		}

		jsonData, err := getdata(db, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id or server error", "details": err.Error()})
			return
		}

		if jsonData == "" {
			c.JSON(http.StatusNotFound, gin.H{"message": "No result found for this id"})
			return
		}

		c.JSON(http.StatusOK, gin.H{

			"Id": id,

			"result": json.RawMessage(jsonData)})
	}
}
