'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { Loader2 } from 'lucide-react';

export const AuthGuard = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const [isChecking, setIsChecking] = useState(true);

  useEffect(() => {
    const token = typeof window !== 'undefined' ? localStorage.getItem('access_token') : null;
    
    if (!token) {
      router.push('/login');
    } else {
      setTimeout(() => setIsChecking(false), 0);
    }
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
