package msg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConflictErr(c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{"error": "Conflict"})
	c.Abort()
}

func BadRequestErr(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	c.Abort()
}

func InternalServerErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	c.Abort()
}
