package dto

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type FileUploadDTO struct {
	OrganizationUuid any
	File             *multipart.FileHeader
	C                *gin.Context
}
