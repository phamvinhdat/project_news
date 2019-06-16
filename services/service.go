package services

import(
	"mime/multipart"
)
type IImg_service interface{
	Save(*multipart.FileHeader) (string, error)
	Delete(string) error
}
