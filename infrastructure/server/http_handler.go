package server

import (
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/raismaulana/my-gogen-rest-setup/infrastructure/util"
)

// GinHTTPHandler will define basic HTTP configuration with gracefully shutdown
type GinHTTPHandler struct {
	GracefullyShutdown
	Router *gin.Engine
}

func NewGinHTTPHandler(address string) (GinHTTPHandler, error) {

	router := gin.Default()
	router.Static("/public", "./public")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		uni := ut.New(en, en)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		util.Trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, util.Trans)
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

	}
	// CORS
	router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	// PING API
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ready")
	})

	return GinHTTPHandler{
		GracefullyShutdown: NewGracefullyShutdown(router, address),
		Router:             router,
	}, nil

}

// RunApplication is implementation of RegistryContract.RunApplication()
func (r *GinHTTPHandler) RunApplication() {
	r.RunWithGracefullyShutdown()
}
