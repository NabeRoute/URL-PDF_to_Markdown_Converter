// Set the script loaded flag at the very beginning
window.scriptLoaded = true;
console.log('app.js scriptLoaded flag set to true');
console.log('app.js starting to load');

document.addEventListener('DOMContentLoaded', () => {
    console.log('DOM fully loaded and app.js is running');
    
    // 要素の参照を取得
    const urlModeBtn = document.getElementById('url-mode-btn');
    const pdfModeBtn = document.getElementById('pdf-mode-btn');
    const urlContent = document.getElementById('url-content');
    const pdfContent = document.getElementById('pdf-content');
    const backToModeBtnUrl = document.getElementById('back-to-mode-btn-url');
    const backToModeBtnPdf = document.getElementById('back-to-mode-btn-pdf');
    const urlInput = document.getElementById('url-input');
    const pdfInput = document.getElementById('pdf-input');
    const convertBtn = document.getElementById('convert-btn');
    const convertPdfBtn = document.getElementById('convert-pdf-btn');
    const copyBtn = document.getElementById('copy-btn');
    const downloadBtn = document.getElementById('download-btn');
    const markdownResult = document.getElementById('markdown-result');
    const resultContainer = document.querySelector('.result-container');
    const loadingElement = document.getElementById('loading');
    const errorMessage = document.getElementById('error-message');
    const modeSelection = document.getElementById('mode-selection');
    
    // デバッグ用：各要素が見つかったかを確認
    console.log('Elements found:', {
        urlModeBtn: !!urlModeBtn,
        pdfModeBtn: !!pdfModeBtn,
        urlContent: !!urlContent,
        pdfContent: !!pdfContent,
        backToModeBtnUrl: !!backToModeBtnUrl,
        backToModeBtnPdf: !!backToModeBtnPdf,
        urlInput: !!urlInput,
        pdfInput: !!pdfInput,
        convertBtn: !!convertBtn,
        convertPdfBtn: !!convertPdfBtn,
        copyBtn: !!copyBtn,
        downloadBtn: !!downloadBtn,
        markdownResult: !!markdownResult,
        resultContainer: !!resultContainer,
        loadingElement: !!loadingElement,
        errorMessage: !!errorMessage,
        modeSelection: !!modeSelection
    });
    
    // モード選択関数 - 選択したモードだけを表示する
    function showMode(mode) {
        // すべてのコンテンツを非表示
        urlContent.classList.add('hidden');
        pdfContent.classList.add('hidden');
        
        // モード選択画面を非表示
        modeSelection.classList.add('hidden');
        
        // 選択されたモードのみ表示
        mode.classList.remove('hidden');
    }
    
    // モード選択ボタンのイベントリスナー
    urlModeBtn.addEventListener('click', () => {
        console.log('URL mode selected');
        showMode(urlContent);
    });
    
    pdfModeBtn.addEventListener('click', () => {
        console.log('PDF mode selected');
        showMode(pdfContent);
    });
    
    // 戻るボタンのイベントリスナー
    backToModeBtnUrl.addEventListener('click', () => {
        resetToModeSelection();
    });
    
    backToModeBtnPdf.addEventListener('click', () => {
        resetToModeSelection();
    });

    // 強化されたエラー表示関数
    function showError(message) {
        console.error('Error:', message);
        
        try {
            if (errorMessage) {
                const pElement = errorMessage.querySelector('p');
                if (pElement) {
                    pElement.textContent = message;
                } else {
                    // p要素がない場合、直接テキストを設定
                    errorMessage.textContent = message;
                }
                errorMessage.classList.remove('hidden');
            } else {
                // エラーメッセージ要素がない場合はアラートを表示
                alert('エラー: ' + message);
            }
        } catch (e) {
            console.error('Error displaying error message:', e);
            alert('エラー: ' + message);
        }
    }

    function hideError() {
        if (errorMessage) {
            errorMessage.classList.add('hidden');
        }
    }

    function showLoading() {
        if (loadingElement) {
            loadingElement.classList.remove('hidden');
        }
    }

    function hideLoading() {
        if (loadingElement) {
            loadingElement.classList.add('hidden');
        }
    }

    function showResults() {
        if (resultContainer) {
            resultContainer.classList.remove('hidden');
        }
    }

    function hideResults() {
        if (resultContainer) {
            resultContainer.classList.add('hidden');
        }
    }

    // モード選択に戻る関数
    function resetToModeSelection() {
        // コンテンツを非表示
        urlContent.classList.add('hidden');
        pdfContent.classList.add('hidden');
        // 結果とエラーをリセット
        hideResults();
        hideError();
        // 入力をリセット
        if (urlInput) urlInput.value = '';
        if (pdfInput) pdfInput.value = '';
        // モード選択を表示
        modeSelection.classList.remove('hidden');
    }

    // URLのバリデーション
    function isValidURL(string) {
        try {
            new URL(string);
            return true;
        } catch (_) {
            return false;
        }
    }

    // ドメイン名を抽出
    function extractDomain(url) {
        try {
            const hostname = new URL(url).hostname;
            return hostname.replace(/^www\./, '');
        } catch (_) {
            return 'webpage';
        }
    }

    // 変換ボタンのイベントリスナー
    if (convertBtn) {
        console.log('Setting up URL convert button click handler');
        convertBtn.addEventListener('click', async () => {
            console.log('URL Convert button clicked');
            const url = urlInput ? urlInput.value.trim() : '';
            
            // 入力検証
            if (!url) {
                showError('URLを入力してください。');
                return;
            }

            if (!isValidURL(url)) {
                showError('有効なURLを入力してください。');
                return;
            }

            // 変換処理開始
            try {
                hideError();
                showLoading();
                hideResults();
                
                console.log('Sending fetch request to /convert with URL:', url);
                const response = await fetch('/convert', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ url }),
                });

                // レスポンス処理
                const data = await response.json();
                hideLoading();

                if (!response.ok) {
                    showError(data.error || 'エラーが発生しました。');
                    return;
                }

                // 結果表示
                if (markdownResult) {
                    markdownResult.textContent = data.markdown;
                }
                showResults();
            } catch (error) {
                hideLoading();
                showError('サーバーとの通信中にエラーが発生しました。');
                console.error('Error:', error);
            }
        });
    } else {
        console.error('Convert button not found in DOM');
    }
    
    // PDF変換ボタンのイベントリスナー
    if (convertPdfBtn) {
        console.log('Setting up PDF convert button click handler');
        convertPdfBtn.addEventListener('click', async () => {
            console.log('PDF Convert button clicked');
            
            // PDFファイルの検証
            if (!pdfInput || !pdfInput.files || pdfInput.files.length === 0) {
                showError('PDFファイルを選択してください。');
                return;
            }
            
            const file = pdfInput.files[0];
            if (file.type !== 'application/pdf') {
                showError('選択されたファイルはPDFではありません。');
                return;
            }
            
            // FormDataの作成
            const formData = new FormData();
            formData.append('pdfFile', file);
            
            // 変換処理開始
            try {
                hideError();
                showLoading();
                hideResults();
                
                console.log('Sending fetch request to /convert-pdf');
                const response = await fetch('/convert-pdf', {
                    method: 'POST',
                    body: formData,
                });
                
                // レスポンス処理
                const data = await response.json();
                hideLoading();
                
                if (!response.ok) {
                    showError(data.error || 'エラーが発生しました。');
                    return;
                }
                
                // 結果表示
                if (markdownResult) {
                    markdownResult.textContent = data.markdown;
                }
                showResults();
            } catch (error) {
                hideLoading();
                showError('サーバーとの通信中にエラーが発生しました。');
                console.error('Error:', error);
            }
        });
    } else {
        console.error('PDF convert button not found in DOM');
    }

    // クリップボードにコピーボタン
    if (copyBtn && markdownResult) {
        copyBtn.addEventListener('click', () => {
            if (!markdownResult.textContent) return;
            
            navigator.clipboard.writeText(markdownResult.textContent)
                .then(() => {
                    const originalText = copyBtn.textContent;
                    copyBtn.textContent = 'コピーしました！';
                    setTimeout(() => {
                        copyBtn.textContent = originalText;
                    }, 2000);
                })
                .catch(err => {
                    console.error('クリップボードへのコピーに失敗しました:', err);
                    showError('クリップボードへのコピーに失敗しました。');
                });
        });
    }

    // Markdownファイルをダウンロードボタン
    if (downloadBtn && markdownResult) {
        downloadBtn.addEventListener('click', () => {
            if (!markdownResult.textContent) return;

            const blob = new Blob([markdownResult.textContent], { type: 'text/markdown' });
            const url = URL.createObjectURL(blob);
            const a = document.createElement('a');
            
            // 現在の状態を確認してファイル名を作成
            let filename = 'converted-document';
            
            if (!urlContent.classList.contains('hidden') && urlInput) {
                // URLモードがアクティブな場合はURLから
                const domain = extractDomain(urlInput.value);
                filename = domain;
            } else if (!pdfContent.classList.contains('hidden') && pdfInput && pdfInput.files && pdfInput.files.length > 0) {
                // PDFモードがアクティブな場合はPDFファイル名から
                const pdfFilename = pdfInput.files[0].name.replace('.pdf', '');
                filename = `${pdfFilename}`;
            }
            
            const timestamp = new Date().toISOString().replace(/[:.]/g, '-');
            a.download = `${filename}-${timestamp}.md`;
            
            a.href = url;
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
            URL.revokeObjectURL(url);
        });
    }
});

console.log('app.js fully loaded');