package dto

import "github.com/gin-gonic/gin"

type FileDownloadDTO struct {
	OrganizationUuid string `json:"organizationUuid" mapstructure:"organizationUuid"`
	FileUuid         string `json:"fileUuid" mapstructure:"fileUuid"`
	Context          *gin.Context
}

func NewFileDownloadDTO(organizationUuid string, fileUuid string) *FileDownloadDTO {
	return &FileDownloadDTO{
		OrganizationUuid: organizationUuid,
		FileUuid:         fileUuid,
	}
}

func NewFileDownloadDTOWithContext(c *gin.Context) *FileDownloadDTO {
	return &FileDownloadDTO{Context: c}
}
