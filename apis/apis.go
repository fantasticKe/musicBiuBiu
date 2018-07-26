package apis

import (
	"github.com/gin-gonic/gin"
	"musicBiuBiu/spider"
	"net/http"
)

func Biubiu(c *gin.Context) {

	c.String(http.StatusOK, "hello, musicBiuBiu")
}

func Spider(c *gin.Context) {
	songId := c.Param("songId")
	ntesMusic.GetAllComments(songId)
	c.String(http.StatusOK, songId)
}
