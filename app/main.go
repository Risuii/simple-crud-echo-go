package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"

	"github.com/Risuii/config"
	"github.com/Risuii/internal"
)

func main() {
	cfg := config.New()

	router := mux.NewRouter()

	usecase := internal.NewUseCase()

	internal.NewHandler(router, *usecase)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.App.Port),
		Handler: router,
	}

	port := os.Getenv("PORT")

	fmt.Println("SERVER ON")
	fmt.Println("PORT :", port)
	log.Fatal(server.ListenAndServe())
}
