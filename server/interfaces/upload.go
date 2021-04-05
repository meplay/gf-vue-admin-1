package interfaces

import "mime/multipart"

type Oss interface {
	Upload(file *multipart.FileHeader) (filepath string, key string, err error)
	Delete(key string) error
}
