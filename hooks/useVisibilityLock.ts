import { useState, useEffect, useCallback, useRef } from 'react';

// Interface definitions
export interface QueueItem {
  id: string;
  speaker: string;
  content: string;
  speakerIcon: string;
  timestamp: Date;
  status: 'generating' | 'queued' | 'displaying' | 'completed';
  isVisible: boolean; // Visibility state tracking
  canCancel: boolean; // Whether this item can be cancelled
  displayOrder: number;
}

export interface InterruptState {
  isActive: boolean;
  timeRemaining: number; // in seconds
  targetItemId: string | null;
  countdownTimer: NodeJS.Timeout | null;
}

export interface QueueStats {
  generating: number;
  queued: number;
  completed: number;
  total: number;
}

// API service class
export class VisibilityLockAPI {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  async getQueue(): Promise<QueueItem[]> {
    const response = await fetch(`${this.baseUrl}/queue`);
    if (!response.ok) throw new Error('Failed to fetch queue');
    const data = await response.json();
    
    // Transform API response to match our interface
    return data.map((item: any, index: number) => ({
      id: item.id || `item-${index}`,
      speaker: item.speaker || 'Unknown',
      content: item.content || '',
      speakerIcon: this.getSpeakerIcon(item.speaker),
      timestamp: new Date(item.timestamp || Date.now()),
      status: item.status || 'queued',
      isVisible: item.isVisible || false,
      canCancel: item.canCancel !== false, // Default to true unless explicitly false
      displayOrder: item.displayOrder || index + 1,
    }));
  }

  async addToQueue(speaker: string, content: string): Promise<void> {
    const response = await fetch(`${this.baseUrl}/queue`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ 
        speaker, 
        content, 
        timestamp: new Date().toISOString(),
        status: 'queued',
        isVisible: false,
        canCancel: true
      }),
    });
    if (!response.ok) throw new Error('Failed to add to queue');
  }

  async cancelItem(itemId: string): Promise<void> {
    const response = await fetch(`${this.baseUrl}/queue/${itemId}`, {
      method: 'DELETE',
    });
    if (!response.ok) throw new Error('Failed to cancel item');
  }

  async markAsVisible(itemId: string): Promise<void> {
    const response = await fetch(`${this.baseUrl}/queue/${itemId}/visible`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ isVisible: true, canCancel: false }),
    });
    if (!response.ok) throw new Error('Failed to mark as visible');
  }

  async interruptQueue(itemId: string, action: string, playerContent?: string): Promise<void> {
    const response = await fetch(`${this.baseUrl}/queue/${itemId}/interrupt`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ action, playerContent }),
    });
    if (!response.ok) throw new Error('Failed to interrupt queue');
  }

  private getSpeakerIcon(speaker: string): string {
    const iconMap: Record<string, string> = {
      'amaterasu': '☀️',
      'orihime': '🎭',
      'guide': '🧚‍♀️',
      'childhood_friend': '👧',
      'goblin': '👹',
      'gm': '🎮',
      'player': '⚔️',
      'system': '🤖',
    };
    
    const normalizedSpeaker = speaker.toLowerCase().replace(/\s+/g, '_');
    return iconMap[normalizedSpeaker] || '💬';
  }
}

