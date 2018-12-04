package controllers

import (
	"gin2018/app/hotel/ui/api/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HotelController struct{}

func (HotelController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "index",
	})
}

func (HotelController) Store(c *gin.Context) {
	// 1. Validate
	var formRequest requests.StoreRequest
	if errors := c.ShouldBind(&formRequest); errors != nil {
		formRequest.Errors(c, errors)
		return
	}

	// 2. Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "store",
	})
}

func (HotelController) Show(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":      c.Param("id"),
		"message": "show",
	})
}

func (HotelController) Update(c *gin.Context) {
	// 1. Validate
	var formRequest requests.StoreRequest
	if errors := c.ShouldBind(&formRequest); errors != nil {
		formRequest.Errors(c, errors)
		return
	}

	// 2. Response
	c.JSON(http.StatusOK, gin.H{
		"id":      c.Param("id"),
		"message": "update",
	})
}

func (HotelController) Destroy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":      c.Param("id"),
		"message": "destroy",
	})
}
