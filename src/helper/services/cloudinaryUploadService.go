package services

import (
	 "context"
    "fmt"
    "mime/multipart"
	"strings"
    "github.com/bchanona/Api_chombi.git/src/helper/config"
    "github.com/cloudinary/cloudinary-go/v2/api/admin/search"
    "github.com/cloudinary/cloudinary-go/v2/api/uploader"
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

func (s *CloudinaryService) UploadPdf(file multipart.File, fileName string) (string, error) {
    cld := config.InitCloudinary()

    // Quitar TODAS las extensiones posibles
    nameWithoutExt := fileName
    nameWithoutExt = strings.TrimSuffix(nameWithoutExt, ".pdf")
    nameWithoutExt = strings.TrimSuffix(nameWithoutExt, ".PDF")


    resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
        PublicID:     nameWithoutExt,
        Folder:       "pdfs",
        ResourceType: "image",
    })
    if err != nil {
        return "", err
    }
    return resp.SecureURL, nil
}


func (s *CloudinaryService) GetPdfUrl() ([]string, error) {
    cld := config.InitCloudinary()

    resp, err := cld.Admin.Search(context.Background(), search.Query{
        Expression: "folder:pdfs",
        MaxResults: 100,
    })

    if err != nil {
        return nil, fmt.Errorf("error en search: %v", err)
    }

    urls := []string{}
    for _, asset := range resp.Assets {
        urls = append(urls, asset.SecureURL)
    }
    return urls, nil
}