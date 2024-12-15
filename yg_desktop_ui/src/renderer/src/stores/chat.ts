import { ApiResponse, Get } from '@renderer/api';
import { ChatRecordData, ChatRecordResponse, SessionData, SessionListResponse } from '@renderer/interface/chat';
import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useChatStore = defineStore('chat', () => {
  const isConnected = ref(false);
  const ws = ref<WebSocket | null>(null);
  // 消息列表
  const messages = ref<any[]>([]);
  // 聊天记录列表
  const chatRecordList = ref<ChatRecordData[]>([]);
  // 会话列表
  const sessionList = ref<SessionData[]>([]);
  // 当前选中的会话id
  const selectedSessionId = ref<number>(0);
  // 当前选中的会话信息
  const sessionInfo = ref<SessionData | null>(null);
  const Init = () => {
    getSessionList();
  };

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
      console.log('Received message:', event.data);
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

  // 获取会话列表
  async function getSessionList() {
    const res: ApiResponse<SessionListResponse> = await Get<SessionListResponse>('/chat/session_list');
    console.log('会话列表:', res.data.list);
    sessionList.value = res.data.list;
  }

  // 获取聊天记录
  async function getChatRecord() {
    const res: ApiResponse<ChatRecordResponse> = await Get<ChatRecordResponse>('/chat/chat_record', { friend_id: selectedSessionId.value });
    console.log('聊天记录:', res.data.list);

    // 格式化时间
    chatRecordList.value = res.data.list.map(record => ({
      ...record,
      created_at: formatTimestamp(record.created_at),
      updated_at: formatTimestamp(record.updated_at)
    }));
    // 获取会话信息
    getSessionInfo();
  }

  // 通过当前选中的会话id，获取sessionlist中的用户信息
  function getSessionInfo() {
    const session = sessionList.value.find(session => session.friend_id === selectedSessionId.value);
    if (session) {
      sessionInfo.value = session;
    }
  }

  // 格式化时间戳
  function formatTimestamp(timestamp: string): string {
    const date = new Date(timestamp);
    return new Intl.DateTimeFormat('zh-CN', {
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit'
    }).format(date);
  }

  function setSelectedSessionId(id: number) {
    selectedSessionId.value = id;
  }

  return {
    isConnected,
    messages,
    sessionList,
    chatRecordList,
    selectedSessionId,
    sessionInfo,
    connectWebSocket,
    sendMessage,
    closeWebSocket,
    getSessionList,
    Init,
    getChatRecord,
    setSelectedSessionId,
    getSessionInfo
  };
}, {
  persist: {
    enabled: true,
    strategies: [
      { key: 'chatRecordList', storage: localStorage, paths: ['chatRecordList'] },
      // { key: 'sessionList', storage: localStorage, paths: ['sessionList'] },
      { key: 'selectedSessionId', storage: localStorage, paths: ['selectedSessionId'] },
      { key: 'sessionInfo', storage: localStorage, paths: ['sessionInfo'] }
    ]
  }
});

