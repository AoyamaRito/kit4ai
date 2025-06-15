# Kit4AI - AI専用ASCII UI設計書生成ツール

**🤖 Claude Code専用ツール | AI間通信用 | 人間の直接使用不可**

**🇯🇵 日本語 | [🇺🇸 English README](README.md)**

**Claude Code**が自動的に完璧に整列したASCII UI設計書を生成するための専用ツールです。AIエージェントが構造化された設計ドキュメントを作成し、人間の開発者がレビュー・実装できるように設計されています。

## ⚠️ 重要: このツールについて

- **🤖 AI専用**: Claude Code AIエージェント専用設計
- **📋 設計書生成**: 実際のUIではなく、設計ドキュメントを作成
- **🚫 人間向けではない**: 人間の直接操作を想定していません
- **🔄 AI→人間ワークフロー**: Claude Code → Kit4AI → 設計仕様書 → 人間開発者

## Kit4AIの機能

Claude Codeができること:
1. **UI設計書生成**: Webインターフェース用の詳細なASCIIモックアップ作成
2. **設計案提示**: UIコンセプトの視覚的表現提供
3. **レイアウト文書化**: 開発チーム向けの構造化仕様書生成
4. **視覚的コミュニケーション**: AI提案と人間実装の橋渡し

## Claude Codeの使用方法

**注意: これらの例はClaude Codeが内部的にこのツールを使用する方法を示しています**

```bash
# Claude Codeがエンタープライズダッシュボード仕様書を生成
go run main.go --template=enterprise --width=80

# Claude Codeがモバイル UI仕様書を作成
go run main.go --template=mobile --width=60

# Claude Codeが設計ドキュメントにUI仕様を挿入
go run main.go --template=simple --insert=design_doc.md:25 --backup

# Claude Codeが利用可能オプションを確認
go run main.go --help
```

**人間の開発者の皆様**: Claude Codeから生成された仕様書を受け取ることになります。このツールを直接使用することはありません。

## Claude Code生成仕様書例

**Claude Codeが人間開発者向けに生成する仕様書の例:**

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

## 目次

- [デモ - 実際の出力例](#デモ---実際の出力例)
- [概要](#概要)
- [インストール](#インストール)
- [使用方法](#使用方法)
- [API リファレンス](#api-リファレンス)
- [含まれている例](#含まれている例)
- [ベストプラクティス](#ベストプラクティス)
- [貢献](#貢献)
- [Language / 言語](#language--言語)

## 概要

Kit4AIは、AIが直接ASCIIアートを作成する際に発生するレイアウトのずれ問題を解決します。構造化されたキャンバスシステムと自動文字フィルタリング機能により、開発者はMarkdownドキュメントで完璧にレンダリングされる一貫した、プロフェッショナルなUI仕様書を作成できます。

## 主な機能

- **コマンドラインインターフェース**: テンプレート、幅オプション、ヘルプシステムを備えた完全なCLI
- **複数テンプレート**: エンタープライズダッシュボード、モバイルインターフェース、シンプルレイアウト
- **レスポンシブ設計**: 異なるキャンバス幅（60, 72, 80, 100, 120文字）に自動適応
- **ドキュメント挿入**: 既存ファイルの指定行に直接UI挿入
- **バックアップサポート**: 既存ファイル変更時の自動バックアップ作成
- **完璧な整列**: ByteCanvasシステムによりレイアウトのずれを防止
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

### コマンドラインインターフェース

Kit4AIは主にコマンドラインツールとして様々なオプションで使用します：

```bash
# 基本使用法 - デフォルトのエンタープライズUIを生成
go run main.go

# テンプレート選択
go run main.go --template=mobile     # モバイルスマートフォンインターフェース
go run main.go --template=enterprise # 複雑なダッシュボード（デフォルト）
go run main.go --template=simple     # 基本的な2パネルレイアウト

# 幅設定
go run main.go --width=60   # コンパクト（モバイル/狭い画面）
go run main.go --width=72   # 印刷対応（A4）
go run main.go --width=80   # 標準（レガシー互換）
go run main.go --width=100  # ワイド（モダンディスプレイ）
go run main.go --width=120  # ウルトラワイド（大型モニター）

# カスタム出力ファイル
go run main.go --output=my_dashboard.txt

# ドキュメント挿入
go run main.go --template=mobile --insert=document.txt:10 --backup
```

### ドキュメント挿入機能

既存ファイルの指定行に直接UIを挿入：

```bash
# 25行目にモバイルUIを挿入（バックアップ付き）
go run main.go --template=mobile --width=60 --insert=readme.txt:25 --backup

# ドキュメント末尾にエンタープライズダッシュボードを挿入
go run main.go --template=enterprise --insert=design_doc.txt:999

# バックアップなしでシンプルレイアウトを挿入
go run main.go --template=simple --width=100 --insert=specification.md:15
```

### プログラム的使用

高度な用途では、Kit4AIをGoライブラリとしても使用できます：

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

## コマンドラインオプション

### 基本オプション

- `--template` - 生成するUIテンプレート（enterprise, mobile, simple）
- `--width` - キャンバス幅（60, 72, 80, 100, 120）
- `--output` - 出力ファイル名（デフォルト：自動生成）
- `--help` - ヘルプ情報を表示
- `--version` - バージョン情報を表示

### ドキュメント挿入

- `--insert file:line` - 既存ファイルの指定行にUIを挿入
- `--backup` - 挿入前にバックアップ（.bak）を作成

### テンプレート

- **enterprise** - ナビゲーション、メトリクス、チャート付きの複雑なダッシュボードUI
- **mobile** - メッセージレイアウト付きスマートフォンインターフェース
- **simple** - 基本的な2パネルレイアウト

### 幅オプション

- **60** - コンパクト（モバイル/狭い画面）
- **72** - 印刷対応（A4用紙互換）
- **80** - 標準（レガシーターミナル互換）
- **100** - ワイド（モダンディスプレイ）
- **120** - ウルトラワイド（大型モニター）

## API リファレンス（ライブラリ使用）

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

## AI→人間ワークフロー

**Claude CodeがKit4AIを使用して開発チームを支援する方法:**

1. **🤖 AI分析**: Claude Codeがプロジェクト要件を分析
2. **📋 仕様書生成**: Kit4AIを使用して視覚的モックアップを作成
3. **👥 人間レビュー**: 開発者が構造化された設計ドキュメントを受け取り
4. **🔧 実装**: チームが仕様書に基づいて実際のUIを構築
5. **🔄 反復**: Claude Codeがバリエーションと改良版を生成

## 開発チーム向け使用ケース

- **📋 設計ドキュメント**: プロジェクト文書での視覚的仕様書
- **🎯 要件明確化**: UIコンセプトの明確な視覚的コミュニケーション
- **🚀 高速プロトタイピング**: インターフェースアイデアの迅速な視覚化
- **📱 クロスプラットフォーム計画**: 汎用テキストベース仕様書
- **🔄 AI支援設計**: UI設計提案でのAI活用

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

**Language / 言語:**  
[🇺🇸 English README](README.md) | 🇯🇵 **日本語**

*Kit4AIは、Claude Codeが完璧に整列したASCII UI仕様書を生成し、AI分析と人間開発ワークフローを橋渡しします。*