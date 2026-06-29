'use client';

import { AnimatePresence } from 'framer-motion';
import { useToastStore } from './toast-store';
import { Toast } from './Toast';

export function ToastContainer() {
  const toasts = useToastStore((state) => state.toasts);

  return (
    <div className="fixed top-4 left-1/2 -translate-x-1/2 z-50 flex flex-col gap-2 pointer-events-none w-full max-w-sm px-4">
      <AnimatePresence>
        {toasts.map((t) => (
          <Toast key={t.id} {...t} />
        ))}
      </AnimatePresence>
    </div>
  );
}
