package interfaces

import "mime/multipart"

type Oss interface {
	Upload(file *multipart.FileHeader) (path string, key string, err error)
	Delete(key string) error
}
