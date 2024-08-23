package dto

import "time"

type FileListDTO struct {
	FileName   string    `json:"fileName"`
	FileUuid   string    `json:"fileUuid"`
	FileSuffix string    `json:"fileSuffix"`
	UpdateTime time.Time `json:"updateTime"`
}
