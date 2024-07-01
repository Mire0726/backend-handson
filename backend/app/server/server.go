package app

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

// Serve はHTTPサーバを起動します
func Serve(addr string) {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		Skipper:      echomiddleware.DefaultCORSConfig.Skipper,
		AllowOrigins: echomiddleware.DefaultCORSConfig.AllowOrigins,
		AllowMethods: echomiddleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: []string{"Content-Type", "Accept", "Origin", "X-Token", "Authorization"},
	}))

	// ルーティングの設定
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to TodoApp!")
	})

	// // POSTリクエストのルート
	// e.POST("/todos", handler.HandleTodoCreate)
	// e.GET("/todos", handler.HandleGetAllTodos)
	// e.PUT("/todo", handler.HandleUpdateTodoByTitle)
	// e.DELETE("/todos/:id", handler.HandleDeleteTodo)

	// サーバーの起動
	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("failed to start server. %+v", err)
	}
}
