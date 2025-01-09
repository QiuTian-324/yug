import { ChatStore } from '@renderer/stores/chat/chat';

export function connectWebSocket(token: string) {
  const chatStore = ChatStore();

  if (chatStore.ws) {
    chatStore.ws.close();
  }
  chatStore.ws = new WebSocket(`ws://127.0.0.1:9000/api/chat/ws?token=${token}`);

  chatStore.ws.onopen = () => {
    chatStore.isConnected = true;
    console.log('WebSocket 连接已建立');
  };

  chatStore.ws.onmessage = (event) => {
    console.log('Received message:', event.data);
    const message = JSON.parse(event.data);
    console.log('收到消息:', message);
  };

  chatStore.ws.onclose = () => {
    chatStore.isConnected = false;
    console.log('WebSocket 连接已关闭');
  };

  chatStore.ws.onerror = (error) => {
    chatStore.isConnected = false;
    console.error('WebSocket 错误:', error);
  };
}
