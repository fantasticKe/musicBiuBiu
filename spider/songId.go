package ntesMusic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	Rex         = `<ul class="f-hide"><li><a href="/song\?id=\d*?">.*</a></li></ul>`
	SongNameRex = `<li><a href="/song\?id=\d*?">(.*?)</a></li>`
	SongIdRex   = `<li><a href="/song\?id=(\d*?)">.*?</a></li>`
)

func SongId(listId string) (songMap map[string]string, err error) {
	var songs = make(map[string]string)

	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, "https://music.163.com/discover/toplist?id="+listId, nil)
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	request.Header.Set("Cookie", "_iuqxldmzr_=32; _ntes_nnid=f158b6e8fe90cf72cee74376a65e9339,1527516758509; _ntes_nuid=f158b6e8fe90cf72cee74376a65e9339; WM_TID=6HGcMyznTB5f%2FilGrSxH6osEfu4HaL4Z; __remember_me=true; __utmz=94650624.1532179258.5.4.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; hb_MA-BFF5-63705950A31C_source=www.baidu.com; __f_=1532439176486; __utma=187553192.1197532187.1532439410.1532439410.1532439410.1; __utmz=187553192.1532439410.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __oc_uuid=9e2cf7f0-8f46-11e8-ac7a-89775e8e5afa; __utmc=94650624; playerid=25550276; WM_NI=Ue6TpT0jd1fZF3JuRZ0kEkvKj2y%2FS9rvCgT2O8bqSmQoacDkxw%2BhmXioSxCNBFM%2FlZsGFXNgvFICbbJBkuntVfvgq62Bu27Y10mXCpM8CRxpBFEemOGQEnxjmLv0tbJGOFc%3D; WM_NIKE=9ca17ae2e6ffcda170e2e6eea4c880a8869b92ae5c8786bf8feb40b0948d8dcd6fb28eb887c648aab8ffb8ea2af0fea7c3b92a97bbfab5f360f4b6acd2e55286ed8696d152f4b58b8cf541aaeeb7d0d233b2aa8d88d73a8de8ae92d26793b1af84c761bca9e589f054abb19d99c246829a83d0ef52b1998ab9cd6db69a87a4e16a8a9ea596ed8096baba88f43ef29ebbd3b653f3f0a3b9c27c8fe88db3fb7c829a8b86f045858f8b8ffb62af94bad1f76a8a9981a7e237e2a3; JSESSIONID-WYYY=UbMzDw8pp7ugHsbAeaWG%5CY%2BvYTSRWSytNekYqqmj0d%5ChxVRAVuowfjpojueq%2FIHpvh%2FahsOpEpPKW9fxxquYk8512eHdV2xFitPk7ZeGEt4oBFzIKngdfw%2Ftg7130TaYo7%2FKZaYf7%2Fh0TlOT%2Brat8zB%5CnAt83lI3v64rSMDkGRJg6gsF%3A1532753716818; MUSIC_U=9de0b492647e466198b5bb0393731d8bb4e139e0be72e0620cef52331cdfdc2ed466dc58db998dded62d64ca66d2e60e0cdf63e06733d5840331dcb041be5556305842396b5dfc01; __csrf=a2cd542a8e80677afa75ab2d052bdf20; __utma=94650624.677775120.1527516760.1532613459.1532752108.12; __utmb=94650624.6.10.1532752108")
	// 发起请求
	response, reqErr := client.Do(request)
	// 错误处理
	if reqErr != nil {
		fmt.Println("Fatal error ", reqErr.Error())
		return songs, reqErr
	}
	defer response.Body.Close()
	resBody, _ := ioutil.ReadAll(response.Body)
	exp := regexp.MustCompile(Rex)
	result := exp.FindAllStringSubmatch(string(resBody), -1)
	songNameExp := regexp.MustCompile(SongNameRex)
	songIdRex := regexp.MustCompile(SongIdRex)
	songNames := songNameExp.FindAllStringSubmatch(result[0][0], -1)
	songIds := songIdRex.FindAllStringSubmatch(result[0][0], -1)
	for i, name := range songNames {
		songs[name[1]] = songIds[i][1]
	}
	return songs, nil
}
