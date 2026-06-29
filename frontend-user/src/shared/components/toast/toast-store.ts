import { create } from 'zustand';

export type ToastType = 'success' | 'error' | 'info' | 'warning';

export interface ToastProps {
  id: string;
  title?: string;
  description: string;
  type?: ToastType;
  duration?: number;
}

interface ToastState {
  toasts: ToastProps[];
  addToast: (toast: Omit<ToastProps, 'id'>) => void;
  removeToast: (id: string) => void;
}

export const useToastStore = create<ToastState>((set) => ({
  toasts: [],
  addToast: (toast) => {
    const id = Math.random().toString(36).substring(2, 9);
    set((state) => ({
      toasts: [...state.toasts, { ...toast, id, duration: toast.duration ?? 3000 }],
    }));
  },
  removeToast: (id) =>
    set((state) => ({
      toasts: state.toasts.filter((t) => t.id !== id),
    })),
}));

// Utility function to call from anywhere
export const toast = {
  success: (description: string, title?: string) => useToastStore.getState().addToast({ type: 'success', description, title }),
  error: (description: string, title?: string) => useToastStore.getState().addToast({ type: 'error', description, title }),
  info: (description: string, title?: string) => useToastStore.getState().addToast({ type: 'info', description, title }),
  warning: (description: string, title?: string) => useToastStore.getState().addToast({ type: 'warning', description, title }),
};
