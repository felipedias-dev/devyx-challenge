package main

import (
	"database/sql"
	"fmt"

	_ "github.com/devyx-tech/felipe-challegend/api/docs"
	"github.com/devyx-tech/felipe-challegend/configs"
	"github.com/devyx-tech/felipe-challegend/internal/infra/database"
	"github.com/devyx-tech/felipe-challegend/internal/infra/web/handlers"
	"github.com/devyx-tech/felipe-challegend/internal/infra/web/webserver"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Devyx Challenge API
// @version 1.0
// @description This is a doc for Products server.

// @contact.name Felipe Dias
// @contact.email felipedias.dev@gmail.com

// @host localhost:3333
// @BasePath /
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productRepository := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepository)

	webserver := webserver.NewWebServer(configs.WebServerPort)

	webserver.AddHandler("docs", httpSwagger.WrapHandler)
	webserver.AddHandler("createProduct", productHandler.Create)
	webserver.AddHandler("listProducts", productHandler.List)
	webserver.AddHandler("getProduct", productHandler.GetByID)
	webserver.AddHandler("updateProduct", productHandler.Update)
	webserver.AddHandler("deleteProduct", productHandler.Delete)

	webserver.RegisterHandlers()

	webserver.Start()
}
