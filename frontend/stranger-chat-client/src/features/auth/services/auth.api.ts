import { apiClient } from '@/shared/lib/api-client';
import { AuthResponse, LoginRequest, RegisterRequest } from '../types';

export const authApi = {
  login: async (payload: LoginRequest): Promise<AuthResponse> => {
    const response = await apiClient.post<AuthResponse>('/auth/login', payload);
    return response.data;
  },

  register: async (payload: RegisterRequest): Promise<AuthResponse> => {
    const response = await apiClient.post<AuthResponse>('/auth/register', payload);
    return response.data;
  },
};
