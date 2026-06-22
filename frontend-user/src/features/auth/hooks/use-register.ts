'use client';

import { useState } from 'react';
import { authApi } from '../services/auth.api';
import { useAuthStore } from '../store/auth.store';
import { RegisterRequest } from '../types';

export const useRegister = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const setAuth = useAuthStore((state) => state.setAuth);

  const register = async (payload: RegisterRequest) => {
    setIsLoading(true);
    setError(null);
    try {
      const response = await authApi.register(payload);
      setAuth(response);
      return true;
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError('Something went wrong');
      }
      return false;
    } finally {
      setIsLoading(false);
    }
  };

  return { register, isLoading, error };
};
