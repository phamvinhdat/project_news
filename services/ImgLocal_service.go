package services

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ImgLocalService struct{}

func NewImgLocalService() IImg_service {
	return &ImgLocalService{}
}

//Save a file and return path and error
func (i *ImgLocalService) Save(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	//create empty file
	tempFile, err := ioutil.TempFile("public/view/images", "news-*"+filepath.Ext(fileHeader.Filename))
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	tempFile.Write(fileBytes)
	fileInfo, _ := tempFile.Stat()
	filePath := "images/" + fileInfo.Name()
	return filePath, nil
}

//Delete a file with path
func (i *ImgLocalService) Delete(path string) error {
	_, file := filepath.Split(path)
	err := os.Remove("public/view/images/" + file)
	if err == nil {
		return nil
	}

	return err
}
