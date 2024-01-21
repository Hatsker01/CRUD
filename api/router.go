package api

import (
	_ "github.com/CRUD/api/docs"
	v1 "github.com/CRUD/api/handler/v1"
	"github.com/CRUD/config"
	"github.com/CRUD/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Db     *sqlx.DB
	Conf   config.Config
	Logger logger.Logger
}

func New(option Option) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(),
		gin.Recovery(),
	)
	handlerV1 := v1.New(&v1.HandlerV1Config{
		Db:     option.Db,
		Logger: option.Logger,
		Cfg:    option.Conf,
	})
	api := router.Group("/v1")

	// user
	api.POST("/user", handlerV1.CreateUser)
	api.PUT("/user", handlerV1.UpdateUser)
	api.GET("/user", handlerV1.GetAllUsers)
	api.GET("/user/:id", handlerV1.GetSingleUser)
	api.DELETE("/user/:id", handlerV1.DeleteUser)
	api.GET("/user/country/:id", handlerV1.GetUserWithCountry)

	// country
	api.POST("/country", handlerV1.CreateCountry)
	api.GET("/country/:id", handlerV1.GetSingleCountry)
	api.GET("/country/user/:id", handlerV1.GetUserCountries)
	api.PUT("/country", handlerV1.UpdateCountry)
	api.DELETE("/country/:id", handlerV1.DeleteCountry)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
