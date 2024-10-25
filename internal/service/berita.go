package service

import (
	"future-path/entity"
	"future-path/internal/repository"
	"future-path/model"
)

type IBeritaService interface {
	CreateBerita(beritaReq *model.CreateBerita) (*entity.Berita, error)
	GetBeritaSingkat() ([]entity.Berita, error)
	GetBeritaFull(id uint) (*entity.Berita, error)
}

type BeritaService struct {
	BeritaRepository repository.IBeritaRepository
}

func NewBeritaService(BeritaRepository repository.IBeritaRepository) IBeritaService {
	return &BeritaService{BeritaRepository}
}

func (bs *BeritaService) CreateBerita(beritaReq *model.CreateBerita) (*entity.Berita, error) {
	berita := &entity.Berita{
		Judul_Berita: beritaReq.Judul_Berita,
		Isi_Berita:   beritaReq.Isi_Berita,
	}

	berita, err := bs.BeritaRepository.CreateBerita(berita)
	if err != nil {
		return nil, err
	}
	return berita, nil
}

func (bs *BeritaService) GetBeritaSingkat() ([]entity.Berita, error) {
	return bs.BeritaRepository.GetBeritaSingkat()
}

func (bs *BeritaService) GetBeritaFull(id uint) (*entity.Berita, error) {
	return bs.BeritaRepository.GetBeritaFull(id)
}
