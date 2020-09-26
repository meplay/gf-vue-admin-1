package utils

import "mime/multipart"

var Oss OSS

type OSS interface {
	Upload(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}
