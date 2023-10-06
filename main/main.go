package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhileshsirohi/SignIn-SignUp-Authentication-GOLANG/pkg/routes"
)

func main() {
	port := "9020"
	router := gin.New()

	// router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Run(":" + port)

}
