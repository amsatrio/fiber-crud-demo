package domain

import (
	"io"
	"mime/multipart"
)

type MFileRepository interface {
	Delete(id string) error
	Upload(file *multipart.FileHeader, id int64, mUserId uint) error
	Stream(id uint) (io.ReadCloser, error)
}
