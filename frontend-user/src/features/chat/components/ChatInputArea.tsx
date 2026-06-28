'use client';

import { FormEvent } from 'react';
import { SendHorizontal } from 'lucide-react';
import { ConnectionStatus } from '../types';
import { Button } from '@/shared/components/Button';
import { Input } from '@/shared/components/Input';

interface ChatInputAreaProps {
  status: ConnectionStatus;
  inputValue: string;
  setInputValue: (value: string) => void;
  onSendMessage: (e: FormEvent) => void;
}

export function ChatInputArea({ status, inputValue, setInputValue, onSendMessage }: ChatInputAreaProps) {
  return (
    <form
      onSubmit={onSendMessage}
      className="p-4 bg-white/70 dark:bg-[#1c1c1e]/70 backdrop-blur-xl border-t border-black/5 dark:border-white/5 sticky bottom-0"
    >
      <div className="flex items-center gap-3 max-w-4xl mx-auto">
        <Input
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          placeholder={status === 'connected' ? 'Nhập tin nhắn...' : 'Đang chờ kết nối...'}
          disabled={status !== 'connected'}
          className="flex-1 rounded-full h-12 px-6 bg-zinc-100 dark:bg-zinc-800 border-transparent focus:bg-white dark:focus:bg-zinc-900 focus:border-[#007AFF] focus:ring-2 focus:ring-[#007AFF]/20 transition-all duration-200"
        />
        <Button
          type="submit"
          disabled={!inputValue.trim() || status !== 'connected'}
          className="w-12 h-12 rounded-full bg-[#007AFF] text-white flex items-center justify-center disabled:opacity-50 disabled:cursor-not-allowed hover:bg-[#007AFF]/90 active:scale-[0.98] transition-all"
        >
          <SendHorizontal className="w-5 h-5" />
        </Button>
      </div>
    </form>
  );
}
