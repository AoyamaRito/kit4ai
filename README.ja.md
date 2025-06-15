# Kit4AI - ASCII アート UI 仕様書作成ツール

AIがWebインターフェースを生成するために使用できる、完璧に整列したASCIIアートUI仕様書を作成するための強力なGoベースのツールです。

## デモ - 実際の出力例

### エンタープライズ ダッシュボード UI

```
+------------------------------------------------------------------------------+
| ENTERPRISE CONTROL PANEL v2.4.1                           2024-06-15 14:32:17|
+------------------------------------------------------------------------------+
| [F1]File [F2]Edit [F3]View [F4]Tools [F5]Reports [F6]Admin [ESC]Exit         |
+------------------------------------------------------------------------------+
+------------------++----------------------------------------------------------+
|NAVIGATION        ||+----------------+ +-----------------+ +-----------------+|
| [1] Dashboard    |||REVENUE METRICS | |PERFORMANCE      | |SECURITY         ||
|>[2] Analytics    |||Daily: $47,892  | |Avg Resp: 245ms  | |Threats: 0       ||
| [3] Users        |||Weekly: $312,456| |Uptime: 99.97%   | |Blocked: 127     ||
| [4] Settings     |||Monthly: $1.2M  | |Errors: 0.03%    | |Firewall: ON     ||
|                  |||Growth: +12.5%  | |Requests: 847K   | |SSL: Valid       ||
|QUICK ACTIONS     |||Target: 87%     | |Cache Hit: 94%   | |Backup: OK       ||
| [R] Refresh      ||+----------------+ +-----------------+ +-----------------+|
| [B] Backup       ||+------------------------------------+ +-----------------+|
| [M] Maintenance  |||ANALYTICS CHART - Last 7 Days       | |LIVE ACTIVITY    ||
|                  |||Revenue |#######*****:::::....      | |14:32 Login: admin|
|SYSTEM STATUS     |||Traffic |****#######****::....      | |14:31 Order #4891||
| CPU: 67%         |||Users   |:::*****########**..       | |14:30 User signup||
| RAM: 4.2/8GB     |||Errors  |.....::::*****......       | |14:29 Payment OK ||
| Online: 1,247    ||+        Mon Tue Wed Thu Fri Sat Sun-+ |14:28 Backup done||
+------------------++----------------------------------------------------------+
| Connected: DB-MAIN | Cache: REDIS-01 | Queue: 247 | Alerts: 0  F10:Settings  |
+------------------------------------------------------------------------------+
```

### スマートフォン UI

```
+----------------------------------------------------------+
| 12:34 PM               5G                     [====] 87% |
+----------------------------------------------------------+
| <-                   MESSAGES                       [+]  |
+----------------------------------------------------------+
| [1] New Message               Active: 3                  |
| [2] Contacts                  Total: 127                 |
| [3] Recent Chats              Unread: 5                  |
| [4] Settings                  Status: Online             |
+----------------------------------------------------------+
| John Doe                               2:30 PM           |
| Hey, are you free for lunch?                             |
| Jane Smith                             1:45 PM           |
| Meeting at 3 PM confirmed                                |
+----------------------------------------------------------+
|    [HOME]     [CHAT]     [CALL]     [MORE]     [USER]    |
+----------------------------------------------------------+
```

## 概要

Kit4AIは、AIが直接ASCIIアートを作成する際に発生するレイアウトのずれ問題を解決します。構造化されたキャンバスシステムと自動文字フィルタリング機能により、開発者はMarkdownドキュメントで完璧にレンダリングされる一貫した、プロフェッショナルなUI仕様書を作成できます。

## 主な機能

- **完璧な整列**: ByteCanvasシステムによりレイアウトのずれを防止
- **マルチ幅サポート**: 設定可能なキャンバス幅（60, 72, 80, 100, 120文字）
- **ASCIIフィルター**: 全角文字を自動除去してレイアウトずれを防止
- **レイヤーシステム**: Z順序による複雑なUI構成
- **Markdown対応**: ドキュメント埋め込み用に最適化された出力

## 解決される問題

### Kit4AI導入前
```
+------------------+
| 不整列なUI      |  ← 全角文字によるレイアウトずれ
| レイアウト      |
+------------------+
```

### Kit4AI導入後
```
+------------------+
| Perfect Layout   |  ← ASCIIフィルターによる完璧な整列
| Clean Design     |
+------------------+
```

## アーキテクチャ

### コアコンポーネント

- **Canvas**: 基本的なruneベースのグリッドシステム
- **ByteCanvas**: 安定したASCIIアートのための8bit処理
- **TextLayer**: 全角文字サポート（整列のため非推奨）
- **LayerSystem**: Z順序による多層レイヤー構成
- **Config System**: 柔軟な幅設定

### キャンバス設定

```go
StandardConfig    = 80x100   // レガシー互換
WideConfig        = 100x100  // モダンディスプレイ
UltraWideConfig   = 120x100  // 大型モニター
CompactConfig     = 60x80    // モバイル/狭い画面
PrintConfig       = 72x90    // A4用紙対応
```

## インストール

### 前提条件

- Go 1.19以降
- Git

### ステップ1: リポジトリをクローン

```bash
git clone https://github.com/AoyamaRito/kit4ai.git
cd kit4ai
```

### ステップ2: Goモジュールを初期化

```bash
go mod init kit4ai
go mod tidy
```

### ステップ3: インストール確認

```bash
go run main.go
```

以下のような出力が表示されれば成功です：
```
Complex Enterprise Dashboard UI created: complex_enterprise_ui.txt
```

### ステップ4: テスト実行（オプション）

```bash
go test ./pkg/canvas/...
```

