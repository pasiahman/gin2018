package main

import (
	"gin2018/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Api{}.Include(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
