export interface User {
  id: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string | null;
  username: string;
  email: string;
  display_name: string;
  avatar_url?: string;
  status: 'online' | 'offline' | 'idle';
  current_room_id?: string | null;
  is_banned: boolean;
  banned_at?: string | null;
}

export interface AuthResponse {
  message: string;
  data: User;
  access_token: string;
  refresh_token: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface ApiErrorResponse {
  error: string;
}
