import { defineStore } from 'pinia';

export type ToastType = 'success' | 'error' | 'info';

interface ToastState {
  visible: boolean;
  message: string;
  type: ToastType;
  timeoutId: number | null;
}

export const useToastStore = defineStore('toast', {
  state: (): ToastState => ({
    visible: false,
    message: '',
    type: 'success',
    timeoutId: null,
  }),
  actions: {
    show(message: string, type: ToastType = 'success', duration: number = 3000) {
      // Clear existing timeout if any
      if (this.timeoutId) {
        clearTimeout(this.timeoutId);
      }

      this.message = message;
      this.type = type;
      this.visible = true;

      this.timeoutId = window.setTimeout(() => {
        this.visible = false;
      }, duration);
    },
    hide() {
      this.visible = false;
      if (this.timeoutId) {
        clearTimeout(this.timeoutId);
        this.timeoutId = null;
      }
    }
  },
});