// Custom hook for managing visibility lock system
export function useVisibilityLock(apiEndpoint: string) {
  const [queueItems, setQueueItems] = useState<QueueItem[]>([]);
  const [currentDisplayItem, setCurrentDisplayItem] = useState<QueueItem | null>(null);
  const [interruptState, setInterruptState] = useState<InterruptState>({
    isActive: false,
    timeRemaining: 0,
    targetItemId: null,
    countdownTimer: null,
  });
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [sessionCost, setSessionCost] = useState(0.0012);

  const api = useRef(new VisibilityLockAPI(apiEndpoint));
  const pollIntervalRef = useRef<NodeJS.Timeout | null>(null);

  // Load queue items from API
  const loadQueue = useCallback(async () => {
    try {
      setIsLoading(true);
      setError(null);
      const items = await api.current.getQueue();
      setQueueItems(items);
      
      // Find current display item (first item being displayed)
      const displayItem = items.find(item => item.status === 'displaying') || null;
      setCurrentDisplayItem(displayItem);
      
      // Update session cost (mock calculation)
      setSessionCost(prev => prev + 0.0001);
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to load queue';
      setError(errorMessage);
      console.error('Failed to load queue:', err);
    } finally {
      setIsLoading(false);
    }
  }, []);

  // Initialize queue and polling
  useEffect(() => {
    loadQueue();
    
    // Set up polling
    pollIntervalRef.current = setInterval(loadQueue, 2000); // Poll every 2 seconds
    
    return () => {
      if (pollIntervalRef.current) {
        clearInterval(pollIntervalRef.current);
      }
    };
  }, [loadQueue]);

  // Handle item visibility change (when displayed to player)
  const markItemAsVisible = useCallback(async (item: QueueItem) => {
    if (item.isVisible) return; // Already visible

    try {
      await api.current.markAsVisible(item.id);
      
      setQueueItems(prev => 
        prev.map(qItem => 
          qItem.id === item.id 
            ? { ...qItem, isVisible: true, canCancel: false, status: 'completed' as const }
            : qItem
        )
      );
      
      setCurrentDisplayItem({ ...item, isVisible: true, canCancel: false, status: 'completed' });
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to mark item as visible';
      setError(errorMessage);
      console.error('Failed to mark item as visible:', err);
    }
  }, []);

  // Start interrupt countdown for unseen items
  const startInterruptCountdown = useCallback((targetItem: QueueItem) => {
    if (interruptState.isActive) return; // Already active

    const countdown = setInterval(() => {
      setInterruptState(prev => {
        const newTimeRemaining = Math.max(0, prev.timeRemaining - 0.1);
        
        if (newTimeRemaining <= 0) {
          // Time's up, automatically display the item
          clearInterval(countdown);
          markItemAsVisible(targetItem);
          return {
            isActive: false,
            timeRemaining: 0,
            targetItemId: null,
            countdownTimer: null,
          };
        }
        
        return { ...prev, timeRemaining: newTimeRemaining };
      });
    }, 100); // Update every 100ms for smooth countdown

    setInterruptState({
      isActive: true,
      timeRemaining: 2.0, // 2 second window
      targetItemId: targetItem.id,
      countdownTimer: countdown,
    });
  }, [interruptState.isActive, markItemAsVisible]);

  // Handle interrupt actions
  const handleInterruptAction = useCallback(async (
    action: 'cancel' | 'continue' | 'player_action', 
    playerContent?: string
  ) => {
    if (!interruptState.targetItemId) return;

    try {
      // Clear countdown timer
      if (interruptState.countdownTimer) {
        clearInterval(interruptState.countdownTimer);
      }

      await api.current.interruptQueue(interruptState.targetItemId, action, playerContent);
      
      if (action === 'cancel') {
        // Remove cancelled item from queue
        setQueueItems(prev => prev.filter(item => item.id !== interruptState.targetItemId));
      } else if (action === 'continue') {
        // Allow item to display normally
        const targetItem = queueItems.find(item => item.id === interruptState.targetItemId);
        if (targetItem) {
          await markItemAsVisible(targetItem);
        }
      }

      // Reset interrupt state
      setInterruptState({
        isActive: false,
        timeRemaining: 0,
        targetItemId: null,
        countdownTimer: null,
      });

    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to handle interrupt action';
      setError(errorMessage);
      console.error('Failed to handle interrupt action:', err);
    }
  }, [interruptState, queueItems, markItemAsVisible]);

  // Cancel an item
  const cancelItem = useCallback(async (itemId: string) => {
    const item = queueItems.find(q => q.id === itemId);
    if (!item || !item.canCancel) return;

    try {
      await api.current.cancelItem(itemId);
      setQueueItems(prev => prev.filter(item => item.id !== itemId));
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to cancel item';
      setError(errorMessage);
      console.error('Failed to cancel item:', err);
    }
  }, [queueItems]);

  // Add item to queue
  const addToQueue = useCallback(async (speaker: string, content: string) => {
    try {
      await api.current.addToQueue(speaker, content);
      await loadQueue(); // Refresh queue
    } catch (err) {
      const errorMessage = err instanceof Error ? err.message : 'Failed to add to queue';
      setError(errorMessage);
      console.error('Failed to add to queue:', err);
    }
  }, [loadQueue]);

  // Auto-trigger interrupt window for next unseen item
  useEffect(() => {
    if (!interruptState.isActive && !currentDisplayItem) {
      const nextUnseenItem = queueItems.find(item => 
        !item.isVisible && 
        item.status === 'queued' && 
        item.canCancel
      );
      
      if (nextUnseenItem) {
        // Small delay to allow UI to update
        const timer = setTimeout(() => {
          startInterruptCountdown(nextUnseenItem);
        }, 500);
        
        return () => clearTimeout(timer);
      }
    }
  }, [queueItems, currentDisplayItem, interruptState.isActive, startInterruptCountdown]);

  // Calculate queue statistics
  const queueStats: QueueStats = {
    generating: queueItems.filter(item => item.status === 'generating').length,
    queued: queueItems.filter(item => item.status === 'queued').length,
    completed: queueItems.filter(item => item.status === 'completed').length,
    total: queueItems.length,
  };

  // Get progress percentage for countdown bar
  const getCountdownProgress = (): number => {
    return interruptState.isActive ? ((2.0 - interruptState.timeRemaining) / 2.0) * 100 : 0;
  };

  // Get unseen cancelable items
  const unseenCancelableItems = queueItems.filter(item => 
    !item.isVisible && 
    item.status !== 'generating' && 
    item.canCancel
  );

  // Cleanup on unmount
  useEffect(() => {
    return () => {
      if (interruptState.countdownTimer) {
        clearInterval(interruptState.countdownTimer);
      }
      if (pollIntervalRef.current) {
        clearInterval(pollIntervalRef.current);
      }
    };
  }, [interruptState.countdownTimer]);

  return {
    // State
    queueItems,
    currentDisplayItem,
    interruptState,
    isLoading,
    error,
    sessionCost,
    queueStats,
    unseenCancelableItems,
    
    // Actions
    markItemAsVisible,
    handleInterruptAction,
    cancelItem,
    addToQueue,
    loadQueue,
    
    // Utilities
    getCountdownProgress,
  };
}