import { create } from 'zustand';
import { User, AuthData } from '../types';

interface AuthState {
  user: User | null;
  accessToken: string | null;
  isAuthenticated: boolean;
  setAuth: (data: AuthData) => void;
  setAccessToken: (token: string) => void;
  setUser: (user: User) => void;
  logout: () => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  user: null,
  accessToken: null,
  isAuthenticated: false,

  setAuth: (data: AuthData) => {
    if (typeof window !== 'undefined') {
      localStorage.setItem('access_token', data.access_token);
    }
    set({ user: data.user, accessToken: data.access_token, isAuthenticated: true });
  },

  setAccessToken: (token: string) => {
    if (typeof window !== 'undefined') {
      localStorage.setItem('access_token', token);
    }
    set({ accessToken: token, isAuthenticated: true });
  },

  setUser: (user: User) => {
    set({ user });
  },

  logout: () => {
    if (typeof window !== 'undefined') {
      localStorage.removeItem('access_token');
    }
    set({ user: null, accessToken: null, isAuthenticated: false });
  },
}));
