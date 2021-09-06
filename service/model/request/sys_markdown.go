package request

import uuid "github.com/satori/go.uuid"

// User register structure
type SaveMd struct {
	UUID  uuid.UUID `json:"uuid"`
	Editor  string `json:"editor"`
	Title   string `json:"title"`
	MkValue string `json:"mkValue"`
	ParentId string `json:"parentId" gorm:"default:0"`
}
