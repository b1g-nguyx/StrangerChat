'use client';

import { motion } from 'framer-motion';
import { Message } from '../types';

interface ChatMessageBubbleProps {
  message: Message;
}

export function ChatMessageBubble({ message }: ChatMessageBubbleProps) {
  const isMe = message.sender === 'me';

  return (
    <motion.div
      initial={{ opacity: 0, y: 10 }}
      animate={{ opacity: 1, y: 0 }}
      className={`flex ${isMe ? 'justify-end' : 'justify-start'} w-full`}
    >
      <div
        className={`max-w-[75%] px-4 py-3 rounded-2xl ${
          isMe
            ? 'bg-[#007AFF] text-white rounded-br-sm'
            : 'bg-white dark:bg-[#1c1c1e] text-zinc-900 dark:text-zinc-50 rounded-bl-sm shadow-sm'
        }`}
      >
        <p className="text-[15px] leading-relaxed break-words">{message.text}</p>
        <span className={`text-[11px] mt-1 block ${isMe ? 'text-white/70' : 'text-zinc-400'}`}>
          {new Date(message.timestamp).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
        </span>
      </div>
    </motion.div>
  );
}
