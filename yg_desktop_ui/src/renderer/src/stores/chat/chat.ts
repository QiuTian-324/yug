import { ApiResponse, Get } from '@renderer/api';
import { formatLocalizedSessionTime, formatTimestamp } from '@renderer/utils/time';
import { defineStore } from 'pinia';
import { ref } from 'vue';
import { ChatRecordData, ChatRecordResponse, SessionData, SessionListResponse } from './type';

export const ChatStore = defineStore('chat', () => {
  const isConnected = ref(false);
  const ws = ref<WebSocket | null>(null);
  // 消息列表
  const messages = ref<any[]>([]);
  // 聊天记录列表
  const chatRecordList = ref<ChatRecordData[]>([]);
  // 会话列表
  const sessionList = ref<SessionData[]>([]);
  // 当前选中的会话id
  const selectedSessionId = ref<string>('');
  // 当前选中的会话信息
  const currentSessionInfo = ref<SessionData>({
    user_id: '',
    friend_id: '',
    nickname: '',
    avatar: '',
    unread_num: 0,
    last_msg: '',
    last_msg_at: ''
  });


  const Init = () => {
    getSessionList();
  };

  const sendMessage = (message: any) => {
    if (ws.value && ws.value.readyState === WebSocket.OPEN) {
      ws.value.send(JSON.stringify(message));
    } else {
      console.error('WebSocket 连接未建立或已关闭');
    }
  }

  const closeWebSocket = () => {
    if (ws.value) {
      ws.value.close();
      ws.value = null;
      isConnected.value = false;
    }
  }


  async function getSessionList() {
    const res: ApiResponse<SessionListResponse> = await Get<SessionListResponse>('/chat/session_list');
    // 格式化会话列表中的时间
    sessionList.value = res.data.list.map(session => ({
      ...session,
      last_msg_at: formatLocalizedSessionTime(session.last_msg_at)
    }));
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
    setSessionInfo();
  }

  // 设置当前会话信息
  function setSessionInfo() {
    const session = sessionList.value.find(session => session.friend_id == selectedSessionId.value);
    if (session) {
      currentSessionInfo.value = session;
    }
  }


  function setSelectedSessionId(id: string) {
    selectedSessionId.value = id;
  }

  return {
    ws,
    isConnected,
    messages,
    sessionList,
    chatRecordList,
    selectedSessionId,
    currentSessionInfo,
    getSessionList,
    Init,
    getChatRecord,
    setSelectedSessionId,
    setSessionInfo,
    sendMessage,
    closeWebSocket
  };
}, {
  persist: {
    enabled: true,
    strategies: [
      {
        key: "chatStore",
        storage: localStorage,
        paths: ["sessionList", "selectedSessionId"]
      }
      // { key: 'chatRecordList', storage: localStorage, paths: ['chatRecordList'] },
      // // { key: 'sessionList', storage: localStorage, paths: ['sessionList'] },
      // { key: 'selectedSessionId', storage: localStorage, paths: ['selectedSessionId'] },
      // { key: 'currentSessionInfo', storage: localStorage, paths: ['currentSessionInfo'] }
    ]
  }
});

