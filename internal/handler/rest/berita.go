package rest

import (
	"errors"
	"future-path/model"
	"future-path/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *Rest) CreateBerita(ctx *gin.Context) {
	var beritaReq model.CreateBerita

	if err := ctx.ShouldBindJSON(&beritaReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid Request", err)
		return
	}

	berita, err := r.service.BeritaService.CreateBerita(&beritaReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create news", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "News created", berita)

}

func (r *Rest) GetBeritaSingkat(ctx *gin.Context) {
	berita, err := r.service.BeritaService.GetBeritaSingkat()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get short news", err)
		return
	}

	var beritaResponse []model.GetBerita
	for _, b := range berita {
		beritaResponse = append(beritaResponse, model.GetBerita{
			Judul_Berita: b.Judul_Berita,
			Isi_Berita:   b.Isi_Berita,
		})
	}

	response.Success(ctx, http.StatusOK, "Short new retrieved", beritaResponse)

}

func (r *Rest) GetBeritaFull(ctx *gin.Context) {
	idParam := ctx.Param("id_berita")

	if idParam == "" {
		response.Error(ctx, http.StatusBadRequest, "Invalid news ID", errors.New("ID cannot be empty"))
	}

	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid news ID", err)
		return
	}

	berita, err := r.service.BeritaService.GetBeritaFull(uint(id))

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get full news", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Full news retrieved", berita)
}
