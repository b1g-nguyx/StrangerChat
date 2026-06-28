'use client';

import { ConnectionStatus } from '../types';

interface ChatHeaderProps {
  status: ConnectionStatus;
}

export function ChatHeader({ status }: ChatHeaderProps) {
  return (
    <header className="h-16 flex items-center justify-between px-6 bg-white/70 dark:bg-[#1c1c1e]/70 backdrop-blur-xl border-b border-black/5 dark:border-white/5 sticky top-0 z-10">
      <h1 className="text-lg font-semibold text-zinc-900 dark:text-zinc-50">Stranger Chat</h1>
      <div className="flex items-center gap-2">
        <div className={`w-2 h-2 rounded-full ${status === 'connected' ? 'bg-green-500' : status === 'connecting' ? 'bg-yellow-500 animate-pulse' : 'bg-red-500'}`} />
        <span className="text-sm font-medium text-zinc-500 dark:text-zinc-400">
          {status === 'connected' ? 'Đã kết nối' : status === 'connecting' ? 'Đang tìm kiếm...' : 'Ngắt kết nối'}
        </span>
      </div>
    </header>
  );
}
