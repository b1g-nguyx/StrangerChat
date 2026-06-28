'use client';

import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { MessageCircle, User, LogOut, LogIn } from 'lucide-react';
import { Button } from '../Button';
import { useAuthStore, authApi } from '@/features/auth';

export const Header = () => {
  const router = useRouter();
  const { user, isAuthenticated, logout } = useAuthStore();

  const handleLogout = async () => {
    try {
      await authApi.logout();
    } catch (error) {
      console.error('Logout failed', error);
    } finally {
      logout();
      router.push('/');
    }
  };

  return (
    <header className="sticky top-0 z-50 w-full bg-white/70 dark:bg-[#1c1c1e]/70 backdrop-blur-xl border-b border-black/5 dark:border-white/10 transition-all duration-300">
      <div className="max-w-6xl mx-auto px-4 h-16 flex items-center justify-between">
        <Link href="/" className="flex items-center gap-2 group">
          <div className="w-8 h-8 bg-[#007AFF]/10 dark:bg-[#007AFF]/20 rounded-full flex items-center justify-center transition-all group-hover:scale-105">
            <MessageCircle className="w-5 h-5 text-[#007AFF]" />
          </div>
          <span className="font-bold text-lg text-zinc-900 dark:text-zinc-50 tracking-tight">Stranger Chat</span>
        </Link>
        
        <nav className="flex items-center gap-4">
          {isAuthenticated ? (
            <>
              <div className="flex items-center gap-2 text-zinc-600 dark:text-zinc-400">
                <User className="w-4 h-4" />
                <span className="text-sm font-medium">{user?.username}</span>
              </div>
              <Button variant="ghost" size="sm" onClick={handleLogout} className="gap-2">
                <LogOut className="w-4 h-4" />
                <span className="hidden sm:inline">Logout</span>
              </Button>
            </>
          ) : (
            <Button variant="secondary" size="sm" onClick={() => router.push('/login')} className="gap-2">
              <LogIn className="w-4 h-4" />
              <span>Login</span>
            </Button>
          )}
        </nav>
      </div>
    </header>
  );
};
