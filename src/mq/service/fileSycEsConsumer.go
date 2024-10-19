package service

import "UploadFileProject/src/entity/dto"

type SynFileEsConsumer struct{}

var SynFileEsConsumerImpl = &SynFileEsConsumer{}

func (synFile *SynFileEsConsumer) SynFileEs(data dto.SynEsDTO) {
	//message := `{
	//
	//}`
}
