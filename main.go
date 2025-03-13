package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"url-to-markdown/handlers"
)

func main() {
	// Ginルーターを初期化
	router := gin.Default()

	// Add logging middleware
	router.Use(func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log request details
		latency := time.Since(start)
		log.Printf("[%s] %s %s - %v", 
			c.Request.Method, 
			c.Request.URL.Path, 
			c.ClientIP(),
			latency,
		)
	})

	// 静的ファイルの提供 - add file server debugging
	router.Static("/static", "./static")
	router.StaticFS("/debug-static", http.Dir("./static"))
	
	// Serve static files with additional header
	router.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Next()
	})

	// HTMLテンプレートのロード
	router.LoadHTMLGlob("templates/*")

	// ルートエンドポイント
	router.GET("/", func(c *gin.Context) {
		log.Println("Serving index.html")
		c.HTML(200, "index.html", gin.H{
			"title": "URL to Markdown Converter",
			"timestamp": time.Now().Unix(),
		})
	})

	// API エンドポイント
	router.POST("/convert", handlers.ConvertURLToMarkdown)
	
	// PDFアップロード＆変換エンドポイント
	router.POST("/convert-pdf", handlers.ConvertPDFToMarkdown)

	// サーバー起動
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}