package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"musicBiuBiu/orm"
	"musicBiuBiu/spider"
	"net/http"
	"sync"
)

func Biubiu(c *gin.Context) {

	c.String(http.StatusOK, "hello, musicBiuBiu")
}

func Spider(c *gin.Context) {
	listId := c.Param("songId")
	SpiderComments(listId)
	c.String(http.StatusOK, listId)
}

func GetCommentsByMusicName(c *gin.Context) {
	musicName := c.Param("musicName")
	var comments []orm.Comment
	var err error
	if comments, err = orm.GetCommentsByMusicName(musicName); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, comments)
}

/**
根据歌单id抓取该歌单所有歌曲的评论
*/
func SpiderComments(songId string) {
	var wg sync.WaitGroup

	nteComments := ntesMusic.NteComments{
		Wc: make(chan orm.Comment, 100),
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if hotMusic, err := ntesMusic.SongId(songId); err == nil {
			fmt.Println(hotMusic)
			nteComments.GetListComments(hotMusic)
		}
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for com := range nteComments.Wc {
				id := orm.AddComment(com)
				fmt.Println(id)
			}
		}()
	}
	wg.Wait()
}
