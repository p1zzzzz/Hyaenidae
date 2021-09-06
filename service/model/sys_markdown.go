package model

import (
	"Hyaenidae/global"
	uuid "github.com/satori/go.uuid"
)

type Markdown struct {
	global.Hyaenidae_MODEL
	UUID     uuid.UUID    `json:"uuid" gorm:"comment:MdUUID"`  //md文件UUID
	Editor   string     `json:"editor" gorm:"comment:编辑者"`     // 编辑者
	Title    string     `json:"title" gorm:"comment:标题"`       // 标题
	MkValue  string     `json:"mkvalue" gorm:"comment:文章内容;type:longtext"`   // 文章内容
	ParentId string     `json:"parentId" gorm:"comment:父菜单ID;default:0"` // ID
	Children []Markdown `json:"children" gorm:"-"`
}