## 使用方法

### 基本例

```go
package main

import (
    "fmt"
    "kit4ai/pkg/canvas"
)

func main() {
    // 設定を指定
    canvas.SetConfig(canvas.StandardConfig)
    
    // キャンバスを作成
    ui := canvas.NewByteCanvas()
    
    // フレームを描画
    ui.DrawBox(0, 0, 79, 10)
    
    // テキストを追加（全角文字は自動フィルタリング）
    ui.WriteBytesASCII(2, 2, "Hello World!")
    
    // 出力
    fmt.Println(ui.String())
}
```

### 複雑なUI例

```go
// エンタープライズダッシュボードを作成
canvas.SetConfig(canvas.StandardConfig)
dashboard := canvas.NewByteCanvas()

// タイトルバー
dashboard.DrawBox(0, 0, 79, 2)
dashboard.WriteBytesASCII(2, 1, "ENTERPRISE DASHBOARD v2.1")

// マルチパネルレイアウト
dashboard.DrawBox(0, 3, 25, 15)  // サイドバー
dashboard.DrawBox(26, 3, 79, 15) // メインコンテンツ

// 自動ASCIIフィルタリングでコンテンツを追加
dashboard.WriteBytesASCII(2, 5, "Navigation Menu")
dashboard.WriteBytesASCII(28, 5, "Analytics Data")
```

## API リファレンス

### ByteCanvasメソッド

- `NewByteCanvas()` - 現在の設定で新しいキャンバスを作成
- `DrawBox(x1, y1, x2, y2)` - 矩形フレームを描画
- `WriteBytes(x, y, text)` - 生テキストを書き込み
- `WriteBytesASCII(x, y, text)` - 全角文字フィルタリング付きで書き込み
- `FilterASCII(text)` - 文字列から全角文字を除去
- `String()` - 末尾空行除去付きで文字列に変換

### 設定メソッド

- `SetConfig(config)` - キャンバス寸法を設定
- `SetStandardWidth()` - 80文字（レガシー）
- `SetWideWidth()` - 100文字（モダン）
- `SetCompactWidth()` - 60文字（モバイル）
- `GetCurrentWidth()` - アクティブな幅を取得
- `GetConfigName()` - 設定の説明を取得

## 含まれている例

リポジトリには複数のUI例が含まれています：

- **エンタープライズダッシュボード**: 複雑なマルチパネル管理画面
- **スマートフォンUI**: ナビゲーション付きモバイルアプリレイアウト
- **銀行アプリ**: セキュリティ機能付き金融インターフェース
- **POSターミナル**: 小売店舗向け販売システム
- **病院システム**: 医療管理インターフェース

## 生成されるUI仕様書

すべての例はMarkdown互換のテキストファイルを生成します：

```
Configuration: Standard (80x100) - Legacy Compatible
Features: Multi-panel layout, real-time data, charts, logs
ASCII Filter: Enabled (all full-width characters removed)

Layout:
+------------------------------------------------------------------------------+
| ENTERPRISE CONTROL PANEL v2.4.1                           2024-06-15 14:32:17|
+------------------------------------------------------------------------------+
```

## 技術的決定

### なぜByteCanvas？
- Unicode整列問題を排除
- 一貫した8bit文字処理
- 全環境での安定した配置

### なぜASCII限定？
- 汎用的な互換性
- Markdownでのレイアウトずれ防止
- 全テキストエディターでの一貫したレンダリング
- プロフェッショナルな外観

### なぜレイヤーシステム？
- 複雑なUI構成
- Z順序管理
- モジュラー開発
- 個別コンポーネントの簡単なテスト

## ベストプラクティス

1. **テキストコンテンツには常にWriteBytesASCII()を使用**
2. **キャンバス作成前に設定を指定**
3. **ターゲットディスプレイに適した幅を使用**
4. **異なる設定でテスト**
5. **UI要素をキャンバス境界内に保持**

## ASCIIフィルターの詳細

自動ASCIIフィルターは以下を除去します：
- 日本語文字（ひらがな、カタカナ、漢字）
- 全角Unicode文字（0xFF01-0xFF5E）
- Unicode句読点（0x3000-0x303F）
- 127を超える文字（非ASCII）

保持するもの：
- 標準ASCII（0-127）
- 数字、文字、記号
- 罫線文字（フレーム用）

## ファイル構造

```
kit4ai/
├── pkg/canvas/
│   ├── canvas.go      # 基本runeキャンバス
│   ├── bytecanvas.go  # ASCII最適化キャンバス
│   ├── textlayer.go   # 全角テキストサポート
│   ├── layer.go       # レイヤー構成システム
│   └── config.go      # 設定管理
├── main.go            # 実装例
├── *.txt             # 生成されたUI仕様書
├── README.md         # 英語版ドキュメント
└── README.ja.md      # 日本語版ドキュメント（このファイル）
```

## 使用ケース

- **AI UI生成**: AIシステム用の構造化テンプレート提供
- **ドキュメント**: 技術文書へのUIモックアップ埋め込み
- **プロトタイピング**: 高速ASCIIベースインターフェース設計
- **クロスプラットフォーム**: 汎用テキストベースUI仕様書
- **レガシーシステム**: ターミナルベースインターフェース設計

## 貢献

1. リポジトリをフォーク
2. 機能ブランチを作成
3. 新機能のテストを追加
4. ASCIIフィルター互換性を確保
5. ドキュメントを更新
6. プルリクエストを送信

## ライセンス

MITライセンス - 詳細はLICENSEファイルを参照

---

*Kit4AIは、WebインターフェースプロジェクトのためのAIシステムが完璧に整列したASCIIアートUI仕様書を作成できるようにします。*