package service

import (
	"future-path/entity"
	"future-path/internal/repository"
)

type ISekolahService interface {
	GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error)
	GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error)
}

type SekolahService struct {
	SekolahRepository repository.ISekolahRepository
}

func NewSekolahService(SekolahRepository repository.ISekolahRepository) ISekolahService {
	return &SekolahService{SekolahRepository}
}

func (ss *SekolahService) GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error) {
	sekolah, err := ss.SekolahRepository.GetSekolahNegeri(namaSekolah)
	if err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (ss *SekolahService) GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error) {
	sekolah, err := ss.SekolahRepository.GetSekolahSwasta(namaSekolah)
	if err != nil {
		return nil, err
	}
	return sekolah, nil
}
