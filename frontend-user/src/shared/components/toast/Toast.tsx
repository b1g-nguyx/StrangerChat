'use client';

import { useEffect } from 'react';
import { motion } from 'framer-motion';
import { CheckCircle2, AlertCircle, Info, AlertTriangle, X } from 'lucide-react';
import { ToastProps, useToastStore } from './toast-store';

export function Toast({ id, title, description, type = 'info', duration = 3000 }: ToastProps) {
  const removeToast = useToastStore((state) => state.removeToast);

  useEffect(() => {
    if (duration > 0) {
      const timer = setTimeout(() => {
        removeToast(id);
      }, duration);
      return () => clearTimeout(timer);
    }
  }, [duration, id, removeToast]);

  const icons = {
    success: <CheckCircle2 className="w-5 h-5 text-green-500" />,
    error: <AlertCircle className="w-5 h-5 text-red-500" />,
    info: <Info className="w-5 h-5 text-blue-500" />,
    warning: <AlertTriangle className="w-5 h-5 text-yellow-500" />
  };

  return (
    <motion.div
      layout
      initial={{ opacity: 0, y: -20, scale: 0.95 }}
      animate={{ opacity: 1, y: 0, scale: 1 }}
      exit={{ opacity: 0, scale: 0.95, transition: { duration: 0.2 } }}
      className="pointer-events-auto flex items-start gap-3 p-4 w-full max-w-sm bg-white/80 dark:bg-[#1c1c1e]/80 backdrop-blur-xl border border-black/5 dark:border-white/5 rounded-2xl shadow-xl shadow-black/5 dark:shadow-black/20"
    >
      <div className="shrink-0 mt-0.5">{icons[type]}</div>
      <div className="flex-1 min-w-0">
        {title && <h4 className="text-sm font-semibold text-zinc-900 dark:text-zinc-50 mb-1">{title}</h4>}
        <p className="text-sm text-zinc-600 dark:text-zinc-400">{description}</p>
      </div>
      <button 
        onClick={() => removeToast(id)}
        className="shrink-0 text-zinc-400 hover:text-zinc-600 dark:hover:text-zinc-200 transition-colors p-1 -m-1"
      >
        <X className="w-4 h-4" />
      </button>
    </motion.div>
  );
}
