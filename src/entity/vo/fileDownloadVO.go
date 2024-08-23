package vo

type FileDownloadVO struct {
	OrganizationUuid string `json:"organizationUuid"`
	FileUuid         string `json:"fileUuid"`
	FileData         string `json:"fileData"`
}
