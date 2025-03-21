package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Accès autorisé à l'endpoint protégé"})
}
