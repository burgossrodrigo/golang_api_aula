package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	eg := r.Group("/v1")
}
