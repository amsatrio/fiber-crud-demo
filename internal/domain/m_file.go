package domain

import (
	"io"
	"mime/multipart"
)

type MFileRepository interface {
	Delete(id string) error
	Upload(file *multipart.FileHeader, id int64) error
	Stream(id uint) (io.ReadCloser, error)
}
