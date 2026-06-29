import { Button } from '@/shared/components/Button';
import { Settings, MessageCircle } from 'lucide-react';
import { Header } from '@/shared/components/layout/Header';
import { Footer } from '@/shared/components/layout/Footer';
import { HeroSlider } from '@/shared/components/layout/HeroSlider';
import Link from 'next/link';

export default function Home() {
  return (
    <div className="flex flex-col min-h-[100dvh] bg-transparent">
      <Header />
      
      <main className="flex-1 flex flex-col items-center justify-center p-4 py-12">
        {/* Animated Slider Section */}
        <HeroSlider />

        {/* Call To Action Block */}
        <div className="w-full max-w-2xl text-center flex flex-col items-center mt-4">
          <h1 className="text-4xl sm:text-5xl font-bold text-zinc-900 dark:text-zinc-50 tracking-tight mb-4">
            Ready to jump in?
          </h1>
          <p className="text-lg text-zinc-500 dark:text-zinc-400 mb-10 max-w-lg">
            No registration required. Just click below and start chatting with a random stranger immediately.
          </p>

          <div className="flex flex-col sm:flex-row w-full max-w-md gap-4 justify-center">
            <Link href="/chat" className="w-full sm:w-auto flex-1">
              <Button size="lg" className="w-full flex items-center justify-center gap-2 text-lg">
                <MessageCircle className="w-5 h-5" />
                Start Chatting
              </Button>
            </Link>
            <Button variant="secondary" size="lg" className="w-full sm:w-auto flex items-center justify-center gap-2">
              <Settings className="w-5 h-5" />
              Settings
            </Button>
          </div>
        </div>
      </main>

      <Footer />
    </div>
  );
}
