package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type HotelController struct {}

func (HotelController) Index(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "index",
    })
}

func (HotelController) Store(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{
        "message": "store",
    })
}

func (HotelController) Show(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "id": c.Param("id"),
        "message": "show",
    })
}

func (HotelController) Update(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "id": c.Param("id"),
        "message": "update",
    })
}

func (HotelController) Destroy(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "id": c.Param("id"),
        "message": "destroy",
    })
}
