package main

import (
	"Dzaakk/go-restful-api/app"
	"Dzaakk/go-restful-api/controller"
	"Dzaakk/go-restful-api/helper"
	"Dzaakk/go-restful-api/middleware"
	"Dzaakk/go-restful-api/repository"
	"Dzaakk/go-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
