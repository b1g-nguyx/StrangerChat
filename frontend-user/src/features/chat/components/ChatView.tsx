'use client';

import { useState, useEffect, useRef, FormEvent } from 'react';
import { AnimatePresence, motion } from 'framer-motion';

import { ChatHeader } from './ChatHeader';
import { ChatInputArea } from './ChatInputArea';
import { ChatMessageBubble } from './ChatMessageBubble';
import { MatchmakingView } from './MatchmakingView';
import { VideoCallView } from './VideoCallView';
import { useChatSocket } from '../hooks/useChatSocket';
import { useWebRTC } from '../hooks/useWebRTC';
import { Button } from '@/shared/components/Button';
import { RefreshCw } from 'lucide-react';

export function ChatView() {
  const { status, messages, sendMessage, findNewMatch, roomId } = useChatSocket();
  const { localStream, remoteStream, isVideoEnabled, enableVideo, stopCall } = useWebRTC(roomId);
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

      <main className="flex-1 overflow-hidden relative">
        {status === 'connecting' && messages.length === 0 ? (
          <MatchmakingView />
        ) : (
          <div className="w-full h-full flex flex-col">
            
            {/* Video Section (Top Row) */}
            {status === 'connected' && (
              <div className="w-full flex-shrink-0 border-b border-black/5 dark:border-white/5 bg-zinc-100 dark:bg-[#111] p-4 flex items-center justify-center">
                <VideoCallView 
                  localStream={localStream}
                  remoteStream={remoteStream}
                  isVideoEnabled={isVideoEnabled}
                  onToggleVideo={isVideoEnabled ? stopCall : enableVideo}
                  status={status}
                />
              </div>
            )}

            {/* Chat Section */}
            <div className="flex-1 flex flex-col overflow-hidden bg-white dark:bg-[#1c1c1e]/50">
              <div className="flex-1 overflow-y-auto p-4 sm:p-6 flex flex-col gap-4">
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
                      className="bg-zinc-900 text-white dark:bg-white dark:text-black rounded-full h-12 px-6 flex items-center gap-2 shadow-xl"
                    >
                      <RefreshCw className="w-5 h-5" />
                      Tìm người lạ mới
                    </Button>
                  </motion.div>
                )}
              </div>

              {/* Input Area */}
              <ChatInputArea
                status={status}
                inputValue={inputValue}
                setInputValue={setInputValue}
                onSendMessage={handleSendMessage}
              />
            </div>

          </div>
        )}
      </main>
    </div>
  );
}
