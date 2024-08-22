package vo

import "UploadFileProject/src/entity/dto"

type FileListVO struct {
	dto.FileSearchDTO
	FileList []dto.FileListDTO `json:"fileList"`
}
