// package authentication

// import (
// 	"database/sql"
// 	"net/http"
// 	"os"
// 	"path/filepath"

// 	resultprocess "github.com/AliMumtaz001/GoTask/result"
// 	"github.com/AliMumtaz001/GoTask/utils"
// 	"github.com/gin-gonic/gin"
// )

// var db *sql.DB

// func Upload(c *gin.Context) {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "File not found"})
// 		return
// 	}

// 	// ensure the path
// 	if err := os.MkdirAll("upload", os.ModePerm); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create upload directory"})
// 		return
// 	}

// 	// save the file
// 	Path := filepath.Join("upload", file.Filename)
// 	err = c.SaveUploadedFile(file, Path)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
// 		return
// 	}

// 	content, err := os.ReadFile(Path)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read the uploaded file"})
// 		return
// 	}

// 	data := string(content)
// 	res := utils.CombineFunc(data)

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "File uploaded and analyzed successfully",
// 		"file":    file.Filename,
// 		"result":  res,
// 	})
// 	err = resultprocess.SaveResult(db, res)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save result to database"})
// 	}
// }


package authentication

import (
    "database/sql"
    "net/http"
    "os"
    "path/filepath"

    resultprocess "github.com/AliMumtaz001/GoTask/result"
    "github.com/AliMumtaz001/GoTask/utils"
    "github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, db *sql.DB) { // Add db parameter
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "File not found"})
        return
    }

    if err := os.MkdirAll("upload", os.ModePerm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create upload directory"})
        return
    }

    Path := filepath.Join("upload", file.Filename)
    err = c.SaveUploadedFile(file, Path)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
        return
    }

    content, err := os.ReadFile(Path)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read the uploaded file"})
        return
    }

    data := string(content)
    res := utils.CombineFunc(data)

    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "File uploaded and analyzed successfully",
        "file":    file.Filename,
        "result":  res,
    })

    err = resultprocess.SaveResult(db, res) // Pass db to SaveResult
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save result to database"})
    }
}