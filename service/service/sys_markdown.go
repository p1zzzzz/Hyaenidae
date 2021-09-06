package service

import (
	"Hyaenidae/global"
	"Hyaenidae/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

//@function: SaveMarkdown
//@description: 保存markdown
//@param: u model.Markdown
//@return: err error, MarkdownInter model.Markdown

func SaveMd(reqMark model.Markdown) (err error, MarkdownInter model.Markdown) {
	var m model.Markdown
	fmt.Println(reqMark)
	if !errors.Is(global.Hyaenidae_DB.Where("uuid = ?", reqMark.UUID).First(&m).Error, gorm.ErrRecordNotFound) { // 判断markdown是否存在
		err = global.Hyaenidae_DB.Where("uuid = ?", reqMark.UUID).Updates(&reqMark).Error
	}else{
		err = global.Hyaenidae_DB.Create(&reqMark).Error
	}
	return err, reqMark
}

//@function: DeleteMarkdown
//@description: 删除markdown
//@param: id float64
//@return: err error

func DeleteMd(uuid string) (err error) {
	var m model.Markdown
	err = global.Hyaenidae_DB.Where("uuid = ?", uuid).Delete(&m).Error
	return err
}

//@function: GetMdById
//@description: 查询markdown
//@param: id float64
//@return: err error

func GetMdById(id float64) (err error, markdown *model.Markdown) {
	var m model.Markdown
	err = global.Hyaenidae_DB.Where("id = ?", id).First(&m).Error
	return err,&m
}

//@function: GetMdByUuid
//@description: 查询markdown
//@param: uuid string
//@return: err error

func GetMdByUuid(uuid string) (err error, markdown *model.Markdown) {
	var m model.Markdown
	err = global.Hyaenidae_DB.Where("uuid = ?", uuid).First(&m).Error
	return err,&m
}



//@function: GetMdTree
//@description: 获取markdown树
//@return: err error, menus []model.Markdown

func GetMdTree() (err error, mds []model.Markdown) {
	err, treeMap := getMdTreeMap()
	mds = treeMap["0"]
	for i := 0; i < len(mds); i++ {
		err = getMdChildrenList(&mds[i], treeMap)
	}
	return err, mds
}

//@function: getMdTreeMap
//@description: 获取markdown树map
//@return: err error, treeMap map[string][]model.SysBaseMenu

func getMdTreeMap() (err error, treeMap map[string][]model.Markdown) {
	var allMenus []model.Markdown
	treeMap = make(map[string][]model.Markdown)
	err = global.Hyaenidae_DB.Order("ID").Select("id", "created_at","updated_at","deleted_at","uuid","editor","title","parent_id").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}



//@function: getMdChildrenList
//@description: 获取markdown的子markdown
//@param: menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//@return: err error

func getMdChildrenList(mds *model.Markdown, treeMap map[string][]model.Markdown) (err error) {
	mds.Children = treeMap[strconv.Itoa(int(mds.ID))]
	for i := 0; i < len(mds.Children); i++ {
		err = getMdChildrenList(&mds.Children[i], treeMap)
	}
	return err
}



