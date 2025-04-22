// package authentication

// import (
//     "database/sql"
//     "encoding/json"
//     "net/http"
//     "strconv"

//     "github.com/AliMumtaz001/GoTask/utils"
//     "github.com/gin-gonic/gin"
// )

// func getdata(db *sql.DB, userID string, page, pageSize int) (string, int, error) {
//     userIDInt, err := strconv.Atoi(userID)
//     if err != nil {
//         return "", 0, err
//     }

//     // Calculate offset for pagination
//     offset := (page - 1) * pageSize

//     // Query to get total count of records for the user
//     var totalRecords int
//     countQuery := `SELECT COUNT(*) FROM results WHERE user_id = $1`
//     err = db.QueryRow(countQuery, userIDInt).Scan(&totalRecords)
//     if err != nil {
//         return "", 0, err
//     }

//     // Query to fetch paginated results
//     var results []utils.Multiples
//     query := `
//         SELECT words, digits, special_char, lines, spaces, punctuation, consonants, vowels, sentences, paragraphs 
//         FROM results 
//         WHERE user_id = $1 
//         ORDER BY id 
//         LIMIT $2 OFFSET $3`
//     rows, err := db.Query(query, userIDInt, pageSize, offset)
//     if err != nil {
//         return "", 0, err
//     }
//     defer rows.Close()

//     for rows.Next() {
//         var result utils.Multiples
//         err := rows.Scan(
//             &result.Words,
//             &result.Digits,
//             &result.SpecialChar,
//             &result.Lines,
//             &result.Spaces,
//             &result.Punctuation,
//             &result.Consonants,
//             &result.Vowels,
//             &result.Sentences,
//             &result.Paragraphs,
//         )
//         if err != nil {
//             return "", 0, err
//         }
//         results = append(results, result)
//     }

//     if err = rows.Err(); err != nil {
//         return "", 0, err
//     }

//     if len(results) == 0 {
//         return "", totalRecords, nil
//     }

//     jsonData, err := json.Marshal(results)
//     if err != nil {
//         return "", 0, err
//     }

//     return string(jsonData), totalRecords, nil
// }

// func GetResultHandler(db *sql.DB) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         userID := c.Query("user_id") // Use user_id instead of id
//         if userID == "" {
//             c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id query parameter"})
//             return
//         }

//         // Get page and pageSize from query parameters (default to page 1, 4 records per page)
//         pageStr := c.Query("page")
//         pageSizeStr := c.Query("page_size")
//         page, err := strconv.Atoi(pageStr)
//         if err != nil || page < 1 {
//             page = 1
//         }
//         pageSize, err := strconv.Atoi(pageSizeStr)
//         if err != nil || pageSize < 1 {
//             pageSize = 4 // Default to 4 records per page
//         }

//         jsonData, totalRecords, err := getdata(db, userID, page, pageSize)
//         if err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or server error", "details": err.Error()})
//             return
//         }

//         if jsonData == "" {
//             c.JSON(http.StatusNotFound, gin.H{"message": "No results found for this user"})
//             return
//         }

//         c.JSON(http.StatusOK, gin.H{
//             "user_id":       userID,
//             "page":          page,
//             "page_size":     pageSize,
//             "total_records": totalRecords,
//             "total_pages":   (totalRecords + pageSize - 1) / pageSize,
//             "results":       json.RawMessage(jsonData),
//         })
//     }
// }

package authentication

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/AliMumtaz001/GoTask/utils"
    "github.com/gin-gonic/gin"
)

func getdata(db *sql.DB, userID string, page, pageSize int) (string, int, error) {
    userIDInt, err := strconv.Atoi(userID)
    if err != nil {
        return "", 0, err
    }

    // Calculate offset for pagination
    offset := (page - 1) * pageSize

    // Query to get total count of records for the user
    var totalRecords int
    countQuery := `SELECT COUNT(*) FROM results WHERE user_id = $1`
    err = db.QueryRow(countQuery, userIDInt).Scan(&totalRecords)
    if err != nil {
        return "", 0, err
    }

    // Query to fetch paginated results (removed ORDER BY id)
    var results []utils.Multiples
    query := `
        SELECT words, digits, special_char, lines, spaces, punctuation, consonants, vowels, sentences, paragraphs 
        FROM results 
        WHERE user_id = $1 
        LIMIT $2 OFFSET $3`
    rows, err := db.Query(query, userIDInt, pageSize, offset)
    if err != nil {
        return "", 0, err
    }
    defer rows.Close()

    for rows.Next() {
        var result utils.Multiples
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
            return "", 0, err
        }
        results = append(results, result)
    }

    if err = rows.Err(); err != nil {
        return "", 0, err
    }

    if len(results) == 0 {
        return "", totalRecords, nil
    }

    jsonData, err := json.Marshal(results)
    if err != nil {
        return "", 0, err
    }

    return string(jsonData), totalRecords, nil
}

func GetResultHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.Query("user_id") // Use user_id instead of id
        if userID == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user_id query parameter"})
            return
        }

        // Get page and pageSize from query parameters (default to page 1, 4 records per page)
        pageStr := c.Query("page")
        pageSizeStr := c.Query("page_size")
        page, err := strconv.Atoi(pageStr)
        if err != nil || page < 1 {
            page = 1
        }
        pageSize, err := strconv.Atoi(pageSizeStr)
        if err != nil || pageSize < 1 {
            pageSize = 4 // Default to 4 records per page
        }

        jsonData, totalRecords, err := getdata(db, userID, page, pageSize)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id or server error", "details": err.Error()})
            return
        }

        if jsonData == "" {
            c.JSON(http.StatusNotFound, gin.H{"message": "No results found for this user"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "user_id":       userID,
            "page":          page,
            "page_size":     pageSize,
            "total_records": totalRecords,
            "total_pages":   (totalRecords + pageSize - 1) / pageSize,
            "results":       json.RawMessage(jsonData),
        })
    }
}