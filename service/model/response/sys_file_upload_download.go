package response

import "Hyaenidae/model"

type SysFileResponse struct {
	File model.FileUploadAndDownload `json:"file"`
}

