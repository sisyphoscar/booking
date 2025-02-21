package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home is the handler for the home page
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}
