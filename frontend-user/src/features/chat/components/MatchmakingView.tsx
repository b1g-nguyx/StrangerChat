'use client';

import { motion } from 'framer-motion';
import { Loader2 } from 'lucide-react';

export function MatchmakingView() {
  return (
    <div className="flex-1 flex flex-col items-center justify-center p-6 text-center">
      <motion.div
        animate={{ scale: [1, 1.1, 1] }}
        transition={{ repeat: Infinity, duration: 2 }}
        className="w-24 h-24 bg-[#007AFF]/10 rounded-full flex items-center justify-center mb-6"
      >
        <Loader2 className="w-10 h-10 text-[#007AFF] animate-spin" />
      </motion.div>
      <h2 className="text-xl font-semibold text-zinc-900 dark:text-zinc-50 mb-2">Đang tìm người lạ...</h2>
      <p className="text-zinc-500 dark:text-zinc-400">Vui lòng chờ trong giây lát, chúng tôi đang kết nối bạn với một người ngẫu nhiên.</p>
    </div>
  );
}
