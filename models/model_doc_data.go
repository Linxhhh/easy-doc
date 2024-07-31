package models

type DocDataModel struct {
	Model
	DocId     uint     `gorm:"column:docId;comment:" json:"docId"`
	DocTitle  string   `gorm:"column:docTitle;comment:文档标题" json:"docTitle"`
	DiggCount int      `gorm:"column:diggCount;comment:点赞量" json:"diggCount"`
	ReadCount int      `gorm:"column:readCount;comment:浏览量" json:"readCount"`
}

/*
	文档数据表，记录文档的浏览量、点赞量等数据！
*/
