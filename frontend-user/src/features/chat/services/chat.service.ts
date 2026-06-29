import { Message, ConnectionStatus } from '../types';
import { toast } from '@/shared/components/toast';

type StatusCallback = (status: ConnectionStatus) => void;
type MessageCallback = (message: Message | Message[] | ((prev: Message[]) => Message[])) => void;
type RoomCallback = (roomId: string | null) => void;
type WebRTCCallback = (type: string, payload: unknown) => void;

class ChatSocketService {
  private ws: WebSocket | null = null;
  private token: string | null = null;
  private roomId: string | null = null;
  
  private onStatusChange?: StatusCallback;
  private onMessageReceived?: MessageCallback;
  private onRoomChange?: RoomCallback;
  private onWebRTCSignal?: WebRTCCallback;

  // Set event listeners
  public subscribe(
    onStatusChange: StatusCallback,
    onMessageReceived: MessageCallback,
    onRoomChange: RoomCallback
  ) {
    this.onStatusChange = onStatusChange;
    this.onMessageReceived = onMessageReceived;
    this.onRoomChange = onRoomChange;
  }

  public subscribeWebRTC(onWebRTCSignal: WebRTCCallback) {
    this.onWebRTCSignal = onWebRTCSignal;
  }

  public connect(token: string) {
    this.token = token;
    if (this.ws) {
      this.disconnect();
    }

    const wsUrl = process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:8081/ws/chat';
    const wsEndpoint = `${wsUrl}?token=${token}`;
    const ws = new WebSocket(wsEndpoint);
    this.ws = ws;

    this.onStatusChange?.('connecting');

    ws.onopen = () => {
      if (this.ws !== ws) return;
      this.findMatch();
    };

    ws.onmessage = (event) => {
      if (this.ws !== ws) return;
      try {
        const data = JSON.parse(event.data);

        switch (data.type) {
          case 'MATCHED':
            this.roomId = data.room_id;
            this.onRoomChange?.(this.roomId);
            this.onStatusChange?.('connected');
            this.onMessageReceived?.([]); // Xóa tin nhắn cũ
            toast.success('Đã tìm thấy người lạ! Bắt đầu trò chuyện ngay.', 'Ghép phòng thành công');
            break;
            
          case 'CHAT':
            this.onMessageReceived?.((prev: Message[]) => [
              ...prev,
              {
                id: Date.now().toString() + Math.random(),
                text: data.content,
                sender: 'stranger',
                timestamp: new Date(),
              },
            ]);
            break;

          case 'PARTNER_LEFT':
            this.onStatusChange?.('disconnected');
            this.roomId = null;
            this.onRoomChange?.(null);
            this.onMessageReceived?.((prev: Message[]) => [
              ...prev,
              {
                id: Date.now().toString(),
                text: "Người lạ đã thoát khỏi phòng.",
                sender: 'stranger',
                timestamp: new Date(),
              },
            ]);
            toast.info('Người lạ đã rời khỏi cuộc trò chuyện.', 'Kết thúc');
            break;

          case 'WEBRTC_OFFER':
          case 'WEBRTC_ANSWER':
          case 'WEBRTC_ICE_CANDIDATE':
            this.onWebRTCSignal?.(data.type, data.payload);
            break;
        }
      } catch (err) {
        console.error('Lỗi parse tin nhắn WS:', err);
      }
    };

    ws.onclose = () => {
      if (this.ws !== ws) return;
      this.onStatusChange?.('disconnected');
      this.roomId = null;
      this.onRoomChange?.(null);
    };

    ws.onerror = (error) => {
      if (this.ws !== ws) return;
      console.error('WebSocket Error:', error);
      this.onStatusChange?.('disconnected');
    };
  }

  public disconnect() {
    if (this.ws) {
      if (this.ws.readyState === WebSocket.OPEN) {
        this.ws.send(JSON.stringify({ type: 'LEAVE_ROOM' }));
      }
      this.ws.close();
      this.ws = null;
    }
    this.roomId = null;
    this.onStatusChange?.('disconnected');
    this.onRoomChange?.(null);
  }

  public findMatch() {
    if (this.ws && this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify({ type: 'FIND_MATCH' }));
    }
  }

  public sendMessage(text: string) {
    if (!text.trim() || !this.roomId || !this.ws || this.ws.readyState !== WebSocket.OPEN) return;

    const newMsg: Message = {
      id: Date.now().toString(),
      text: text.trim(),
      sender: 'me',
      timestamp: new Date(),
    };

    this.onMessageReceived?.((prev: Message[]) => [...prev, newMsg]);

    this.ws.send(
      JSON.stringify({
        type: 'CHAT',
        room_id: this.roomId,
        content: text.trim(),
      })
    );
  }

  public sendWebRTCSignal(type: string, payload: unknown) {
    if (!this.roomId || !this.ws || this.ws.readyState !== WebSocket.OPEN) return;
    this.ws.send(
      JSON.stringify({
        type: type,
        room_id: this.roomId,
        payload: payload,
      })
    );
  }

  public getStatus() {
    return this.ws ? this.ws.readyState : WebSocket.CLOSED;
  }
}

export const chatSocketService = new ChatSocketService();
