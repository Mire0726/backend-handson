package main

import (
	app "backend/app/server"
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// コマンドライン引数の解析
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
	flag.Parse()

	// サーバーの起動
	app.Serve(addr)
}
