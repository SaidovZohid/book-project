package main

import (
	"github.com/gin-gonic/gin"
	bo "book/bookTesting"
)

type handler struct {
	storage *bo.DBManager
}

func NewServer(storage *bo.DBManager) *gin.Engine {
	r := gin.Default()

	h := handler{
		storage: storage,
	}
	
	r.GET("/book/:id", h.GetBook)
	r.POST("/book", h.CreateBook)
	r.DELETE("/book/delete/:id", h.DeleteBook)
	r.GET("/book/getbooks", h.GetAllBook)
	r.POST("/book/update", h.UpdateBook)
	return r
}