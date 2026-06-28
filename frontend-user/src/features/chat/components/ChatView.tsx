'use client';

import { useState, useEffect, useRef, FormEvent } from 'react';
import { AnimatePresence, motion } from 'framer-motion';

import { ChatHeader } from './ChatHeader';
import { ChatInputArea } from './ChatInputArea';
import { ChatMessageBubble } from './ChatMessageBubble';
import { MatchmakingView } from './MatchmakingView';
import { useChatSocket } from '../hooks/useChatSocket';
import { Button } from '@/shared/components/Button';
import { RefreshCw } from 'lucide-react';

export function ChatView() {
  const { status, messages, sendMessage, findNewMatch } = useChatSocket();
  const [inputValue, setInputValue] = useState('');
  const messagesEndRef = useRef<HTMLDivElement>(null);

  // Auto-scroll to bottom
  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);

  const handleSendMessage = (e: FormEvent) => {
    e.preventDefault();
    if (!inputValue.trim()) return;
    sendMessage(inputValue);
    setInputValue('');
  };

  return (
    <div className="flex flex-col h-[100dvh] bg-[#f5f5f7] dark:bg-black overflow-hidden">
      <ChatHeader status={status} />

      <main className="flex-1 overflow-y-auto flex flex-col gap-4 relative">
        {status === 'connecting' && messages.length === 0 ? (
          <MatchmakingView />
        ) : (
          <div className="flex flex-col gap-4 p-4 sm:p-6 pb-2">
            <AnimatePresence initial={false}>
              {messages.map((msg) => (
                <ChatMessageBubble key={msg.id} message={msg} />
              ))}
            </AnimatePresence>
            <div ref={messagesEndRef} />
            {status === 'disconnected' && messages.length > 0 && (
              <motion.div 
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                className="flex justify-center mt-4 mb-4"
              >
                <Button 
                  onClick={findNewMatch}
                  className="bg-zinc-900 text-white dark:bg-white dark:text-black rounded-full h-12 px-6 flex items-center gap-2"
                >
                  <RefreshCw className="w-5 h-5" />
                  Tìm người lạ mới
                </Button>
              </motion.div>
            )}
          </div>
        )}
      </main>

      <ChatInputArea
        status={status}
        inputValue={inputValue}
        setInputValue={setInputValue}
        onSendMessage={handleSendMessage}
      />
    </div>
  );
}
