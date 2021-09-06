package response

import "Hyaenidae/model"

type SysMarkdownResponse struct {
	Markdown  model.Markdown `json:"markdown"`
}

type SysMarkdownsResponse struct {
	Markdowns []model.Markdown `json:"markdown"`
}
