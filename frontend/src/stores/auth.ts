import { defineStore } from 'pinia';
import api from '@/lib/axios';

interface User {
  id: number;
  role: string;
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('token') || null,
    user: null as User | null,
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
      
      // Decode token to get role (simple implementation, ideally backend returns user object)
      // For now, let's assume backend returns user object or we fetch it
      try {
        const profileRes = await api.get('/profile');
        // We need to fetch the role from the user object, but profile endpoint returns DonorProfile
        // Let's rely on the token payload if possible, or update login response
        // Actually, let's update login response in backend to return role, or fetch it.
        // For now, let's assume the token has the role and we decode it, OR we fetch a new /auth/me endpoint.
        // But we don't have /auth/me. 
        // Let's use a workaround: The backend middleware sets 'role' in context from token.
        // The token is a JWT. We can decode it here.
        const payload = JSON.parse(atob(response.data.token.split('.')[1]));
        this.user = { 
            id: payload.user_id, 
            role: payload.role 
        };
      } catch (e) {
        console.error("Failed to decode token or fetch user", e);
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
    // Add a hydrate action to restore user state on app load
    hydrate() {
        if (this.token) {
            try {
                // simple jwt decode
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
