import { defineStore } from 'pinia';
import api from '@/lib/axios';

interface User {
  id: number;
  role: string;
}

const isTokenExpired = (token: string): boolean => {
  try {
    const part = token.split('.')[1];
    if (!part) return true;
    const payload = JSON.parse(atob(part));
    if (!payload.exp) return false;
    const now = Math.floor(Date.now() / 1000);
    return payload.exp < now;
  } catch (e) {
    return true;
  }
};

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: null as User | null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token && !isTokenExpired(state.token),
  },
  actions: {
    setToken(token: string) {
      this.token = token;
      localStorage.setItem('token', token);
    },
    async login(credentials: any) {
      const response = await api.post('/auth/login', credentials);
      this.setToken(response.data.token);
      
      try {
        const payload = JSON.parse(atob(response.data.token.split('.')[1]));
        this.user = { 
            id: payload.user_id, 
            role: payload.role 
        };
      } catch (e) {
        console.error("Failed to decode token", e);
      }
    },
    async register(data: any) {
      await api.post('/auth/register', data);
    },
    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem('token');
    },
    hydrate() {
        if (this.token) {
            if (isTokenExpired(this.token)) {
                this.logout();
                return;
            }
            try {
                const part = this.token.split('.')[1];
                if (!part) throw new Error("Invalid token");
                const payload = JSON.parse(atob(part));
                this.user = { 
                    id: payload.user_id, 
                    role: payload.role 
                };
            } catch (e) {
                this.logout();
            }
        }
    }
  },
});
