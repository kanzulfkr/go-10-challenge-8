package handler

import (
	"C8/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h HttpServer) CreateBook(c *gin.Context) {
	var newBook model.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	res, err := h.app.CreateBook(newBook)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Book": res,
	})
}

func (h HttpServer) ListBook(c *gin.Context) {

	res, err := h.app.ListBook()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Book": res,
	})
}

func (h HttpServer) GetBook(c *gin.Context) {
	id := c.Param("id")

	res, err := h.app.GetBook(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Book": res,
	})
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var newBook model.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	res, err := h.app.UpdateBook(id, newBook)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Book": res,
	})
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	count, err := h.app.DeleteBook(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d rows affected", count),
	})
}
