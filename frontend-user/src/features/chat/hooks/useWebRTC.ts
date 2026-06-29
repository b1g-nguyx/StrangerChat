import { useState, useEffect, useRef, useCallback } from 'react';
import { chatSocketService } from '../services/chat.service';
import { toast } from '@/shared/components/toast';

const configuration = {
  iceServers: [
    { urls: 'stun:stun.l.google.com:19302' },
    { urls: 'stun:stun1.l.google.com:19302' },
  ],
};

export function useWebRTC(roomId: string | null) {
  const [localStream, setLocalStream] = useState<MediaStream | null>(null);
  const [remoteStream, setRemoteStream] = useState<MediaStream | null>(null);
  const [isVideoEnabled, setIsVideoEnabled] = useState(false);
  
  const peerConnection = useRef<RTCPeerConnection | null>(null);

  // Initialize peer connection
  const initPeerConnection = useCallback(() => {
    if (peerConnection.current) return;

    const pc = new RTCPeerConnection(configuration);

    pc.onicecandidate = (event) => {
      if (event.candidate) {
        chatSocketService.sendWebRTCSignal('WEBRTC_ICE_CANDIDATE', event.candidate);
      }
    };

    pc.ontrack = (event) => {
      if (event.streams && event.streams[0]) {
        setRemoteStream(event.streams[0]);
      }
    };

    peerConnection.current = pc;
  }, []);

  // Cleanup WebRTC
  const cleanup = useCallback(() => {
    if (peerConnection.current) {
      peerConnection.current.close();
      peerConnection.current = null;
    }
    if (localStream) {
      localStream.getTracks().forEach(track => track.stop());
    }
    setLocalStream(null);
    setRemoteStream(null);
    setIsVideoEnabled(false);
  }, [localStream]);

  // Handle incoming signals
  useEffect(() => {
    chatSocketService.subscribeWebRTC(async (type, payload) => {
      if (!peerConnection.current) initPeerConnection();
      const pc = peerConnection.current!;

      try {
        if (type === 'WEBRTC_OFFER') {
          await pc.setRemoteDescription(new RTCSessionDescription(payload));
          const answer = await pc.createAnswer();
          await pc.setLocalDescription(answer);
          chatSocketService.sendWebRTCSignal('WEBRTC_ANSWER', answer);
          
          // Check if remote stopped sharing
          const hasVideo = pc.getReceivers().some(r => r.track && r.track.readyState === 'live');
          if (!hasVideo) setRemoteStream(null);
          
        } else if (type === 'WEBRTC_ANSWER') {
          await pc.setRemoteDescription(new RTCSessionDescription(payload));
          
          // Check if remote stopped sharing
          const hasVideo = pc.getReceivers().some(r => r.track && r.track.readyState === 'live');
          if (!hasVideo) setRemoteStream(null);
          
        } else if (type === 'WEBRTC_ICE_CANDIDATE') {
          await pc.addIceCandidate(new RTCIceCandidate(payload));
        }
      } catch (err) {
        console.error('WebRTC error handling signal:', err);
      }
    });

    return () => {
      chatSocketService.subscribeWebRTC(() => {});
    };
  }, [initPeerConnection]);

  // When room is disconnected, cleanup
  useEffect(() => {
    if (!roomId) {
      cleanup();
    }
  }, [roomId, cleanup]);

  const startCall = async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
      setLocalStream(stream);
      setIsVideoEnabled(true);
      
      initPeerConnection();
      const pc = peerConnection.current!;
      
      stream.getTracks().forEach((track) => {
        pc.addTrack(track, stream);
      });

      const offer = await pc.createOffer();
      await pc.setLocalDescription(offer);
      chatSocketService.sendWebRTCSignal('WEBRTC_OFFER', offer);
      
    } catch (err) {
      console.error('Error starting video call:', err);
      toast.error('Vui lòng cấp quyền truy cập Camera và Microphone trong trình duyệt.', 'Lỗi thiết bị');
    }
  };
  
  // Accept the call (start our own camera and add tracks if not already added)
  // This can happen automatically when we receive an offer, or user can click a button to join.
  // We'll auto-add tracks if we receive an offer and we choose to enable video.
  const enableVideo = async () => {
    try {
      const stream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
      setLocalStream(stream);
      setIsVideoEnabled(true);
      
      initPeerConnection();
      const pc = peerConnection.current!;
      
      stream.getTracks().forEach((track) => {
        // Prevent adding duplicate tracks
        const senders = pc.getSenders();
        const exists = senders.find(s => s.track === track);
        if (!exists) {
          pc.addTrack(track, stream);
        }
      });
      
      // We might need renegotiation here (send an offer) because we added new tracks
      const offer = await pc.createOffer();
      await pc.setLocalDescription(offer);
      chatSocketService.sendWebRTCSignal('WEBRTC_OFFER', offer);

    } catch (err) {
      console.error('Error enabling video:', err);
    }
  };

  const disableVideo = async () => {
    if (localStream) {
      localStream.getTracks().forEach(track => track.stop());
      setLocalStream(null);
      setIsVideoEnabled(false);
    }
    
    if (peerConnection.current) {
      const pc = peerConnection.current;
      // Remove all local tracks from the connection
      pc.getSenders().forEach(sender => {
        if (sender.track) {
          pc.removeTrack(sender);
        }
      });
      
      // Renegotiate without local tracks
      try {
        const offer = await pc.createOffer();
        await pc.setLocalDescription(offer);
        chatSocketService.sendWebRTCSignal('WEBRTC_OFFER', offer);
      } catch (err) {
        console.error('Error renegotiating after disabling video:', err);
      }
    }
  };

  return {
    localStream,
    remoteStream,
    isVideoEnabled,
    startCall,
    enableVideo,
    stopCall: disableVideo,
    cleanup
  };
}
