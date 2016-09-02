package main

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	absTemplate := filepath.Dir(os.Args[0])
	r.LoadHTMLGlob(absTemplate + "/../templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})

	})
	r.POST("/", func(c *gin.Context) {
		c.Request.ParseMultipartForm(32 << 20)
		c.JSON(200, gin.H{
			"message": "ok",
			"forms":   c.Request.MultipartForm.Value,
		})
	})
	r.Run()
}
