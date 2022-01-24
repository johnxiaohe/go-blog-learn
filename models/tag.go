package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreateBy   string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 查询标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)

	return tag.ID > 0
}

func AddTag(name string, state int, createBy string) bool {
	if !ExistTagByName(name) {
		db.Create(&Tag{
			Name:     name,
			State:    state,
			CreateBy: createBy,
		})
	}

	return true
}
