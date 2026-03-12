package ports

import "mime/multipart"

type ImageStorage interface {
	UploadImage(file multipart.File, fileName string) (string, error)
}