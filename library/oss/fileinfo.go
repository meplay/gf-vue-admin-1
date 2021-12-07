package oss

import (
	"io"
	"io/fs"
	"net/http"
	"net/textproto"
	"time"

	"github.com/flipped-aurora/gf-vue-admin/interfaces"
)

type Option func(*FileInfo)

type FileInfo struct {
	Filename        string      // base name of the file
	FileContentType string      // Content-Type of the given data
	Filesize        int64       // length in bytes for regular files; system-dependent for others
	FileMode        fs.FileMode // file mode bits
	FileModTime     time.Time   // modification time
	FileIsDir       bool        // abbreviation for Mode().IsDir()
	FileSys         interface{} // underlying data source (can return nil)
	FileHeader      textproto.MIMEHeader
	reader          io.Reader
}

func NewFileInfo(options ...Option) interfaces.FileInfo {
	info := &FileInfo{}
	for i := 0; i < len(options); i++ {
		options[i](info)
	}
	return info
}

func NewFileInfoByFs(info fs.FileInfo, options ...Option) interfaces.FileInfo {
	_info := &FileInfo{
		Filename:    info.Name(),
		Filesize:    info.Size(),
		FileMode:    info.Mode(),
		FileModTime: info.ModTime(),
		FileIsDir:   info.IsDir(),
		FileSys:     info.Sys(),
	}
	for i := 0; i < len(options); i++ {
		options[i](_info)
	}

	return _info
}

func WithFilename(filename string) Option {
	return func(info *FileInfo) {
		info.Filename = filename
	}
}

func WithFilesize(filesize int64) Option {
	return func(info *FileInfo) {
		info.Filesize = filesize
	}
}

func WithFileMode(fileMode fs.FileMode) Option {
	return func(info *FileInfo) {
		info.FileMode = fileMode
	}
}

func WithFileModTime(fileModTime time.Time) Option {
	return func(info *FileInfo) {
		info.FileModTime = fileModTime
	}
}

func WithFileIsDir(fileIsDir bool) Option {
	return func(info *FileInfo) {
		info.FileIsDir = fileIsDir
	}
}

func WithFileSys(fileSys interface{}) Option {
	return func(info *FileInfo) {
		info.FileSys = fileSys
	}
}

func WithReader(reader io.Reader) Option {
	return func(info *FileInfo) {
		info.reader = reader
	}
}

func WithFileContentType(fileContentType string) Option {
	return func(info *FileInfo) {
		info.FileContentType = fileContentType
	}
}

func (f *FileInfo) Name() string {
	return f.Filename
}

func (f *FileInfo) Size() int64 {
	return f.Filesize
}

func (f *FileInfo) Mode() fs.FileMode {
	return f.FileMode
}

func (f *FileInfo) ModTime() time.Time {
	return f.FileModTime
}

func (f *FileInfo) IsDir() bool {
	return f.FileIsDir
}

func (f *FileInfo) Sys() interface{} {
	return f.FileSys
}

func (f *FileInfo) ContentType() string {
	if f.FileContentType != "" {
		return f.FileContentType
	}
	buf := make([]byte, 0, f.Filesize)
	_, _ = f.reader.Read(buf)
	return http.DetectContentType(buf)
}
