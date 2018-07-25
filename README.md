# musicBiuBiu

##歌曲评论弹幕biubiubiu

1. 抓取网易云歌曲评论存放到mongo中
2. 页面输入歌曲名称建立websocket连接。
3. 根据歌曲名称查询歌曲评论，并通过websocket推送到页面已弹幕的形式展示。
4. 若mongo中没有歌曲评论则开启实时爬虫去抓取歌曲评论。