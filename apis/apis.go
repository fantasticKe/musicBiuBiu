package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Biubiu(c *gin.Context) {

	c.String(http.StatusOK, "hello, musicBiuBiu")
}
