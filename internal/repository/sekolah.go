package repository

import (
	"future-path/entity"

	"gorm.io/gorm"
)

type ISekolahRepository interface {
	GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error)
	GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error)
	GetAllSekolah(limit, offset int) ([]*entity.Sekolah, error)
	GetSekolahDetail(id int) (*entity.Sekolah, error)
}

type SekolahRepository struct {
	db *gorm.DB
}

func NewSekolahRepository(db *gorm.DB) ISekolahRepository {
	return &SekolahRepository{db}
}

func (sk *SekolahRepository) GetSekolahNegeri(namaSekolah string) ([]*entity.Sekolah, error) {
	var sekolah []*entity.Sekolah
	query := "%" + namaSekolah + "%"
	if err := sk.db.Debug().Where("id_kepemilikan = ? AND nama_sekolah LIKE ?", 1, query).First(&sekolah).Error; err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (sk *SekolahRepository) GetSekolahSwasta(namaSekolah string) ([]*entity.Sekolah, error) {
	var sekolah []*entity.Sekolah
	query := "%" + namaSekolah + "%"
	if err := sk.db.Debug().Where("id_kepemilikan = ? AND nama_sekolah LIKE ?", 2, query).Find(&sekolah).Error; err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (sk *SekolahRepository) GetAllSekolah(limit, offset int) ([]*entity.Sekolah, error) {
	var sekolah []*entity.Sekolah
	if err := sk.db.Debug().Limit(limit).Offset(offset).Find(&sekolah).Error; err != nil {
		return nil, err
	}
	return sekolah, nil
}

func (sk *SekolahRepository) GetSekolahDetail(id int) (*entity.Sekolah, error) {
	var sekolah entity.Sekolah
	if err := sk.db.Debug().Where("id_sekolah = ?", id).First(&sekolah).Error; err != nil {
		return nil, err
	}
	return &sekolah, nil
}
