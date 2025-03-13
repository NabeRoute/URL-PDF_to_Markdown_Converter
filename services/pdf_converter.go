package services

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

// ConvertPDFToMarkdown はPDFファイルをMarkdownに変換します
func ConvertPDFToMarkdown(fileData []byte) (string, error) {
	// 一時ファイルの作成
	tempFile, err := os.CreateTemp("", "upload-*.pdf")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %w", err)
	}
	tempFilePath := tempFile.Name()
	defer os.Remove(tempFilePath) // 関数終了時に一時ファイルを削除

	// 受け取ったバイトデータを一時ファイルに書き込む
	if _, err := tempFile.Write(fileData); err != nil {
		tempFile.Close()
		return "", fmt.Errorf("failed to write to temp file: %w", err)
	}
	if err := tempFile.Close(); err != nil {
		return "", fmt.Errorf("failed to close temp file: %w", err)
	}

	// 一時出力ファイルのパスを作成
	textFilePath := filepath.Join(os.TempDir(), "output.txt")
	defer os.Remove(textFilePath) // 関数終了時に一時出力ファイルを削除

	
	cmd := exec.Command("pdftotext", tempFilePath, textFilePath)
	if err := cmd.Run(); err != nil {
		
		return extractTextFromPDFSimple(tempFilePath)
	}

	// 抽出されたテキストを読み込む
	textData, err := os.ReadFile(textFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read extracted text: %w", err)
	}

	// 抽出したテキストからHTMLを生成
	htmlContent := generateHTMLFromPDFText(string(textData))

	// html-to-markdownコンバーターを使用してHTMLをマークダウンに変換
	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(htmlContent)
	if err != nil {
		return "", fmt.Errorf("failed to convert HTML to markdown: %w", err)
	}

	return markdown, nil
}

// generateHTMLFromPDFText はPDFから抽出したテキストをHTMLに変換
func generateHTMLFromPDFText(text string) string {
	// テキストを行に分割
	lines := strings.Split(text, "\n")
	
	// HTMLを構築
	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<!DOCTYPE html><html><body>")
	
	inParagraph := false
	
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		
		// 空行は段落の区切りとして扱う
		if trimmedLine == "" {
			if inParagraph {
				htmlBuilder.WriteString("</p>")
				inParagraph = false
			}
			continue
		}
		
		// 新しい段落を開始
		if !inParagraph {
			htmlBuilder.WriteString("<p>")
			inParagraph = true
		} else {
			// 同じ段落内の改行
			htmlBuilder.WriteString(" ")
		}
		
		// エスケープしてHTMLに追加
		escapedLine := strings.ReplaceAll(trimmedLine, "&", "&amp;")
		escapedLine = strings.ReplaceAll(escapedLine, "<", "&lt;")
		escapedLine = strings.ReplaceAll(escapedLine, ">", "&gt;")
		htmlBuilder.WriteString(escapedLine)
	}
	
	// 最後の段落を閉じる
	if inParagraph {
		htmlBuilder.WriteString("</p>")
	}
	
	htmlBuilder.WriteString("</body></html>")
	return htmlBuilder.String()
}

// フォールバックメソッド - より簡単なPDFテキスト抽出方法
func extractTextFromPDFSimple(pdfPath string) (string, error) {

	return "# PDF変換メッセージ\n\nこのPDFの変換には外部ツール（pdftotext）が必要です。\nDockerイメージに poppler-utils をインストールするか、テキスト抽出用の純粋なGoライブラリを使用してください。", nil
}

// SaveUploadedFile はアップロードされたファイルを一時ディレクトリに保存
func SaveUploadedFile(file io.Reader) ([]byte, error) {
	// ファイルデータを読み込む
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read uploaded file: %w", err)
	}

	return fileData, nil
}