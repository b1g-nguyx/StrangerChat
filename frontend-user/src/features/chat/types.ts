export type ConnectionStatus = 'disconnected' | 'connecting' | 'connected';

export interface Message {
  id: string;
  text: string;
  sender: 'me' | 'stranger';
  timestamp: Date;
}
