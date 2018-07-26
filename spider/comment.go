package ntesMusic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"musicBiuBiu/orm"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/**
	获取歌曲评论
    id:歌曲id
	limit:每次请求条数
	offset:请求起始点
*/
func GetComments(id string, offset, limit int) (commentReps string, err error) {
	rid := ""
	strOffset := strconv.Itoa(offset)
	strLimit := strconv.Itoa(limit)
	total := "true"
	initStr1 := `{"rid": "` + rid + `","offset": "` + strOffset + `","total": "` + total + `","limit": "` + strLimit + `","csrf_token": ""}`
	params1, key1, err := EncParams(initStr1)
	if err != nil {
		panic(err)
	}
	// 发送POST请求得到最后包含url的结果
	commentReps, err = Comments(params1, key1, id)
	if err != nil {
		fmt.Println(err)
		return commentReps, err
	}
	return commentReps, err

}
func GetAllComments(songId string) {
	offset := 0
	fmt.Printf("开始获取歌曲id:%s的所有评论\n", songId)
	time.Sleep(1 * time.Millisecond)
	i := 0
	startTime := time.Now()
	var data string
	var err error
	for {
		data, err = GetComments(songId, offset, offset+40)
		fmt.Println(data)
		var commentsRep *CommentsRep
		if err != nil {
			log.Println(err)
		}
		if err1 := json.Unmarshal([]byte(data), &commentsRep); err1 != nil {
			panic(err1)
		}
		for _, c := range commentsRep.Comments {
			comment := orm.Comment{
				UserId:     c.User.UserId,
				NickName:   c.User.Nickname,
				AvatarUrl:  c.User.AvatarUrl,
				Content:    c.Content,
				LikesCount: c.LikedCount,
				Time:       c.Time,
			}
			orm.AddComment(comment)
		}

		i++
		if offset > int(commentsRep.Total) {
			log.Printf("这首歌一共请求%d次获取所有评论\n", i)
			log.Printf("offset为%d \n", offset)
			fmt.Printf("停止获取歌曲id:%s的所有评论\n", songId)
			log.Println("获取这首歌所有评论一共花费时间:", time.Now().Sub(startTime))
			break
		}
		offset += 20
	}
}

/**
发送post请求
*/
func Comments(params string, encSecKey string, id string) (commentRep string, err error) {
	client := &http.Client{}
	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	body := strings.NewReader(form.Encode())
	request, _ := http.NewRequest("POST", "http://music.163.com/weapi/v1/resource/comments/R_SO_4_"+id+"?csrf_token=", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Content-Length", (string)(body.Len()))
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.101 Safari/537.36")
	request.Header.Set("Cookie", "_ntes_nnid=f2c441d1440900d6daa9611bab3dc027,1515122355101; _ntes_nuid=f2c441d1440900d6daa9611bab3dc027; __utmz=94650624.1515122355.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); _iuqxldmzr_=32; __remember_me=true; JSESSIONID-WYYY=YXZtk7tOBJ4b3gOVrX2hl5%2BBriZyYVR5kNX3D3G5oWFRcY3J1cvGnMJRZx6JXgVSRNhFKO3O%5CmRiRACwWjrhBnkmK3dgGyTawDSAAmF%2Fct5T%2BhYVRy1BnxCgx%5CYrAUrjnQ8jEJQ1VHJTdNhqS4p9jVxHdRcc7iv5cQn649a%5CsBTc46WR%3A1515402120148; __utma=94650624.753127024.1515122355.1515218466.1515400320.9; __utmc=94650624; MUSIC_U=0120a3f48157438f759f2034b3925668ae731f8ae462a842927650798e0d663c97d1b459676c0cc693e926b3c390b8ba205ba14613b02d6c02d1ccf53040f6087d9739a0cccfd7eebf122d59fa1ed6a2; __csrf=5aa926378397ed694496ebf6486c5dfc; __utmb=94650624.5.10.1515400320")
	// 发起请求
	response, reqErr := client.Do(request)
	// 错误处理
	if reqErr != nil {
		fmt.Println("Fatal error ", reqErr.Error())
		return commentRep, reqErr
	}
	defer response.Body.Close()
	resBody, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(resBody))
	/*err = json.Unmarshal(resBody, &commentRep)
	if err != nil {
		fmt.Println(err)
		return commentRep, err
	}*/
	return string(resBody), nil
}
