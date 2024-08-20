package dto

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type FileUploadDTO struct {
	OrganizationUuid string
	File             *multipart.FileHeader
	Context          *gin.Context
}

func NewFileUploadDTO(file *multipart.FileHeader, organizationUuid string, context *gin.Context) *FileUploadDTO {
	return &FileUploadDTO{File: file, OrganizationUuid: organizationUuid, Context: context}
}
