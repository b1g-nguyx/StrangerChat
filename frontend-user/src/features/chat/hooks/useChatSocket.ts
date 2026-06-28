import { useState, useEffect, useRef, useCallback } from 'react';
import { Message, ConnectionStatus } from '../types';
import { useAuthStore } from '@/features/auth/store/auth.store';

export function useChatSocket() {
  const token = useAuthStore((state) => state.accessToken);
  const [status, setStatus] = useState<ConnectionStatus>('disconnected');
  const [messages, setMessages] = useState<Message[]>([]);
  const [roomId, setRoomId] = useState<string | null>(null);
  const wsRef = useRef<WebSocket | null>(null);

  const connect = useCallback(() => {
    if (!token) {
      console.warn('Cannot connect to WebSocket: Missing token');
      return;
    }

    const wsUrl = process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:8081/ws/chat';
    const wsEndpoint = `${wsUrl}?token=${token}`;
    const ws = new WebSocket(wsEndpoint);

    setStatus('connecting');
    ws.onopen = () => {
      // Initiate matchmaking immediately after connection
      ws.send(JSON.stringify({ type: 'FIND_MATCH' }));
    };

    ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);

        switch (data.type) {
          case 'MATCHED':
            setRoomId(data.room_id);
            setStatus('connected');
            setMessages([]);
            break;
            
          case 'CHAT':
            setMessages((prev) => [
              ...prev,
              {
                id: Date.now().toString() + Math.random(),
                text: data.content,
                sender: 'stranger',
                timestamp: new Date(),
              },
            ]);
            break;

          case 'PARTNER_LEFT':
            setStatus('disconnected');
            setRoomId(null);
            // Optionally add a system message here
            setMessages((prev) => [
              ...prev,
              {
                id: Date.now().toString(),
                text: "Người lạ đã thoát khỏi phòng.",
                sender: 'stranger', // Treat as a system message for UI simplicity
                timestamp: new Date(),
              },
            ]);
            break;
        }
      } catch (err) {
        console.error('Lỗi khi parse tin nhắn WS:', err);
      }
    };

    ws.onclose = () => {
      setStatus('disconnected');
      setRoomId(null);
    };

    ws.onerror = (error) => {
      console.error('WebSocket Error:', error);
      setStatus('disconnected');
    };

    wsRef.current = ws;
  }, []);

  const disconnect = useCallback(() => {
    if (wsRef.current) {
      // Chỉ gửi tin nhắn nếu WebSocket đã kết nối thành công (OPEN)
      if (wsRef.current.readyState === WebSocket.OPEN) {
        wsRef.current.send(JSON.stringify({ type: 'LEAVE_ROOM' }));
      }
      wsRef.current.close();
      wsRef.current = null;
      setStatus('disconnected');
      setRoomId(null);
    }
  }, []);

  // Connect on mount or when token becomes available
  useEffect(() => {
    if (token) {
      // eslint-disable-next-line react-hooks/set-state-in-effect
      connect();
    }
    return () => disconnect();
  }, [connect, disconnect, token]);

  const sendMessage = useCallback((text: string) => {
    if (!text.trim() || status !== 'connected' || !roomId || !wsRef.current) return;

    const newMsg: Message = {
      id: Date.now().toString(),
      text: text.trim(),
      sender: 'me',
      timestamp: new Date(),
    };
    setMessages((prev) => [...prev, newMsg]);

    // Gửi qua WS
    wsRef.current.send(
      JSON.stringify({
        type: 'CHAT',
        room_id: roomId,
        content: text.trim(),
      })
    );
  }, [roomId, status]);

  const findNewMatch = useCallback(() => {
    disconnect();
    setTimeout(() => {
      setMessages([]);
      connect();
    }, 500);
  }, [connect, disconnect]);

  return { status, messages, sendMessage, findNewMatch };
}
