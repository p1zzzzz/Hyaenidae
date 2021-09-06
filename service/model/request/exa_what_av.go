package request

// Find by uuid structure
type TaskList struct {
	SearchKey string `json:"searchkey" form:"searchkey"`
}