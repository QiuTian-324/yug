import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useChatStore = defineStore('chat', () => {
  const isConnected = ref(false);
  const ws = ref<WebSocket | null>(null);
  const messages = ref<any[]>([]);

  function connectWebSocket(token: string | null) {
    console.log('token:', token);

    if (ws.value) {
      ws.value.close();
    }

    ws.value = new WebSocket(`ws://127.0.0.1:9000/api/chat/ws?token=${token}`);

    ws.value.onopen = () => {
      isConnected.value = true;
      console.log('WebSocket 连接已建立');
    };

    ws.value.onmessage = (event) => {
      const message = JSON.parse(event.data);
      messages.value.push(message);
      console.log('收到消息:', message);
    };

    ws.value.onclose = () => {
      isConnected.value = false;
      console.log('WebSocket 连接已关闭');
    };

    ws.value.onerror = (error) => {
      isConnected.value = false;
      console.error('WebSocket 错误:', error);
    };
  }

  function sendMessage(message: any) {
    if (ws.value && ws.value.readyState === WebSocket.OPEN) {
      ws.value.send(JSON.stringify(message));
    } else {
      console.error('WebSocket 连接未建立或已关闭');
    }
  }

  function closeWebSocket() {
    if (ws.value) {
      ws.value.close();
      ws.value = null;
      isConnected.value = false;
    }
  }

  return {
    isConnected,
    messages,
    connectWebSocket,
    sendMessage,
    closeWebSocket
  };
});