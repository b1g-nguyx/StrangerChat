'use client';

import { useEffect, useState } from 'react';
import axios from 'axios';
import { Loader2 } from 'lucide-react';
import { useAuthStore } from '../store/auth.store';
import { authApi } from '../services/auth.api';

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isInitializing, setIsInitializing] = useState(true);
  const setAccessToken = useAuthStore((state) => state.setAccessToken);
  const setUser = useAuthStore((state) => state.setUser);
  const logout = useAuthStore((state) => state.logout);

  useEffect(() => {
    const initializeAuth = async () => {
      try {
        const baseURL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/v1';
        // Attempt to silently refresh token using the HTTP-Only cookie
        const response = await axios.post(
          `${baseURL}/auth/refresh`,
          {},
          { withCredentials: true }
        );
        
        if (response.data && response.data.access_token) {
          setAccessToken(response.data.access_token);
          // Fetch user info using the new access token
          // Since apiClient is configured to use token from Zustand, we can just call getMe()
          // after setAccessToken completes. In Zustand, set actions are synchronous.
          const user = await authApi.getMe();
          setUser(user);
        }
      } catch (error) {
        // If refresh fails (e.g. no cookie or expired), ensure logged out state
        logout();
      } finally {
        setIsInitializing(false);
      }
    };

    initializeAuth();
  }, [setAccessToken, setUser, logout]);

  if (isInitializing) {
    return (
      <div className="flex-1 flex items-center justify-center min-h-[100dvh] bg-background">
        <Loader2 className="w-10 h-10 animate-spin text-[#007AFF]" />
      </div>
    );
  }

  return <>{children}</>;
};
