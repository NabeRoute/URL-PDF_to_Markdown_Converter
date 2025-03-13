package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	md "github.com/JohannesKaufmann/html-to-markdown"
)

// ConvertHTMLToMarkdown 関数はURLからHTMLを取得し、Markdownに変換します
func ConvertHTMLToMarkdown(url string) (string, error) {
	// HTTPクライアントの設定
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// URLに対してGETリクエスト
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	// レスポンスコード確認
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	// goqueryでHTMLドキュメントを解析
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// html-to-markdownコンバーターを初期化（最もシンプルな形式）
	converter := md.NewConverter("", true, nil)

	// HTMLをマークダウンに変換
	html, err := doc.Html()
	if err != nil {
		return "", fmt.Errorf("failed to get HTML content: %w", err)
	}
	
	markdown, err := converter.ConvertString(html)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to markdown: %w", err)
	}

	return markdown, nil
}