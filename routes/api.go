package routes

import (
	HotelUIApiControllers "gin2018/app/hotel/ui/api/controllers"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func (Api) Include(router *gin.Engine) {
	router.GET("/api/hotel", HotelUIApiControllers.HotelController{}.Index)
	router.POST("/api/hotel", HotelUIApiControllers.HotelController{}.Store)
	router.GET("/api/hotel/:id", HotelUIApiControllers.HotelController{}.Show)
	router.PUT("/api/hotel/:id", HotelUIApiControllers.HotelController{}.Update)
	router.DELETE("/api/hotel/:id", HotelUIApiControllers.HotelController{}.Destroy)
}
