'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { Loader2 } from 'lucide-react';
import { useAuthStore } from '../store/auth.store';

import { authApi } from '../services/auth.api';

export const AuthGuard = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const [isChecking, setIsChecking] = useState(true);

  useEffect(() => {
    const verifyToken = async () => {
      const token = typeof window !== 'undefined' ? localStorage.getItem('access_token') : null;
      
      if (!token) {
        router.push('/login');
        return;
      }
      
      // Restore token to memory if missing
      const { accessToken, setAccessToken, logout } = useAuthStore.getState();
      if (!accessToken) {
        setAccessToken(token);
      }

      try {
        // Parse JWT payload without verifying signature (just to check expiration)
        const payloadStr = atob(token.split('.')[1].replace(/-/g, '+').replace(/_/g, '/'));
        const payload = JSON.parse(payloadStr);
        const isExpired = payload.exp * 1000 < Date.now();

        if (isExpired) {
          // Token is expired, forcefully refresh it using the HttpOnly cookie
          await authApi.refreshToken();
        }

        setIsChecking(false);
      } catch (err) {
        // If parsing fails or refresh fails (e.g. refresh token expired), clear state and login
        logout();
        router.push('/login');
      }
    };

    verifyToken();
  }, [router]);

  if (isChecking) {
    return (
      <div className="flex-1 flex items-center justify-center min-h-screen bg-background">
        <Loader2 className="w-8 h-8 animate-spin text-[#007AFF]" />
      </div>
    );
  }

  return <>{children}</>;
};
