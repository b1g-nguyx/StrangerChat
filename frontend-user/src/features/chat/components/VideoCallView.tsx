'use client';

import { useEffect, useRef } from 'react';
import { Video, VideoOff } from 'lucide-react';
import { Button } from '@/shared/components/Button';

interface VideoCallViewProps {
  localStream: MediaStream | null;
  remoteStream: MediaStream | null;
  isVideoEnabled: boolean;
  onToggleVideo: () => void;
  status: 'disconnected' | 'connecting' | 'connected';
}

export function VideoCallView({ localStream, remoteStream, isVideoEnabled, onToggleVideo, status }: VideoCallViewProps) {
  const localVideoRef = useRef<HTMLVideoElement>(null);
  const remoteVideoRef = useRef<HTMLVideoElement>(null);

  useEffect(() => {
    if (localVideoRef.current && localStream) {
      localVideoRef.current.srcObject = localStream;
    }
  }, [localStream]);

  useEffect(() => {
    if (remoteVideoRef.current && remoteStream) {
      remoteVideoRef.current.srcObject = remoteStream;
    }
  }, [remoteStream]);

  if (status !== 'connected') return null;

  return (
    <div className="flex flex-row gap-4 w-full max-w-4xl mx-auto items-center justify-center">
      {/* Remote Video (Stranger) */}
      <div className="relative flex-1 aspect-square sm:aspect-video max-h-[35vh] bg-zinc-200 dark:bg-zinc-800 rounded-2xl overflow-hidden flex items-center justify-center shadow-inner">
        {remoteStream ? (
          <video
            ref={remoteVideoRef}
            autoPlay
            playsInline
            className="w-full h-full object-cover"
          />
        ) : (
          <div className="text-zinc-500 dark:text-zinc-400 flex flex-col items-center">
            <VideoOff className="w-8 h-8 mb-2 opacity-50" />
            <span className="text-sm font-medium">Người lạ tắt cam</span>
          </div>
        )}
      </div>

      {/* Local Video (Me) */}
      <div className="relative flex-1 aspect-square sm:aspect-video max-h-[35vh] bg-zinc-200 dark:bg-zinc-800 rounded-2xl overflow-hidden flex items-center justify-center shadow-inner">
        {localStream ? (
          <video
            ref={localVideoRef}
            autoPlay
            playsInline
            muted
            className="w-full h-full object-cover scale-x-[-1]"
          />
        ) : (
          <div className="text-zinc-500 dark:text-zinc-400 flex flex-col items-center">
            <VideoOff className="w-8 h-8 mb-2 opacity-50" />
            <span className="text-sm font-medium">Bạn tắt cam</span>
          </div>
        )}
        
        {/* Toggle Button */}
        <div className="absolute bottom-4 left-1/2 -translate-x-1/2">
          <Button
            onClick={onToggleVideo}
            className={`rounded-full w-12 h-12 flex items-center justify-center backdrop-blur-md shadow-lg ${
              isVideoEnabled 
                ? 'bg-red-500/90 text-white hover:bg-red-600' 
                : 'bg-zinc-900/90 text-white dark:bg-white/90 dark:text-black hover:bg-zinc-800 dark:hover:bg-zinc-200'
            }`}
          >
            {isVideoEnabled ? <VideoOff className="w-5 h-5" /> : <Video className="w-5 h-5" />}
          </Button>
        </div>
      </div>
    </div>
  );
}
