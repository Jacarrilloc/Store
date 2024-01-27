package main

import (
    "log"
    "os"
    "github.com/gin-gonic/gin"
    "Store/config"
    "Store/routes"
)

func main() {
    // Establecer el modo de Gin (por ejemplo, a "release" en producción)
    gin.SetMode(gin.DebugMode)

    // Conectar a la base de datos
    if err := config.ConnectDatabase(); err != nil {
        log.Fatal("No se pudo conectar a la base de datos: ", err)
        return
    }

    // Crear una instancia de Gin
    router := gin.Default()

    // Configurar las rutas
    routes.SetupRoutes(router)

    // Iniciar el servidor
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Puerto por defecto
    }
    router.Run(":" + port)
}
