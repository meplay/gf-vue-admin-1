package request

type CreateSimpleUpload struct {
	ChunkNumber      string `p:"chunkNumber" v:"required|length:1,1000#请输入当前切片标记|当前切片标记长度为:min到:max位"`
	CurrentChunkPath string `p:"currentChunkPath" v:"required|length:1,1000#请输入切片本地路径|切片本地路径长度为:min到:max位"`
	CurrentChunkSize string `p:"currentChunkSize" v:"required|length:1,1000#请输入当前切片容量|当前切片容量长度为:min到:max位"`
	FilePath         string `p:"filePath" v:"required|length:1,1000#请输入文件本地路径|文件本地路径长度为:min到:max位"`
	Filename         string `p:"filename" v:"required|length:1,1000#请输入文件名|文件名长度为:min到:max位"`
	Identifier       string `p:"identifier" v:"required|length:1,1000#请输入文件标识（md5）|文件标识（md5）长度为:min到:max位"`
	IsDone           bool   `p:"isDone" v:"required|length:1,1000#请输入是否上传完成|是否上传完成长度为:min到:max位"`
	TotalChunks      string `p:"totalChunks" v:"required|length:1,1000#请输入切片总数|切片总数长度为:min到:max位"`
	TotalSize        string `p:"totalSize" v:"required|length:1,1000#请输入总容量|总容量长度为:min到:max位"`
}
