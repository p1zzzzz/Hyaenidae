package service

import (
	"Hyaenidae/global"
	"Hyaenidae/model"
	"Hyaenidae/utils/upload"
	"errors"
	"mime/multipart"
	"strings"
)

//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func Upload(file model.FileUploadAndDownload) error {
	return global.Hyaenidae_DB.Create(&file).Error
}

//@function: FindFile
//@description: 删除文件切片记录
//@param: id uint
//@return: error, model.ExaFileUploadAndDownload

func FindFile(id uint) (error, model.FileUploadAndDownload) {
	var file model.FileUploadAndDownload
	err := global.Hyaenidae_DB.Where("id = ?", id).First(&file).Error
	return err, file
}



//@function: DeleteFile
//@description: 删除文件记录
//@param: file model.ExaFileUploadAndDownload
//@return: err error

func DeleteFile(file model.FileUploadAndDownload) (err error) {
	var fileFromDb model.FileUploadAndDownload
	err, fileFromDb = FindFile(file.ID)
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.Hyaenidae_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}


//@function: UploadFile
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func UploadFile(header *multipart.FileHeader, noSave string) (err error, file model.FileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(err)
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := model.FileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return Upload(f), f
	}
	return
}