import axios from 'axios';

// Get base URL from env or fallback to default
const baseURL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/v1';

export const apiClient = axios.create({
  baseURL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to attach Bearer token
apiClient.interceptors.request.use(
  (config) => {
    // We will retrieve token from localStorage or Zustand state
    const token = typeof window !== 'undefined' ? localStorage.getItem('access_token') : null;
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Response interceptor to handle errors globally
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    // If the server returns a 400 with an "error" field as defined in the contract
    if (error.response && error.response.data && error.response.data.error) {
      return Promise.reject(new Error(error.response.data.error));
    }
    // Generic error fallback
    return Promise.reject(error);
  }
);
