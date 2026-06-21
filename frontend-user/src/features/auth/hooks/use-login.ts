'use client';

import { useState } from 'react';
import { authApi } from '../services/auth.api';
import { useAuthStore } from '../store/auth.store';
import { LoginRequest } from '../types';

export const useLogin = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const setAuth = useAuthStore((state) => state.setAuth);

  const login = async (payload: LoginRequest) => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await authApi.login(payload);
      setAuth(response);
      return true;
    } catch (err: any) {
      setError(err.message || 'Something went wrong');
      return false;
    } finally {
      setIsLoading(false);
    }
  };

  return { login, isLoading, error };
};
