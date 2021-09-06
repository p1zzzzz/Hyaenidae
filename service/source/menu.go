package source

import (
	"Hyaenidae/global"
	"Hyaenidae/model"
	"time"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []model.SysBaseMenu{
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "setting"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 2, Meta: model.Meta{Title: "服务器状态", Icon: "cloudy"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "user", Name: "user", Component: "view/superAdmin/user/user.vue", Sort: 2, Meta: model.Meta{Title: "用户管理", Icon: "coordinate"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "markdown", Name: "markdown", Hidden: false, Component: "view/markdown/markdown.vue", Sort: 4, Meta: model.Meta{Title: "markdown", Icon: "notebook-1"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "userdashboard", Name: "userdashboard", Hidden: true, Component: "view/userdashboard/index.vue", Sort: 5, Meta: model.Meta{Title: "仪表盘", Icon: "setting"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "person", Name: "person", Hidden: false, Component: "view/person/person.vue", Sort: 6, Meta: model.Meta{Title: "个人信息", Icon: "user-solid"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "tool", Name: "tool", Hidden: false, Component: "view/tool/index.vue", Sort: 7, Meta: model.Meta{Title: "tools", Icon: "s-tools"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "cyberchef", Name: "cyberchef", Hidden: false, Component: "view/tool/cyberchef/cyberchef.vue", Sort: 1, Meta: model.Meta{Title: "cyberchef", Icon: "link"}},
	{Hyaenidae_MODEL: global.Hyaenidae_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "whatav", Name: "whatav", Hidden: false, Component: "view/tool/whatav/whatav.vue", Sort: 2, Meta: model.Meta{Title: "whatav", Icon: "question"}},
}

//@description: sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return global.Hyaenidae_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 5}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
