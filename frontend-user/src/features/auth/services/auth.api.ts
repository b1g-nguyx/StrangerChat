import { apiClient } from '@/shared/lib/api-client';
import { AuthData, APIResponse, LoginRequest, RegisterRequest, User } from '../types';

export const authApi = {
  login: async (payload: LoginRequest): Promise<AuthData> => {
    const response = await apiClient.post<APIResponse<AuthData>>('/auth/login', payload);
    return response.data.data;
  },

  register: async (payload: RegisterRequest): Promise<AuthData> => {
    const response = await apiClient.post<APIResponse<AuthData>>('/auth/register', payload);
    return response.data.data;
  },

  logout: async (): Promise<void> => {
    await apiClient.post('/auth/logout');
  },

  refreshToken: async (): Promise<string> => {
    // Gọi trực tiếp để xin cấp lại token bằng HttpOnly cookie
    const response = await apiClient.post<APIResponse<AuthData>>('/auth/refresh');
    const newAccessToken = response.data.data.access_token;
    
    // Lưu token mới vào Zustand & LocalStorage
    import('@/features/auth/store/auth.store').then(module => {
      module.useAuthStore.getState().setAccessToken(newAccessToken);
    });
    
    return newAccessToken;
  },
};
