package rest

import (
	"fmt"
	"future-path/internal/service"
	"os"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router  *gin.Engine
	service *service.Service
}

func NewRest(service *service.Service) *Rest {
	return &Rest{
		router:  gin.Default(),
		service: service,
	}
}

func (r *Rest) MountEndpoint() {
	routerGroup := r.router.Group("/future-path")
	berita := routerGroup.Group("/berita")
	berita.POST("/create", r.CreateBerita)
	berita.GET("/short-news", r.GetBeritaSingkat)
	berita.GET("/full-news/:id", r.GetBeritaFull)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
