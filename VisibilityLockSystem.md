# VisibilityLockSystem - Fake Gods TRPG 視認ロックシステム

## 概要

VisibilityLockSystemは、Fake Gods TRPGにおける台詞の可視性管理と割り込み制御を行うReactコンポーネントです。プレイヤーが一度見た台詞はキャンセル不可能になり、未表示の台詞のみがキャンセル可能という「視認ロック」機能を実装しています。

## 主な機能

### 🔒 視認ロック機能
- **表示済み台詞**: 一度プレイヤーに表示されると自動的にロックされ、キャンセル不可能
- **未表示台詞**: まだ表示されていない台詞はキャンセル可能
- **ストーリー一貫性**: 確定した台詞は物語に組み込まれ、一貫性が保たれる

### ⚡ 割り込みシステム
- **2秒間の割り込み窓口**: 台詞表示直前に2秒間の割り込み可能時間を提供
- **カウントダウンタイマー**: リアルタイムで残り時間を表示
- **3つの選択肢**: キャンセル、続行、プレイヤー緊急発言

### 🎮 ユーザーインターフェース
- **パステルカラー**: 既存のQueueChatコンポーネントと統一されたデザイン
- **レスポンシブ**: デスクトップ、タブレット、モバイル対応
- **リアルタイム更新**: 2秒間隔でキュー状態を自動更新

## ファイル構成

```
/Users/AoyamaRito/fake_gods/kit4ai/
├── VisibilityLockSystem.tsx           # メインコンポーネント
├── hooks/useVisibilityLock.ts         # 状態管理フック
├── types/tailwind-pastel.d.ts         # Tailwind型定義
├── examples/
│   └── VisibilityLockIntegration.tsx  # 統合例
└── VisibilityLockSystem.md            # このドキュメント
```

## 使用方法

### 基本的な使用例

```tsx
import React from 'react';
import VisibilityLockSystem from './VisibilityLockSystem';

function App() {
  const handlePlayerAction = (action: string, data?: any) => {
    console.log('プレイヤーアクション:', action, data);
  };

  const handleInterrupt = (itemId: string, action: 'cancel' | 'continue' | 'player_action') => {
    console.log('割り込み処理:', itemId, action);
  };

  return (
    <VisibilityLockSystem
      apiEndpoint="http://localhost:3001/api"
      onPlayerAction={handlePlayerAction}
      onInterrupt={handleInterrupt}
    />
  );
}
```

### フル統合例

完全な統合例は `/examples/VisibilityLockIntegration.tsx` を参照してください。

## API エンドポイント仕様

VisibilityLockSystemは以下のAPIエンドポイントと連携します：

### GET /api/queue
キューアイテムの一覧を取得

**レスポンス例:**
```json
[
  {
    "id": "item-1",
    "speaker": "goblin",
    "content": "ガアア！邪魔するな！",
    "timestamp": "2024-06-22T10:30:00Z",
    "status": "queued",
    "isVisible": false,
    "canCancel": true,
    "displayOrder": 1
  }
]
```

### POST /api/queue
新しいアイテムをキューに追加

**リクエスト:**
```json
{
  "speaker": "player",
  "content": "剣を抜いて警戒する",
  "timestamp": "2024-06-22T10:31:00Z"
}
```

### POST /api/queue/:id/visible
アイテムを表示済みとしてマーク

### POST /api/queue/:id/interrupt
割り込み処理を実行

**リクエスト:**
```json
{
  "action": "cancel|continue|player_action",
  "playerContent": "緊急発言内容"
}
```

### DELETE /api/queue/:id
アイテムをキューから削除

## データ構造

### QueueItem インターフェース

```typescript
interface QueueItem {
  id: string;
  speaker: string;              // 話者名
  content: string;              // 台詞内容
  speakerIcon: string;          // 話者アイコン
  timestamp: Date;              // 作成日時
  status: 'generating' | 'queued' | 'displaying' | 'completed';
  isVisible: boolean;           // 視認状態
  canCancel: boolean;           // キャンセル可能フラグ
  displayOrder: number;         // 表示順序
}
```

### InterruptState インターフェース

```typescript
interface InterruptState {
  isActive: boolean;            // 割り込み窓口の有効状態
  timeRemaining: number;        // 残り時間（秒）
  targetItemId: string | null;  // 対象アイテムID
  countdownTimer: NodeJS.Timeout | null; // タイマー参照
}
```

## カスタムフック: useVisibilityLock

状態管理とAPI通信を担当するカスタムフックです。

### 提供される機能

