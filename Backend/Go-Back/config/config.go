package config

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "os"
    "log"
)

var DB *gorm.DB

func ConnectDatabase() error {
    dsn := os.Getenv("DATABASE_URL") // Deberías configurar esta variable en tus variables de entorno
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    // Automigraciones para tus modelos, por ejemplo: database.AutoMigrate(&models.User{})
    // ...

    DB = database
    return nil
}
