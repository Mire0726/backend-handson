package app

import (
	"backend/app/handler"
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
	e.Use(echomiddleware.CORS())

	// ルーティングの設定
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to TodoApp!")
	})

	// POSTリクエストのルート
	e.POST("/todos", handler.HandleTodoCreate)
	e.GET("/todos", handler.HandleGetAllTodos)
	e.PUT("/todo", handler.HandleUpdateTodoByTitle)
	e.DELETE("/todos/:id", handler.HandleDeleteTodo)

	// サーバーの起動
	log.Printf("Server running on %s", addr)
	if err := e.Start(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
