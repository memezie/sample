package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	custommiddleware "example.com/mod/app/customMiddleware"
	"example.com/mod/route"
	"github.com/labstack/echo/v4/middleware"
)

const ()

func main() {
	// Setup
	e := route.Init()

	// =====middlewareの利用=====
	// panic時にサーバーが落ちないように
	e.Use(middleware.Recover())

	// ログ
	e.Use(custommiddleware.LoggerWithConfig())

	// ベーシック認証
	e.Use(custommiddleware.BasicAuth())

	// =========================

	e.Logger.Fatal(e.Start(":1323"))

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	time.Sleep(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
