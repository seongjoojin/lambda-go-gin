package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "test lambda API"})
}

func main() {
	addr := ":3000"
	mode := os.Getenv("GIN_MODE")
	g := gin.New()
	g.GET("/test", test)

	if mode == "release" {
		log.Fatal(gateway.ListenAndServe(addr, g))
	} else {
		log.Fatal(http.ListenAndServe(addr, g))
	}
}
