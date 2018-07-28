package orm

import (
	"gopkg.in/mgo.v2"
)

type Comment struct {
	Id         int64  `bson:"_id"`
	MusicName  string `bson:"musicName"`
	MusicId    string `bson:"musicId"`
	UserId     int64  `bson:"userId"`
	NickName   string `bson:"nickName"`
	AvatarUrl  string `bson:"avatarUrl"`
	Content    string `bson:"content"`
	LikesCount int    `bson:"likesCount"`
	Time       int64  `bson:"time"`
}

const URL = "localhost:27018"

var (
	mgoSession *mgo.Session
	dataBase   = "maoke"
	collection = "ntesMusic_comments"
)

/**
获取session，如果存在则拷贝一份
*/
func GetSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		if mgoSession, err = mgo.Dial(URL); err != nil {
			panic(err)
		}
	}
	//最大连接池4096
	return mgoSession.Clone()
}

/**
获取collection对象
*/
func WitchCollection(col string, s func(session *mgo.Collection) error) error {
	session := GetSession()
	defer session.Close()
	c := session.DB(dataBase).C(col)
	return s(c)
}

/**
存入评论
*/
func AddComment(com Comment) int64 {
	query := func(c *mgo.Collection) error {
		return c.Insert(com)
	}
	if err := WitchCollection(collection, query); err != nil {
		return 0
	}
	return com.Id
}
