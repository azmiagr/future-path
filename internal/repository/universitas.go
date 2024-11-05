package repository

import (
	"future-path/entity"

	"gorm.io/gorm"
)

type IUniversitasRepository interface {
	GetUnivNegeri(namaUniv string) ([]*entity.Universitas, error)
	GetUnivSwasta(namaUniv string) ([]*entity.Universitas, error)
}

type UniversitasRepository struct {
	db *gorm.DB
}

func NewUniversitasRepository(db *gorm.DB) IUniversitasRepository {
	return &UniversitasRepository{db}
}

func (u *UniversitasRepository) GetUnivNegeri(namaUniv string) ([]*entity.Universitas, error) {
	var univ []*entity.Universitas
	query := "%" + namaUniv + "%"
	if err := u.db.Debug().Where("id_kepemilikan = ? AND nama_universitas LIKE ?", 1, query).Find(&univ).Error; err != nil {
		return nil, err
	}
	return univ, nil
}

func (u *UniversitasRepository) GetUnivSwasta(namaUniv string) ([]*entity.Universitas, error) {
	var univ []*entity.Universitas
	query := "%" + namaUniv + "%"
	if err := u.db.Debug().Where("id_kepemilikan = ? AND nama_universitas LIKE ?", 2, query).Find(&univ).Error; err != nil {
		return nil, err
	}
	return univ, nil
}
