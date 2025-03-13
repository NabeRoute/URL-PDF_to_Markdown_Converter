package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"url-to-markdown/services"
)

// URLRequest はURLリクエストの構造体
type URLRequest struct {
	URL string `json:"url" binding:"required"`
}

// ConvertURLToMarkdown はURLをMarkdownに変換するハンドラー
func ConvertURLToMarkdown(c *gin.Context) {
	var request URLRequest

	// リクエストをバインド
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format or missing URL",
		})
		return
	}

	// URLの検証
	if request.URL == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL cannot be empty",
		})
		return
	}

	// HTML→Markdown変換サービスを呼び出し
	markdown, err := services.ConvertHTMLToMarkdown(request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"markdown": markdown,
	})
}

// ConvertPDFToMarkdown はアップロードされたPDFをMarkdownに変換するハンドラー
func ConvertPDFToMarkdown(c *gin.Context) {
	// PDFファイルを取得
	file, err := c.FormFile("pdfFile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "PDF file is required",
		})
		return
	}

	// ファイルがPDFかどうかの簡易チェック
	if filepath.Ext(file.Filename) != ".pdf" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Uploaded file must be a PDF",
		})
		return
	}

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open uploaded file",
		})
		return
	}
	defer src.Close()

	// アップロードされたファイルを保存
	fileData, err := services.SaveUploadedFile(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// PDFをMarkdownに変換
	markdown, err := services.ConvertPDFToMarkdown(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"markdown": markdown,
	})
}