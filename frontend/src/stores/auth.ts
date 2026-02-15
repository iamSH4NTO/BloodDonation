import { defineStore } from 'pinia';
import api from '@/lib/axios';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
  },
  actions: {
    setToken(token: string) {
      this.token = token;
      localStorage.setItem('token', token);
    },
    async login(credentials: any) {
      const response = await api.post('/auth/login', credentials);
      this.setToken(response.data.token);
    },
    async register(data: any) {
      await api.post('/auth/register', data);
    },
    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem('token');
    },
  },
});
