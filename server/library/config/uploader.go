package config

type Uploader struct {
	ChunkDir  string `mapstructure:"chunk-dir" json:"chunkDir" yaml:"chunk-dir"`
	FinishDir string `mapstructure:"finish-dir" json:"finishDir" yaml:"finish-dir"`
}

func (u *Uploader) GetCheckPath(md5 string) string {
	return u.ChunkDir + md5
}

func (u *Uploader) GetIdentifier(identifier string) string {
	return u.ChunkDir + identifier + "/"
}
