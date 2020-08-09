package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jpohlmann/KeypadCipherDecoder/model"
)

func main() {
	r := gin.Default()
	r.GET("/find", func(c *gin.Context) {
		code := c.Query("code")
		result := model.GetBestDecoding(code)
		c.JSON(200, gin.H {"bestMatch": result})
	})
	r.Run("0.0.0.0:3000")
}