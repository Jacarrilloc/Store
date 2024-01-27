package routes

import (
    "github.com/gin-gonic/gin"
    "Store/controllers" // Asegúrate de tener tus controladores definidos
)

func SetupRoutes(router *gin.Engine) {
    // Aquí van tus rutas, por ejemplo:
    router.GET("/users", controllers.GetAllUsers)
    router.POST("/users", controllers.CreateUser)
    // ...
}
