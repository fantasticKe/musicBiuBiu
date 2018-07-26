package ntesMusic

type User struct {
	LocationInfo string `json:"-"`
	UserId       int64  `json:"userId"`
	AuthStatus   int    `json:"authStatus"`
	RemarkName   string `json:"-"`
	AvatarUrl    string `json:"avatarUrl"`
	Experts      string `json:"-"`
	Nickname     string `json:"nickname"`
	UserType     int    `json:"userType"`
	VipType      int    `json:"vipType"`
	ExpertTags   string `json:"-"`
}

type HotComment struct {
	User        *User  `json:"user"`
	BeReplied   string `json:"-"`
	Time        int64  `json:"time"`
	Linked      bool   `json:"linked"`
	CommentId   string `json:"commentId"`
	LinkedCount int    `json:"linkedCount"`
	content     string `json:"content"`
}

type Comment struct {
	User       *User    `json:"user"`
	BeReplied  []string `json:"-"`
	LikedCount int      `json:"likedCount"`
	Liked      bool     `json:"liked"`
	Time       int64    `json:"time"`
	CommentId  int64    `json:"commentId"`
	Content    string   `json:"content"`
}

type BeReplied struct {
	Id            int64
	User          *User `json:"-"`
	UserID        int64
	Content       string `json:"content"`
	Status        int    `json:"status"`
	CommentsID    int64
	HotCommentsID int64
}

type CommentsRep struct {
	IsMusician  bool       `json:"isMusician"`
	UserId      int        `json:"userId"`
	TopComments []string   `json:"topComments";gorm:"-"`
	MoreHot     bool       `json:"moreHot"`
	HotComments []*Comment `json:"hotComments"`
	Code        int        `json:"code"`
	Comments    []*Comment `json:"comments"`
	Total       int        `json:"total"`
	More        bool       `json:"more"`
}
