package models

type RelationArticleTag struct {
	Model
	TagId     string `json:"tag_id"`
	ArticleId string `json:"article_id"`
	CreatedOn string `json:"-"`
	CreatedBy string `json:"-"`
}
