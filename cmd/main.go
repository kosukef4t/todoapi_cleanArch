package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"myproject/di"
	"myproject/infrastructure/database"
	"myproject/infrastructure/router"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// ポート番号を環境変数から取得
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//echoを起動
	engine := echo.New()

	// データベースの初期化
	db, err := database.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 依存関係を初期化
	services := di.InitializeServices(db)

	// ルーティングの初期化
	router.NewRouter(engine, services.Service, services.AuthService)

	// サーバーの開始
	log.Printf("Starting server on :%s...", port)
	if err := engine.Start(":" + port); err != nil {
		log.Fatal(err)
	}

	// サーバーのシャットダウン処理
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := engine.Shutdown(ctx); err != nil {
		engine.Logger.Fatal(err)
	}
	println("stop server method")
}
