package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	if port != "" {
		r.Run(fmt.Sprintf(":%s", port))
	} else {
		r.Run()
	}

}
