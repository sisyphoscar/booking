package main

import (
	"booking/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	loadStaticFiles(router)

	router = routes.SetRoutes(router)

	log.Println("Server will running on http://localhost:8080")
	router.Run()
}

func loadStaticFiles(router *gin.Engine) {
	router.LoadHTMLGlob("web/templates/*.html")
	router.Static("/static", "web/static")
}
