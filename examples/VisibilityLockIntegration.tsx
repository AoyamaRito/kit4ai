import React, { useState, useCallback } from 'react';
import VisibilityLockSystem from '../VisibilityLockSystem';

/**
 * Example integration of VisibilityLockSystem with a Fake Gods TRPG application
 * This demonstrates how to integrate the visibility lock system with existing queue management
 */

interface GameState {
  currentScene: string;
  playerActions: string[];
  storyLog: string[];
  sessionId: string;
}

interface PlayerAction {
  type: 'interrupt' | 'normal' | 'emergency';
  content: string;
  timestamp: Date;
}

const VisibilityLockIntegration: React.FC = () => {
  const [gameState, setGameState] = useState<GameState>({
    currentScene: '村外れの神社',
    playerActions: [],
    storyLog: [],
    sessionId: `session-${Date.now()}`,
  });

  const [playerInput, setPlayerInput] = useState('');
  const [showVisibilityLock, setShowVisibilityLock] = useState(true);
  const [actionLog, setActionLog] = useState<PlayerAction[]>([]);

  // API endpoint configuration
  const API_ENDPOINT = process.env.REACT_APP_API_ENDPOINT || 'http://localhost:3001/api';

  // Handle player actions triggered by the visibility lock system
  const handlePlayerAction = useCallback((action: string, data?: any) => {
    const newAction: PlayerAction = {
      type: action === 'interrupt_action' ? 'interrupt' : 'normal',
      content: data?.content || '',
      timestamp: new Date(),
    };

    setActionLog(prev => [...prev, newAction]);
    
    // Add to game state
    setGameState(prev => ({
      ...prev,
      playerActions: [...prev.playerActions, newAction.content],
      storyLog: [...prev.storyLog, `[${newAction.type.toUpperCase()}] ${newAction.content}`],
    }));

    console.log(`Player action triggered: ${action}`, data);
  }, []);

  // Handle interrupt decisions
  const handleInterrupt = useCallback((itemId: string, action: 'cancel' | 'continue' | 'player_action') => {
    const logEntry = `Interrupt decision for item ${itemId}: ${action}`;
    
    setGameState(prev => ({
      ...prev,
      storyLog: [...prev.storyLog, `[SYSTEM] ${logEntry}`],
    }));

    console.log(`Interrupt handled: ${itemId} -> ${action}`);
  }, []);

  // Handle manual player input
  const handlePlayerInputSubmit = useCallback(async () => {
    if (!playerInput.trim()) return;

    const newAction: PlayerAction = {
      type: 'normal',
      content: playerInput,
      timestamp: new Date(),
    };

    setActionLog(prev => [...prev, newAction]);
    
    setGameState(prev => ({
      ...prev,
      playerActions: [...prev.playerActions, playerInput],
      storyLog: [...prev.storyLog, `[PLAYER] ${playerInput}`],
    }));

    // Send to queue API
    try {
      const response = await fetch(`${API_ENDPOINT}/queue`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          speaker: 'player',
          content: playerInput,
          timestamp: new Date().toISOString(),
        }),
      });

      if (!response.ok) {
        throw new Error('Failed to add player input to queue');
      }
    } catch (error) {
      console.error('Failed to submit player input:', error);
    }

    setPlayerInput('');
  }, [playerInput, API_ENDPOINT]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-pastel-sky to-pastel-mint p-4">
      {/* Header */}
      <div className="bg-white rounded-xl shadow-lg p-4 mb-4 border-2 border-pastel-teal">
        <div className="flex justify-between items-center">
          <div>
            <h1 className="text-2xl font-bold text-gray-700">Fake Gods TRPG - 視認ロックシステム統合版</h1>
            <p className="text-sm text-gray-600">現在のシーン: {gameState.currentScene}</p>
          </div>
          <div className="flex gap-2">
            <button
              onClick={() => setShowVisibilityLock(!showVisibilityLock)}
              className="bg-pastel-purple hover:shadow-lg text-gray-700 px-4 py-2 rounded-lg font-medium transition-all border border-pastel-purple"
            >
              {showVisibilityLock ? '👁️ ロック表示中' : '👁️ ロック非表示'}
            </button>
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 xl:grid-cols-3 gap-4">
        {/* Visibility Lock System - Takes up 2/3 of the space */}
        {showVisibilityLock && (
          <div className="xl:col-span-2">
            <VisibilityLockSystem
              apiEndpoint={API_ENDPOINT}
              onPlayerAction={handlePlayerAction}
              onInterrupt={handleInterrupt}
            />
          </div>
        )}

        {/* Game State & Controls */}
        <div className={`${showVisibilityLock ? 'xl:col-span-1' : 'xl:col-span-3'} space-y-4`}>
          {/* Player Input */}
          <div className="bg-white rounded-xl shadow-lg p-4 border-2 border-pastel-green">
            <h3 className="text-lg font-semibold text-gray-700 mb-3">プレイヤー入力</h3>
            <div className="space-y-3">
              <textarea
                value={playerInput}
                onChange={(e) => setPlayerInput(e.target.value)}
                placeholder="アクションや発言を入力してください..."
                className="w-full border border-pastel-green rounded-lg px-3 py-2 focus:ring-2 focus:ring-pastel-sage focus:border-transparent resize-none text-sm"
                rows={3}
              />
              <button
                onClick={handlePlayerInputSubmit}
                disabled={!playerInput.trim()}
                className="w-full bg-pastel-green hover:shadow-lg text-gray-700 px-4 py-2 rounded-lg font-medium transition-all border border-pastel-sage disabled:opacity-50 disabled:cursor-not-allowed"
              >
                キューに追加
              </button>
            </div>
          </div>

          {/* Action Log */}
          <div className="bg-white rounded-xl shadow-lg p-4 border-2 border-pastel-yellow">
            <h3 className="text-lg font-semibold text-gray-700 mb-3">アクションログ</h3>
            <div className="max-h-48 overflow-y-auto space-y-2">
              {actionLog.slice(-10).map((action, index) => (
                <div
                  key={index}
                  className={`text-xs p-2 rounded ${
                    action.type === 'interrupt' 
                      ? 'bg-pastel-rose border border-pastel-pink' 
                      : action.type === 'emergency'
                      ? 'bg-pastel-yellow border border-pastel-lemon'
                      : 'bg-pastel-mint border border-pastel-sage'
                  }`}
                >
                  <div className="flex justify-between items-start">
                    <span className="font-medium">
                      {action.type === 'interrupt' ? '🛑' : action.type === 'emergency' ? '⚡' : '💬'} 
                      {action.type.toUpperCase()}
                    </span>
                    <span className="text-gray-500">
                      {action.timestamp.toLocaleTimeString()}
                    </span>
                  </div>
                  <div className="mt-1">{action.content}</div>
                </div>
              ))}
              {actionLog.length === 0 && (
                <div className="text-center text-gray-500 py-4">
                  アクションログはありません
                </div>
              )}
            </div>
          </div>

          {/* Game State */}
          <div className="bg-white rounded-xl shadow-lg p-4 border-2 border-pastel-blue">
            <h3 className="text-lg font-semibold text-gray-700 mb-3">ゲーム状態</h3>
            <div className="space-y-2 text-sm">
              <div>
                <span className="font-medium text-gray-600">セッションID:</span>
                <div className="text-xs text-gray-500 font-mono">{gameState.sessionId}</div>
              </div>
              <div>
                <span className="font-medium text-gray-600">現在のシーン:</span>
                <div className="text-gray-700">{gameState.currentScene}</div>
              </div>
              <div>
                <span className="font-medium text-gray-600">実行済みアクション:</span>
                <div className="text-gray-700">{gameState.playerActions.length}件</div>
              </div>
              <div>
                <span className="font-medium text-gray-600">ストーリーログ:</span>
                <div className="text-gray-700">{gameState.storyLog.length}エントリ</div>
              </div>
            </div>
          </div>

          {/* System Status */}
          <div className="bg-white rounded-xl shadow-lg p-4 border-2 border-pastel-purple">
            <h3 className="text-lg font-semibold text-gray-700 mb-3">システム状態</h3>
            <div className="space-y-2 text-sm">
              <div className="flex justify-between">
                <span className="text-gray-600">視認ロック:</span>
                <span className={`font-medium ${showVisibilityLock ? 'text-green-600' : 'text-red-600'}`}>
                  {showVisibilityLock ? '有効' : '無効'}
                </span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-600">API接続:</span>
                <span className="font-medium text-green-600">接続中</span>
              </div>
              <div className="flex justify-between">
                <span className="text-gray-600">統合モード:</span>
                <span className="font-medium text-blue-600">フル機能</span>
              </div>
            </div>
          </div>

          {/* Integration Notes */}
          <div className="bg-pastel-sage rounded-xl shadow-lg p-4 border-2 border-pastel-mint">
            <h3 className="text-lg font-semibold text-gray-700 mb-3">統合機能説明</h3>
            <div className="space-y-2 text-xs text-gray-600">
              <div>• 視認ロックシステムが自動的に割り込み判定を管理</div>
              <div>• プレイヤーアクションは即座にゲーム状態に反映</div>
              <div>• 全てのアクションがログとして保存され追跡可能</div>
              <div>• API経由でキューシステムと連携</div>
              <div>• ストーリーの一貫性が自動的に保持される</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default VisibilityLockIntegration;