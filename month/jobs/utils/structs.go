package utils

// 定义结构体以匹配JSON数据结构
type Data struct {
	Action    string `json:"action"`
	Timestamp int64  `json:"timestamp"`
	Date      struct {
		ID        int     `json:"id"`
		Avatar    string  `json:"avatar"`
		CreatedAt string  `json:"createdAt"`
		DeletedAt *string `json:"deletedAt"` // 使用指针类型以处理null值
		Mobile    string  `json:"mobile"`
		Name      string  `json:"name"`
		Password  string  `json:"password"`
		Status    int     `json:"status"`
		UpdatedAt string  `json:"updatedAt"`
		Version   string  `json:"version"`
	} `json:"date"`
}
