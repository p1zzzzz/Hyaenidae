package source

import (
	"Hyaenidae/global"
	"Hyaenidae/model"
	"Hyaenidae/utils"

	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

var Admin = new(admin)

type admin struct{}

var admins = []model.SysUser{
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "", NickName: "超级管理员", HeaderImg: "http://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
}

//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	admins[0].Password = utils.ScriptPW("admin123")
	return global.Hyaenidae_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.SysUser{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
