package models

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type File struct {
	Name        string                 `json:"name"`
	Size        int64                  `json:"size"`
	Path        string                 `json:"path"`
	Bucket      string                 `json:"bucket"`
	ContentType string                 `json:"content_type"`
	Metadata    map[string]interface{} `json:"metadata"`
	MD5Hash     string                 `json:"md5_hash"` // md5 hash of the file
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	DeletedAt   string                 `json:"deleted_at"`
}
