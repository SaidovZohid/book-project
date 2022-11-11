package main

import (
	bo "book/bookTesting"

	"github.com/gin-gonic/gin"

	_ "book/docs" // for swagger

	swaggerFiles "github.com/swaggo/files"     // gin-swagger
	ginSwagger "github.com/swaggo/gin-swagger" // swagger embed files
)

type handler struct {
	storage *bo.DBManager
}

// @title           Swagger for book api
// @version         1.0
// @description     This is a book service api.
// @host      localhost:8080
func NewServer(storage *bo.DBManager) *gin.Engine {
	r := gin.Default()

	h := handler{
		storage: storage,
	}

	r.GET("/book/:id", h.GetBook)
	r.POST("/book", h.CreateBook)
	r.DELETE("/book/delete/:id", h.DeleteBook)
	r.GET("/book/getbooks", h.GetAllBook)
	r.PUT("/book/update/:id", h.UpdateBook)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
