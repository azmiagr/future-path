package rest

import (
	"fmt"
	"future-path/internal/service"
	"future-path/pkg/middleware"
	"future-path/pkg/response"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	router     *gin.Engine
	service    *service.Service
	middleware middleware.Interface
}

func NewRest(service *service.Service, middleware middleware.Interface) *Rest {
	return &Rest{
		router:     gin.Default(),
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) MountEndpoint() {
	routerGroup := r.router.Group("/future-path")
	routerGroup.Use(r.middleware.Timeout())
	routerGroup.GET("/testing", testTimeout)

	auth := routerGroup.Group("/auth")
	auth.POST("/register", r.Register)
	auth.POST("/login", r.Login)

	user := routerGroup.Group("/user")
	user.Use(r.middleware.AuthenticateUser)
	user.GET("/berita", r.GetBeritaSingkat)
	user.GET("/full-news", r.GetBeritaFull)
	user.GET("/cari-sekolah/negeri", r.GetSekolahNegeri)
	user.GET("/cari-sekolah/swasta", r.GetSekolahSwasta)
	user.GET("cari-universitas/negeri", r.GetUnivNegeri)
	user.GET("cari-universitas/swasta", r.GetUnivSwasta)

	adminAuth := routerGroup.Group("/admin/auth")
	adminAuth.POST("/register", r.RegisterAdmin)
	adminAuth.POST("/login", r.LoginAdmin)

	admin := routerGroup.Group("/admin")
	admin.Use(r.middleware.AuthenticateAdmin, r.middleware.OnlyAdmin)
	admin.POST("/create-berita", r.CreateBerita)
	admin.PATCH("/update-berita/:id_berita", r.UpdateBerita)
	admin.DELETE("/delete-berita/:id_berita", r.DeleteBerita)
}

func testTimeout(ctx *gin.Context) {
	time.Sleep(3 * time.Second)

	response.Success(ctx, http.StatusOK, "success", nil)
}

func (r *Rest) Run() {
	addr := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")

	r.router.Run(fmt.Sprintf("%s:%s", addr, port))
}
