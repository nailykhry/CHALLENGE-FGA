package helpers

import (
	"github.com/gin-gonic/gin"
)

func GetContentType(copyC *gin.Context) string {
	return copyC.Request.Header.Get("Content-Type")
}
