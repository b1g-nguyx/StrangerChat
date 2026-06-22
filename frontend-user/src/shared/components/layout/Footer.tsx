import Link from 'next/link';
import { MessageCircle, Mail, Globe } from 'lucide-react';

export const Footer = () => {
  return (
    <footer className="w-full bg-transparent border-t border-black/5 dark:border-white/10 mt-auto py-8">
      <div className="max-w-6xl mx-auto px-4 flex flex-col md:flex-row justify-between items-center gap-4">
        <div className="flex items-center gap-2">
          <MessageCircle className="w-5 h-5 text-zinc-400" />
          <span className="font-semibold text-zinc-500 dark:text-zinc-400">Stranger Chat</span>
        </div>
        <div className="text-sm text-zinc-400 dark:text-zinc-500">
          © {new Date().getFullYear()} Stranger Chat. Minimal & Beautiful.
        </div>
        <div className="flex items-center gap-4 text-zinc-400">
          <Link href="#" className="hover:text-zinc-900 dark:hover:text-zinc-50 transition-colors">
            <Globe className="w-5 h-5" />
          </Link>
          <Link href="#" className="hover:text-zinc-900 dark:hover:text-zinc-50 transition-colors">
            <Mail className="w-5 h-5" />
          </Link>
        </div>
      </div>
    </footer>
  );
};
