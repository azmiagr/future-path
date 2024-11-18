package service

import (
	"errors"
	"future-path/entity"
	"future-path/internal/repository"
	"future-path/model"
)

type ISekolahService interface {
	GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error)
	GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error)
	GetAllSekolah(page int) ([]*entity.Sekolah, error, int64)
	GetSekolahDetail(id int) (*entity.Sekolah, error)
	AddSekolah(sekolahReq *model.CreateSekolah) (*entity.Sekolah, error)
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

func (ss *SekolahService) GetAllSekolah(page int) ([]*entity.Sekolah, error, int64) {
	limit := 10
	offset := (page - 1) * limit
	sekolah, err := ss.SekolahRepository.GetAllSekolah(limit, offset)
	if err != nil {
		return nil, err, 0
	}

	totalData, err := ss.SekolahRepository.CountAllSekolah()

	return sekolah, nil, totalData
}

func (ss *SekolahService) GetSekolahDetail(id int) (*entity.Sekolah, error) {
	sekolah, err := ss.SekolahRepository.GetSekolahDetail(id)
	if err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (ss *SekolahService) AddSekolah(sekolahReq *model.CreateSekolah) (*entity.Sekolah, error) {
	sekolah := &entity.Sekolah{
		Nama_Sekolah:      sekolahReq.Nama_Sekolah,
		Alamat_Sekolah:    sekolahReq.Alamat_Sekolah,
		Deskripsi_Sekolah: sekolahReq.Deskripsi_Sekolah,
		ID_Kepemilikan:    sekolahReq.ID_Kepemilikan,
	}

	if sekolah.ID_Kepemilikan != 1 && sekolah.ID_Kepemilikan != 2 {
		return nil, errors.New("invalid ownerships id")
	}

	sekolah, err := ss.SekolahRepository.AddSekolah(sekolah)
	if err != nil {
		return nil, err
	}
	return sekolah, nil
}