```typescript
const {
  // 状態
  queueItems,              // キューアイテム一覧
  currentDisplayItem,      // 現在表示中のアイテム
  interruptState,          // 割り込み状態
  isLoading,              // ローディング状態
  error,                  // エラー状態
  sessionCost,            // セッション費用
  queueStats,             // キュー統計
  unseenCancelableItems,  // 未表示でキャンセル可能なアイテム
  
  // アクション
  markItemAsVisible,      // アイテムを表示済みにマーク
  handleInterruptAction,  // 割り込みアクション処理
  cancelItem,             // アイテムキャンセル
  addToQueue,             // キューに追加
  loadQueue,              // キュー再読み込み
  
  // ユーティリティ
  getCountdownProgress,   // カウントダウン進捗取得
} = useVisibilityLock(apiEndpoint);
```

## 視認ロックのルール

1. **✅ 表示済み台詞**: 🔒確定 → 物語に組み込み済み
2. **⏳ 未表示台詞**: 🛑キャンセル可能 → 割り込みで破棄
3. **🔄 表示直前**: ⚡最後のチャンス → 2秒以内に決断
4. **👁️ 視認瞬間**: 🔒即座に確定 → 取り消し不可
5. **📝 再生成**: 確定台詞以降から → 一貫性保持
6. **⚡ 緊急度**: 表示直前ほど重要 → 物語の分岐点

## スタイリング

### Tailwindクラス

パステルカラーを使用したTailwind CSSクラスを使用：

- **背景色**: `bg-pastel-*` （green, blue, purple, pink等）
- **ボーダー**: `border-pastel-*`
- **グラデーション**: `from-pastel-* to-pastel-*`
- **丸角**: `rounded-xl`
- **影**: `shadow-lg`

### カスタムカラー定義

`types/tailwind-pastel.d.ts` に定義されたパステルカラー：

```typescript
export const PASTEL_COLORS = {
  'pastel-pink': '#FFE4E6',
  'pastel-rose': '#FFE4E1',
  'pastel-purple': '#F3E8FF',
  'pastel-blue': '#E0F2FE',
  'pastel-sky': '#E0F7FA',
  'pastel-aqua': '#E0F4F3',
  'pastel-teal': '#E6FFFA',
  'pastel-green': '#F0FDF4',
  'pastel-mint': '#ECFDF5',
  'pastel-sage': '#F0F9F4',
  'pastel-yellow': '#FEFCE8',
  'pastel-lemon': '#FFFBEB',
};
```

## イベントハンドラー

### onPlayerAction コールバック

```typescript
interface PlayerActionCallback {
  (action: string, data?: any): void;
}

// 使用例
const handlePlayerAction = (action: string, data?: any) => {
  if (action === 'interrupt_action') {
    console.log('プレイヤーが割り込み発言:', data.content);
    // ゲーム状態を更新
    updateGameState(data.content);
  }
};
```

### onInterrupt コールバック

```typescript
interface InterruptCallback {
  (itemId: string, action: 'cancel' | 'continue' | 'player_action'): void;
}

// 使用例
const handleInterrupt = (itemId: string, action: 'cancel' | 'continue' | 'player_action') => {
  console.log(`アイテム ${itemId} に対して ${action} を実行`);
  // 必要に応じて追加処理
  logInterruptAction(itemId, action);
};
```

## エラーハンドリング

### エラー表示

```typescript
// エラーが発生した場合、コンポーネント上部に表示
{error && (
  <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
    <strong>エラー:</strong> {error}
  </div>
)}
```

### よくあるエラー

1. **API接続エラー**: ネットワークやサーバーの問題
2. **認証エラー**: APIキーや認証トークンの問題
3. **データ形式エラー**: レスポンスが期待される形式と異なる
4. **タイムアウトエラー**: API応答が遅い

## パフォーマンス最適化

### 自動ポーリング

- **間隔**: 2秒ごとにキュー状態を更新
- **条件**: コンポーネントがマウントされている間のみ
- **停止**: アンマウント時に自動停止

### メモ化

- **useCallback**: イベントハンドラーのメモ化
- **useMemo**: 計算結果のキャッシュ
- **React.memo**: コンポーネントの再描画最適化

## 開発・デバッグ

### ローカル開発

```bash
# APIサーバー起動（別プロセス）
npm run api-server

# React開発サーバー起動
npm start
```

### デバッグ用ログ

```typescript
// 開発環境でのデバッグログ
if (process.env.NODE_ENV === 'development') {
  console.log('Queue items:', queueItems);
  console.log('Interrupt state:', interruptState);
}
```

## ライセンス・著作権

このコンポーネントはFake Gods TRPGプロジェクトの一部として開発されました。

---

**作成者**: Claude Code  
**最終更新**: 2024-06-22  
**バージョン**: 1.0.0