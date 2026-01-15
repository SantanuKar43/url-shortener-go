package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/SantanuKar43/url-shortener-go/uss"
)

func main() {
	uss.Init()
	startHttpServer()
}

func startHttpServer() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/:shortId", uss.HandleResolve)
	e.GET("/preview/:shortId", uss.HandlePreview)
	e.PUT("/", uss.HandleCreate)
	e.DELETE("/:shortId", uss.HandleDelete)
	e.Logger.Fatal(e.Start(":8000"))
}