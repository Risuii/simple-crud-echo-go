package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Risuii/config"
	"github.com/Risuii/internal"
)

func main() {
	cfg := config.New()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	usecase := internal.NewUseCase()
	handler := internal.Handler{
		UseCases: *usecase,
	}

	e.GET("/", handler.Home)
	e.GET("/palindrome", handler.Palindrome)
	e.GET("/languages/Nomor-3", handler.Languages)
	e.GET("/languages", handler.GetAllLanguages)
	e.POST("/language", handler.AddLanguage)
	e.GET("/language/:id", handler.GetLanguage)
	e.PATCH("/language/:id", handler.UpdateLanguage)
	e.DELETE("/language/:id", handler.DeleteLanguage)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.App.Port),
		Handler: e,
	}

	port := os.Getenv("PORT")

	fmt.Println("SERVER ON")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
	log.Fatal(server.ListenAndServe())
}
