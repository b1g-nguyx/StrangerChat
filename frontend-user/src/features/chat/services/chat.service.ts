import { Message, ConnectionStatus } from '../types';

type StatusCallback = (status: ConnectionStatus) => void;
type MessageCallback = (message: Message | Message[] | ((prev: Message[]) => Message[])) => void;
type RoomCallback = (roomId: string | null) => void;

class ChatSocketService {
  private ws: WebSocket | null = null;
  private token: string | null = null;
  private roomId: string | null = null;
  
  private onStatusChange?: StatusCallback;
  private onMessageReceived?: MessageCallback;
  private onRoomChange?: RoomCallback;

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

  public connect(token: string) {
    this.token = token;
    if (this.ws) {
      this.disconnect();
    }

    const wsUrl = process.env.NEXT_PUBLIC_WS_URL || 'ws://localhost:8081/ws/chat';
    const wsEndpoint = `${wsUrl}?token=${token}`;
    this.ws = new WebSocket(wsEndpoint);

    this.onStatusChange?.('connecting');

    this.ws.onopen = () => {
      this.findMatch();
    };

    this.ws.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);

        switch (data.type) {
          case 'MATCHED':
            this.roomId = data.room_id;
            this.onRoomChange?.(this.roomId);
            this.onStatusChange?.('connected');
            this.onMessageReceived?.([]); // Xóa tin nhắn cũ
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
            break;
        }
      } catch (err) {
        console.error('Lỗi parse tin nhắn WS:', err);
      }
    };

    this.ws.onclose = () => {
      this.onStatusChange?.('disconnected');
      this.roomId = null;
      this.onRoomChange?.(null);
    };

    this.ws.onerror = (error) => {
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

  public getStatus() {
    return this.ws ? this.ws.readyState : WebSocket.CLOSED;
  }
}

export const chatSocketService = new ChatSocketService();
