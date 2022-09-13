package models

const (
	ARTICLE_STATE_ON  = 10
	ARTICLE_STATE_OFF = 20
)

type Article struct {
	Model
	Title       string `json:"article_title"`
	Content     string `json:"article_content"`
	Description string `json:"article_desc"`
	State       int    `json:"article_state"`
	CreatedOn   string `json:"article_create_on"`
	CreatedBy   string `json:"article_create_by"`
}

type APIArticle struct {
	ID    uint
	Title string
	Desc  string
}

func ArticlesList(pageOffset int, pageSize int, maps interface{}) (articles []Article) {
	db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&articles)
	return
}

func ArticlesDetail(id int) (article Article) {
	db.Select("id, title, content, description, state, created_on").Where("id = ?", id).First(&article)
	return
}
