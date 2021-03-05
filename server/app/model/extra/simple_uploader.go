package model

type SimpleUploader struct {
	Filename         string `orm:"filename" json:"filename" gorm:"comment:文件名"`
	FilePath         string `orm:"file_path" json:"filePath" gorm:"comment:文件本地路径"`
	TotalSize        string `orm:"total_size" json:"totalSize" gorm:"comment:总容量"`
	Identifier       string `orm:"identifier" json:"identifier" gorm:"comment:文件标识（md5）"`
	ChunkNumber      string `orm:"chunk_number" json:"chunkNumber" gorm:"comment:当前切片标记"`
	TotalChunks      string `orm:"total_chunks" json:"totalChunks" gorm:"comment:切片总数"`
	CurrentChunkPath string `orm:"current_chunk_path" json:"currentChunkPath" gorm:"comment:切片本地路径"`
	CurrentChunkSize string `orm:"current_chunk_size" json:"currentChunkSize" gorm:"comment:当前切片容量"`
	IsDone           bool   `orm:"is_done" json:"isDone" gorm:"comment:是否上传完成"`
}

func (s *SimpleUploader) TableName() string {
	return "simple_uploader"
}

