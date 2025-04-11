package Analyse
import (
	// "fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/readfile", func(c *gin.Context) {
		filename := "Test2.txt" // Jo file read karni hai

		// File read karna
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "File read error: " + err.Error()})
			return
		}

		// File ka content response mein send karna
		c.String(http.StatusOK, string(data))
	})

	router.Run(":8080") // Server start
}
