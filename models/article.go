package models

type Article struct {
	Model

	TagID int `json: "tag_id" gorm: "index"`
	Tag   Tag `json: "tag"`

	Title      string `json: "title"`
	Desc       string `json: "desc"`
	Content    string `json: "content"`
	CreateBy   string `json: "create_by"`
	ModifiedBy string `json: "modified_by"`
	State      int    `json: "state"`
}
