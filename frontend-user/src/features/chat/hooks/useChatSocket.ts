import { useState, useEffect, useCallback } from 'react';
import { Message, ConnectionStatus } from '../types';
import { useAuthStore } from '@/features/auth/store/auth.store';
import { chatSocketService } from '../services/chat.service';

export function useChatSocket() {
  const token = useAuthStore((state) => state.accessToken);
  const [status, setStatus] = useState<ConnectionStatus>('disconnected');
  const [messages, setMessages] = useState<Message[]>([]);
  const [roomId, setRoomId] = useState<string | null>(null);

  useEffect(() => {
    // Đăng ký nhận sự kiện từ service
    chatSocketService.subscribe(
      setStatus,
      (msgUpdate) => {
        if (typeof msgUpdate === 'function') {
          setMessages(msgUpdate);
        } else {
          setMessages(msgUpdate as Message[]);
        }
      },
      setRoomId
    );

    if (token) {
      chatSocketService.connect(token);
    }

    return () => {
      chatSocketService.disconnect();
    };
  }, [token]);

  const sendMessage = useCallback((text: string) => {
    chatSocketService.sendMessage(text);
  }, []);

  const findNewMatch = useCallback(() => {
    chatSocketService.disconnect();
    setTimeout(() => {
      setMessages([]);
      if (token) {
        chatSocketService.connect(token);
      }
    }, 500);
  }, [token]);

  return { status, messages, sendMessage, findNewMatch };
}
