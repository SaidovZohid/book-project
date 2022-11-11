package main

import (
	bo "book/bookTesting"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router /book [post]
// @Summary Create book
// @Description Create book
// @Tags books
// @Accept json
// @Produce json
// @Param book body CreateOrUpdateRequest true "Book"
// @Success 200 {object} bo.Book
// @Failure 500 {object} ResponseError
func (h *handler) CreateBook(ctx *gin.Context) {
	var b bo.Book
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	blog, err := h.storage.CreateBook(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, blog)
}


// @Router /book/{id} [get]
// @Summary Get book by id
// @Description Get book by id
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} bo.Book
// @Failure 500 {object} ResponseError
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


// @Router /book/update/{id} [put]
// @Summary Update book by id
// @Description Update book by id
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param book body CreateOrUpdateRequest true "Book"
// @Success 200 {object} bo.Book
// @Failure 500 {object} ResponseError
func (h *handler) UpdateBook(ctx *gin.Context) {
	var b bo.Book
	var err error
	b.Id, err = strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Message: err.Error(),
		})
		return
	}
	err = ctx.ShouldBindJSON(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Message: err.Error(),
		})
		return
	}
	blog, err := h.storage.UpdateBook(&b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ResponseError{
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, blog)
}



// @Router /book/getbooks [get]
// @Summary GET ALL BOOKS
// @Description Get All Books
// @Tags books
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param author query string false "Author"
// @Param price query number false "Price"
// @Success 200 {object} bo.GetBooksRes
// @Failure 500 {object} ResponseError
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

// @Router /book/delete/{id} [delete]
// @Summary deletes book
// @Description Deletes book info
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} ResponOK
// @Failure 500 {object} ResponseError
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
	ctx.JSON(http.StatusOK, ResponOK{
		Message: "Succesfully deleted!",
	})
}