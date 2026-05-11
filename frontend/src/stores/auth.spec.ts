/**
 * @vitest-environment jsdom
 */
import { describe, it, expect, beforeEach, vi } from 'vitest';
import { setActivePinia, createPinia } from 'pinia';
import { useAuthStore } from './auth';
import { Buffer } from 'buffer';

// Mock localStorage
const localStorageMock = (function() {
  let store: Record<string, string> = {};
  return {
    getItem: (key: string) => store[key] || null,
    setItem: (key: string, value: string) => store[key] = value.toString(),
    removeItem: (key: string) => delete store[key],
    clear: () => { store = {}; },
    key: (index: number) => Object.keys(store)[index] || null,
    length: 0
  };
})();

vi.stubGlobal('localStorage', localStorageMock);

// Mock atob/btoa if needed using Buffer
if (typeof btoa === 'undefined') {
    vi.stubGlobal('btoa', (str: string) => Buffer.from(str, 'binary').toString('base64'));
}
if (typeof atob === 'undefined') {
    vi.stubGlobal('atob', (str: string) => Buffer.from(str, 'base64').toString('binary'));
}

describe('Auth Store', () => {
  beforeEach(() => {
    setActivePinia(createPinia());
    localStorage.clear();
  });

  const createToken = (exp: number) => {
    const header = btoa(JSON.stringify({ alg: 'HS256', typ: 'JWT' }));
    const payload = btoa(JSON.stringify({ user_id: 1, role: 'donor', exp }));
    return `${header}.${payload}.signature`;
  };

  it('should be authenticated if token is valid and not expired', () => {
    const authStore = useAuthStore();
    const futureDate = Math.floor(Date.now() / 1000) + 3600;
    const token = createToken(futureDate);
    
    authStore.setToken(token);
    expect(authStore.isAuthenticated).toBe(true);
  });

  it('should not be authenticated if token is expired', () => {
    const authStore = useAuthStore();
    const pastDate = Math.floor(Date.now() / 1000) - 3600;
    const token = createToken(pastDate);
    
    authStore.setToken(token);
    expect(authStore.isAuthenticated).toBe(false);
  });

  it('should not be authenticated if token is invalid', () => {
    const authStore = useAuthStore();
    authStore.setToken('invalid.token');
    expect(authStore.isAuthenticated).toBe(false);
  });

  it('should logout and clear state', () => {
    const authStore = useAuthStore();
    authStore.setToken('some.token');
    authStore.logout();
    
    expect(authStore.token).toBeNull();
    expect(authStore.user).toBeNull();
    expect(localStorage.getItem('token')).toBeNull();
  });

  it('should hydrate user from valid token', () => {
    const authStore = useAuthStore();
    const futureDate = Math.floor(Date.now() / 1000) + 3600;
    const token = createToken(futureDate);
    
    localStorage.setItem('token', token);
    authStore.token = token;
    authStore.hydrate();
    
    expect(authStore.user).toEqual({ id: 1, role: 'donor' });
  });

  it('should clear state on hydrate if token is expired', () => {
    const authStore = useAuthStore();
    const pastDate = Math.floor(Date.now() / 1000) - 3600;
    const token = createToken(pastDate);
    
    localStorage.setItem('token', token);
    authStore.token = token;
    authStore.hydrate();
    
    expect(authStore.token).toBeNull();
    expect(authStore.user).toBeNull();
  });
});
