import React, { useState } from 'react';
import { useVisibilityLock } from './hooks/useVisibilityLock';

interface VisibilityLockSystemProps {
  apiEndpoint: string;
  onPlayerAction?: (action: string, data?: any) => void;
  onInterrupt?: (itemId: string, action: 'cancel' | 'continue' | 'player_action') => void;
}

const VisibilityLockSystem: React.FC<VisibilityLockSystemProps> = ({
  apiEndpoint,
  onPlayerAction,
  onInterrupt,
}) => {
  const [showInputModal, setShowInputModal] = useState(false);
  const [playerInput, setPlayerInput] = useState('');
  
  const {
    currentDisplayItem,
    interruptState,
    queueStats,
    unseenCancelableItems,
    sessionCost,
    error,
    handleInterruptAction,
    cancelItem,
    getCountdownProgress,
  } = useVisibilityLock(apiEndpoint);

  // Open input modal for player action
  const openInputModal = () => {
    if (interruptState.isActive) {
      setShowInputModal(true);
      setPlayerInput('');
    }
  };

  // Handle modal submission
  const handleModalSubmit = async () => {
    if (playerInput.trim()) {
      await handleInterruptWithCallbacks('player_action', playerInput.trim());
      setShowInputModal(false);
      setPlayerInput('');
    }
  };

  // Handle modal cancel
  const handleModalCancel = () => {
    setShowInputModal(false);
    setPlayerInput('');
  };

  // Enhanced interrupt handler with callbacks
  const handleInterruptWithCallbacks = async (
    action: 'cancel' | 'continue' | 'player_action', 
    playerContent?: string
  ) => {
    const targetItemId = interruptState.targetItemId;
    
    // Handle the interrupt action
    await handleInterruptAction(action, playerContent);
    
    // Call provided callbacks
    if (action === 'player_action' && onPlayerAction) {
      onPlayerAction('interrupt_action', { content: playerContent });
    }
    
    if (onInterrupt && targetItemId) {
      onInterrupt(targetItemId, action);
    }
  };

  return (
    <div className="w-full h-screen bg-gradient-to-br from-pastel-purple to-pastel-pink p-4">
      {/* Error Display */}
      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          <strong>エラー:</strong> {error}
        </div>
      )}

      {/* Header */}
      <div className="bg-white rounded-xl shadow-lg p-4 mb-4 border-2 border-pastel-purple">
        <div className="flex justify-between items-center">
          <h1 className="text-xl font-bold text-gray-700 flex items-center gap-2">
            👁️ 視認ロックシステム
          </h1>
          <div className="text-sm text-gray-600">
            見たらキャンセル不可 | 生成中:{queueStats.generating}件 | 🛑割り込み:{interruptState.timeRemaining.toFixed(1)}秒
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4 h-5/6">
        {/* Left Panel - Currently Visible (LOCKED) */}
        <div className="flex flex-col gap-4">
          <div className="bg-pastel-green rounded-xl shadow-lg border-2 border-pastel-sage p-4 flex-1">
            <h2 className="text-lg font-semibold text-gray-700 mb-4 flex items-center gap-2">
              👁️ 表示中 (確定済み)
            </h2>
            
            {currentDisplayItem ? (
              <div className="text-center">
                <div className="text-4xl mb-2">{currentDisplayItem.speakerIcon}</div>
                <div className="text-lg font-semibold text-gray-700 mb-4">
                  {currentDisplayItem.speaker}
                </div>
                <div className="bg-white rounded-lg p-4 mb-4 border border-pastel-sage">
                  <p className="text-gray-800">{currentDisplayItem.content}</p>
                </div>
                <div className="text-sm text-gray-600">
                  🔒 視認済み → キャンセル不可
                </div>
              </div>
            ) : (
              <div className="text-center text-gray-500 py-8">
                表示中のアイテムはありません
              </div>
            )}
          </div>

          {/* Interrupt Window */}
          {interruptState.isActive && (
            <div className="bg-pastel-rose rounded-xl shadow-lg border-2 border-pastel-pink p-4">
              <h3 className="text-lg font-semibold text-gray-700 mb-4 flex items-center gap-2">
                ⚡ 割り込み窓口
              </h3>
              
              <div className="mb-4">
                <div className="text-sm text-gray-700 mb-2">
                  🛑 割り込み可能時間: {interruptState.timeRemaining.toFixed(1)}秒
                </div>
                
                {/* Progress Bar */}
                <div className="w-full bg-gray-200 rounded-full h-3 mb-2">
                  <div 
                    className="bg-gradient-to-r from-pastel-rose to-pastel-pink h-3 rounded-full transition-all duration-100 ease-linear"
                    style={{ width: `${getCountdownProgress()}%` }}
                  ></div>
                </div>
                
                <div className="text-sm text-gray-600">
                  ⚠️ {interruptState.timeRemaining.toFixed(1)}秒後に次の台詞確定
                </div>
              </div>

              {/* Interrupt Actions */}
              <div className="flex flex-col gap-2">
                <button
                  onClick={() => handleInterruptWithCallbacks('cancel')}
                  className="bg-pastel-rose hover:shadow-lg text-gray-700 px-4 py-2 rounded-lg font-medium transition-all border border-pastel-pink"
                >
                  🛑 キャンセル
                </button>
                <button
                  onClick={() => handleInterruptWithCallbacks('player_action', '緊急発言')}
                  className="bg-pastel-yellow hover:shadow-lg text-gray-700 px-4 py-2 rounded-lg font-medium transition-all border border-pastel-lemon"
                >
                  ⚡ 割り込み発言
                </button>
                <button
                  onClick={() => handleInterruptWithCallbacks('continue')}
                  className="bg-pastel-sky hover:shadow-lg text-gray-700 px-4 py-2 rounded-lg font-medium transition-all border border-pastel-aqua"
                >
                  🔄 続行
                </button>
              </div>
            </div>
          )}
        </div>

        {/* Right Panel - Pre-generated Queue (CANCELABLE) */}
        <div className="bg-pastel-aqua rounded-xl shadow-lg border-2 border-pastel-teal p-4">
          <h2 className="text-lg font-semibold text-gray-700 mb-4 flex items-center gap-2">
            🤖 事前生成 (キャンセル可能)
          </h2>
          
          <div className="text-sm text-gray-700 mb-4">
            待機中の台詞: {queueStats.queued}件
          </div>

          <div className="space-y-3 max-h-96 overflow-y-auto">
            {unseenCancelableItems.map((item) => (
              <div key={item.id} className="bg-white rounded-lg p-3 border border-pastel-teal">
                <div className="flex items-start justify-between mb-2">
                  <div className="flex items-center gap-2">
                    <span className="text-lg">{item.speakerIcon}</span>
                    <span className="font-medium text-gray-700">
                      #{item.displayOrder}: {item.speaker}
                    </span>
                  </div>
                  {item.canCancel && !interruptState.isActive && (
                    <button
                      onClick={() => cancelItem(item.id)}
                      className="text-pastel-rose hover:text-red-500 text-sm font-medium transition-colors"
                    >
                      ❌
                    </button>
                  )}
                </div>
                
                <div className="text-sm text-gray-700 mb-2">
                  「{item.content}」
                </div>
                
                <div className="text-xs text-gray-600">
                  状態: {item.canCancel ? '⏳未表示 = 🛑可能' : '🔒確定済み'}
                </div>
              </div>
            ))}
          </div>

          {unseenCancelableItems.length === 0 && (
            <div className="text-center text-gray-500 py-8">
              待機中のアイテムはありません
            </div>
          )}
        </div>
      </div>

      {/* System Rules Footer */}
      <div className="bg-white rounded-xl shadow-lg border-2 border-gray-200 p-4 mt-4">
        <h3 className="text-lg font-semibold text-gray-700 mb-3 flex items-center gap-2">
          📋 視認ロックルール
        </h3>
        
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2 text-sm">
          <div className="text-gray-700">
            ✅ 表示済み台詞: 🔒確定 → 物語に組み込み済み
          </div>
          <div className="text-gray-700">
            ⏳ 未表示台詞: 🛑キャンセル可能 → 割り込みで破棄
          </div>
          <div className="text-gray-700">
            🔄 表示直前: ⚡最後のチャンス → 2秒以内に決断
          </div>
          <div className="text-gray-700">
            👁️ 視認瞬間: 🔒即座に確定 → 取り消し不可
          </div>
          <div className="text-gray-700">
            📝 再生成: 確定台詞以降から → 一貫性保持
          </div>
          <div className="text-gray-700">
            ⚡ 緊急度: 表示直前ほど重要 → 物語の分岐点
          </div>
        </div>
        
        <div className="mt-3 text-xs text-gray-600 text-center">
          セッション費用: ${sessionCost.toFixed(4)} | 処理完了: {queueStats.completed}件
        </div>
      </div>

      {/* Bottom Controls */}
      <div className="fixed bottom-0 left-0 right-0 bg-white border-t-2 border-gray-200 p-4">
        <div className="flex justify-center items-center space-x-6 max-w-md mx-auto">
          {/* Scroll Down Button */}
          <button className="bg-pastel-sky hover:shadow-lg text-gray-700 p-4 rounded-full transition-all border-2 border-pastel-aqua">
            <span className="text-2xl">⬇️</span>
          </button>
          
          {/* Microphone Button for Interruption */}
          <button
            onClick={openInputModal}
            disabled={!interruptState.isActive}
            className={`p-6 rounded-full transition-all border-2 ${
              interruptState.isActive 
                ? 'bg-pastel-rose hover:shadow-lg text-gray-700 border-pastel-pink animate-pulse' 
                : 'bg-gray-200 text-gray-400 border-gray-300 cursor-not-allowed'
            }`}
          >
            <span className="text-3xl">🎤</span>
          </button>
          
          {/* Emergency Action Button */}
          <button className="bg-pastel-yellow hover:shadow-lg text-gray-700 p-4 rounded-full transition-all border-2 border-pastel-lemon">
            <span className="text-2xl">⚡</span>
          </button>
        </div>
        
        <div className="text-xs text-gray-600 text-center mt-2">
          {interruptState.isActive 
            ? `🎤 マイクで割り込み可能 (${interruptState.timeRemaining.toFixed(1)}秒)` 
            : '🎤 割り込み待機中'}
        </div>
      </div>

      {/* Input Modal */}
      {showInputModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
          <div className="bg-white rounded-xl shadow-2xl max-w-md w-full border-2 border-pastel-pink">
            {/* Modal Header */}
            <div className="bg-pastel-rose p-4 rounded-t-xl">
              <h3 className="text-lg font-semibold text-gray-700 text-center flex items-center justify-center gap-2">
                🎤 プレイヤー割り込み発言
              </h3>
              <div className="text-sm text-gray-600 text-center mt-1">
                残り時間: {interruptState.timeRemaining.toFixed(1)}秒
              </div>
            </div>
            
            {/* Modal Content */}
            <div className="p-6">
              <textarea
                value={playerInput}
                onChange={(e) => setPlayerInput(e.target.value)}
                placeholder="発言内容を入力してください..."
                className="w-full border-2 border-pastel-pink rounded-lg p-3 focus:ring-2 focus:ring-pastel-rose focus:border-transparent resize-none text-sm"
                rows={4}
                autoFocus
              />
              
              {/* Voice Input Hint */}
              <div className="mt-3 p-3 bg-pastel-yellow rounded-lg border border-pastel-lemon">
                <div className="text-sm text-gray-700 mb-2">
                  🎤 音声入力も可能（ブラウザの音声認識を使用）
                </div>
                <button
                  onClick={() => {
                    if ('webkitSpeechRecognition' in window) {
                      const recognition = new (window as any).webkitSpeechRecognition();
                      recognition.lang = 'ja-JP';
                      recognition.onresult = (event: any) => {
                        setPlayerInput(event.results[0][0].transcript);
                      };
                      recognition.start();
                    }
                  }}
                  className="text-xs bg-pastel-lemon hover:shadow-md px-3 py-1 rounded-full transition-all"
                >
                  🎙️ 音声で入力
                </button>
              </div>
            </div>
            
            {/* Modal Actions */}
            <div className="flex space-x-3 p-6 pt-0">
              <button
                onClick={handleModalCancel}
                className="flex-1 bg-gray-200 hover:shadow-lg text-gray-700 py-3 rounded-lg font-medium transition-all"
              >
                キャンセル
              </button>
              <button
                onClick={handleModalSubmit}
                disabled={!playerInput.trim()}
                className="flex-1 bg-pastel-rose hover:shadow-lg text-gray-700 py-3 rounded-lg font-medium transition-all disabled:opacity-50 disabled:cursor-not-allowed"
              >
                🎤 割り込み実行
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default VisibilityLockSystem;