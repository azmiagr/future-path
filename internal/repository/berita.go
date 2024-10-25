package repository

import (
	"future-path/entity"

	"gorm.io/gorm"
)

type IBeritaRepository interface {
	CreateBerita(berita *entity.Berita) (*entity.Berita, error)
	GetBeritaSingkat() ([]entity.Berita, error)
	GetBeritaFull(id uint) (*entity.Berita, error)
}

type BeritaRepository struct {
	db *gorm.DB
}

func NewBeritaRepository(db *gorm.DB) IBeritaRepository {
	return &BeritaRepository{db}
}

func (br *BeritaRepository) CreateBerita(berita *entity.Berita) (*entity.Berita, error) {
	if err := br.db.Debug().Create(berita).Error; err != nil {
		return nil, err
	}

	return berita, nil
}

func (br *BeritaRepository) GetBeritaSingkat() ([]entity.Berita, error) {
	var berita []entity.Berita

	if err := br.db.Debug().Select("id_berita, judul_berita, LEFT(isi_berita, 100) as isi_berita").Order("tanggal desc").Limit(5).Find(&berita).Error; err != nil {
		return nil, err
	}
	return berita, nil
}

func (br *BeritaRepository) GetBeritaFull(id uint) (*entity.Berita, error) {
	var berita entity.Berita
	if err := br.db.Debug().First(&berita, id).Error; err != nil {
		return nil, err
	}

	return &berita, nil
}
