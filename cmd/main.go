package main

import (
	"log"
	"net/http"

	"github.com/b3nhard/car-rental/config"
	"github.com/b3nhard/car-rental/internal/database"
	"github.com/b3nhard/car-rental/internal/models"
	"github.com/b3nhard/car-rental/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	c := config.NewConfig()
	db := database.InitDb(c.DB)
	db.AutoMigrate(models.User{})
	s := server.NewServer(db)
	s.MountRoutes()
	log.Fatal(http.ListenAndServe(c.Server.Addr, s.Router))

}
