package services

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/bchanona/Api_chombi.git/src/helper/config"
)

type CloudinaryService struct{}

func (s *CloudinaryService) UploadImage(file multipart.File, fileName string) (string, error) {

	cld := config.InitCloudinary()

	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID: fileName,
		Folder:   "vehicles",
	})

	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}