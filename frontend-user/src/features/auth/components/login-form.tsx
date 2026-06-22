'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { useLogin } from '../hooks/use-login';
import { Input } from '@/shared/components/Input';
import { Button } from '@/shared/components/Button';
import Link from 'next/link';

export const LoginForm = () => {
  const router = useRouter();
  const { login, isLoading, error } = useLogin();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const success = await login({ email, password });
    if (success) {
      router.push('/');
    }
  };

  return (
    <div className="w-full max-w-md p-8 rounded-[32px] bg-white/70 dark:bg-[#1c1c1e]/70 backdrop-blur-xl border border-black/5 dark:border-white/10 shadow-[0_8px_30px_rgb(0,0,0,0.04)] dark:shadow-[0_8px_30px_rgb(0,0,0,0.12)]">
      <div className="mb-8 text-center">
        <h2 className="text-3xl font-bold text-zinc-900 dark:text-zinc-50 tracking-tight">Welcome back</h2>
        <p className="text-zinc-500 dark:text-zinc-400 mt-2 text-sm">Sign in to start chatting</p>
      </div>

      <form onSubmit={handleSubmit} className="space-y-4">
        <Input
          label="Email address"
          type="email"
          placeholder="user@example.com"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        
        <Input
          label="Password"
          type="password"
          placeholder="••••••••"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />

        {error && (
          <div className="p-3 text-sm text-[#FF3B30] bg-[#FF3B30]/10 rounded-2xl border border-[#FF3B30]/20">
            {error}
          </div>
        )}

        <Button type="submit" className="w-full mt-6" size="lg" isLoading={isLoading}>
          Sign In
        </Button>
      </form>

      <p className="mt-8 text-center text-sm text-zinc-500 dark:text-zinc-400">
        Don&apos;t have an account?{' '}
        <Link href="/register" className="text-[#007AFF] hover:opacity-80 transition-opacity font-medium">
          Sign up
        </Link>
      </p>
    </div>
  );
};
