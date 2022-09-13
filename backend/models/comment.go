package models

const (
	COMMENT_STATE_ON  = 1
	COMMENT_STATE_OFF = 0
)

type Comment struct {
	Model
	Content    string `json:"comment_content"`
	Url        string `json:"comment_url"`
	Email      string `json:"comment_email"`
	Tel        string `json:"comment_tel"`
	Name       string `json:"comment_name"`
	State      int    `json:"comment_state"`
	ArticleId  int    `json:"article_id"`
	CreatedOn  string `json:"comment_create_on"`
	ModifiedOn string `json:"comment_create_by"`
}

type APIComment struct {
	ID        uint
	Content   string `json:"comment_content"`
	Url       string `json:"comment_url"`
	Email     string `json:"comment_email"`
	Tel       string `json:"comment_tel"`
	Name      string `json:"comment_name"`
	State     int    `json:"comment_state"`
	ArticleId int    `json:"article_id"`
	CreatedOn string `json:"comment_create_on"`
}

func CommentsListByArticleId(pageOffset int, pageSize int, articleId int) (comments []Comment) {
	maps := make(map[string]interface{})
	maps["article_id"] = articleId
	db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&comments)
	return
}

func CommentsCreate(comment Comment) {
	db.Create(&comment)
}
