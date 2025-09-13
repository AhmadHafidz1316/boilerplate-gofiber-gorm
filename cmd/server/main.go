package main

import (
	"log"

	_ "github.com/AhmadHafidz1316/go-fiber-gorm/docs" // 👉 import docs agar swagger bisa baca
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"

	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/config"
	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/db"
	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/repository"
	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/service"
	h "github.com/AhmadHafidz1316/go-fiber-gorm/internal/transport/http/handler"
	"github.com/AhmadHafidz1316/go-fiber-gorm/internal/transport/http/router"
)

// @title Go Fiber + GORM API
// @version 1.0
// @description This is a sample CRUD API with Fiber v2 + GORM + Swagger
// @host localhost:3000
// @BasePath /api/v1
func main() {
	cfg := config.Load()

	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Automigrate(database); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewCustomerRepository(database)
	svc := service.NewCustomerService(repo)
	handler := h.NewCustomerHandler(svc)

	app := fiber.New()
	router.Register(app, handler)

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Printf("listening on :%s", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
