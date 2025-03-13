# URL/PDF to Markdown Converter

URL/PDF to Markdown Converterは、ウェブページのURLやPDFファイルからマークダウン形式のテキストを生成するシンプルなWebアプリケーションです。

## 機能

- ウェブページURLの入力からMarkdownへの変換
- PDFファイルのアップロードからMarkdownへの変換
- 生成されたMarkdownのコピーとダウンロード機能


## 技術スタック

- **バックエンド**: Go言語（Gin Webフレームワーク）
- **フロントエンド**: HTML, CSS, JavaScript
- **PDF処理**: poppler-utils (pdftotext)
- **HTML解析**: goquery
- **HTML→Markdown変換**: html-to-markdown
- **コンテナ化**: Docker, Docker Compose

## 使用方法

### Dockerを使用した実行（推奨）

1. このリポジトリをクローン
   ```bash
   git clone https://github.com/[あなたのユーザー名]/url-to-markdown-converter.git
   cd url-to-markdown-converter
   ```

2. Docker Composeでビルド＆実行
   ```bash
   docker-compose up --build
   ```

3. ブラウザで開く
   ```
   http://localhost:8080
   ```

### ローカル環境での実行

1. Go 1.16以上をインストール
2. poppler-utilsをインストール（PDF変換機能のため）
   - Ubuntu/Debian: `sudo apt-get install poppler-utils`
   - macOS: `brew install poppler`
   - Windows: [poppler for Windows](https://blog.alivate.com.au/poppler-windows/)をダウンロード

3. 依存関係のインストール
   ```bash
   go mod download
   ```

4. アプリケーションを実行
   ```bash
   go run main.go
   ```

5. ブラウザで開く
   ```
   http://localhost:8080

   ```
## デモ

![デモ](demo/demo.mov)

## ディレクトリ構成

```
.
├── handlers/         # HTTPリクエストハンドラー
├── services/         # ビジネスロジック
├── static/           # 静的ファイル
│   ├── css/          # CSSファイル
│   └── js/           # JavaScriptファイル
├── templates/        # HTMLテンプレート
├── Dockerfile        # Dockerイメージ定義
├── docker-compose.yml # Docker Compose設定
├── go.mod            # Goモジュール定義
├── go.sum            # 依存関係のチェックサム
├── main.go           # アプリケーションのエントリーポイント
└── README.md         # このファイル
```

## 開発環境のセットアップ

1. 前提条件
   - Go 1.16+
   - Git
   - Docker & Docker Compose（オプション）

2. 開発用サーバーの起動
   ```bash
   # ホットリロード機能付きで実行（air等のツールを使用）
   air
   ```



