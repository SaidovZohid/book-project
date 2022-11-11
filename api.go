package main

import (
	bo "book/bookTesting"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateBook(ctx *gin.Context) {
	var b bo.Book
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	blog, err := h.storage.CreateBook(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, blog)
}

func (h *handler) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	blog, err := h.storage.GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func (h *handler) UpdateBook(ctx *gin.Context) {
	var b bo.Book
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	blog, err := h.storage.UpdateBook(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, blog)
}

func (h *handler) GetAllBook(ctx *gin.Context) {
	queryParams, err := validateParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	books, err := h.storage.GetAllBook(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func validateParams(ctx *gin.Context) (*bo.BookParam, error) {
	var (
		limit int64
		page int64 = 1
		price float64
		err error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if ctx.Query("price") != "" {
		price, err = strconv.ParseFloat(ctx.Query("price"), 64)
		if err != nil {
			return nil, err
		}
	}

	return &bo.BookParam{
		Limit: int(limit),
		Page: int(page),
		Author: ctx.Query("author"),
		Price: price,
	}, nil
}

func (h *handler) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err !=  nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = h.storage.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, "Succesfully deleted!")
}