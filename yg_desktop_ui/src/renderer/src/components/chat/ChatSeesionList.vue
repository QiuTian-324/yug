<template>
  <!-- 会话列表 -->
  <div class="h-screen bg-white">
    <ul class="w-full text-black">
      <li
        v-for="conversation in conversations"
        :key="conversation.sessionId"
        :class="[
          'flex items-center gap-3 p-2 mx-2 my-2 rounded-lg cursor-pointer',
          selectedSessionId === conversation.sessionId
            ? 'text-white bg-blue-500'
            : 'hover:text-black hover:bg-gray-100'
        ]"
        @click="selectConversation(conversation.sessionId)"
      >
        <!-- 头像 -->
        <div class="w-[40px] h-[40px]">
          <a-avatar shape="square" :size="40" :src="conversation.avatar" />
        </div>
        <!-- 会话信息 -->
        <div class="flex-1 flex items-center justify-between">
          <!-- 昵称和预览消息 -->
          <div>
            <div class="font-bold text-[14px]">{{ conversation.nickname }}</div>
            <div class="text-[12px] w-[120px] truncate">
              {{ conversation.messagePreview }}
            </div>
            <div
              :class="[
                'text-[11px]',
                selectedSessionId === conversation.sessionId ? 'text-white' : 'text-gray-500'
              ]"
            >
              {{ conversation.timestamp }}
            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { Avatar as AAvatar } from 'ant-design-vue'
import { ref } from 'vue'

// 会话数据接口
interface Conversation {
  sessionId: number
  avatar: string
  nickname: string
  messagePreview: string
  timestamp: string
}

// 模拟会话数据
const conversations: Conversation[] = [
  {
    sessionId: 1,
    avatar: 'https://i.pravatar.cc/100?img=1',
    nickname: 'Alice',
    messagePreview: 'Hey! Are we still meeting later?',
    timestamp: '11:45 AM'
  },
  {
    sessionId: 2,
    avatar: 'https://i.pravatar.cc/100?img=2',
    nickname: 'Bob',
    messagePreview: 'I’ve sent the files to your email.',
    timestamp: '10:30 AM'
  },
  {
    sessionId: 3,
    avatar: 'https://i.pravatar.cc/100?img=3',
    nickname: 'Charlie',
    messagePreview: 'Can you review this document?',
    timestamp: 'Yesterday'
  },
  {
    sessionId: 4,
    avatar: 'https://i.pravatar.cc/100?img=4',
    nickname: 'David',
    messagePreview: 'Great job on the presentation!',
    timestamp: 'Yesterday'
  },
  {
    sessionId: 5,
    avatar: 'https://i.pravatar.cc/100?img=5',
    nickname: 'Eva',
    messagePreview: 'Let’s catch up this weekend.',
    timestamp: '2 days ago'
  }
]

// 选中的会话ID
const selectedSessionId = ref<number | null>(null)

// 选择会话
const selectConversation = (sessionId: number) => {
  selectedSessionId.value = sessionId
}
</script>
