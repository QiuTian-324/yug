<template>
  <!-- 会话列表 -->
  <div class="h-screen bg-light-base">
    <ul class="w-full text-light-text">
      <li
        v-for="conversation in chatStore.sessionList"
        :key="conversation.friend_id"
        :class="[
          'flex items-center gap-3 p-2 mx-2 my-2 rounded-lg cursor-pointer',
          chatStore.selectedSessionId === conversation.friend_id
            ? 'text-white bg-light-active'
            : 'hover:text-light-text hover:bg-light-hover'
        ]"
        @click="selectConversation(conversation.friend_id)"
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
              {{ conversation.last_msg }}
            </div>
            <div
              :class="[
                'text-[11px]',
                chatStore.selectedSessionId === conversation.friend_id ? 'text-white' : 'text-gray-500'
              ]"
            >
              {{ formatTimestamp(conversation.last_msg_at) }}
            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { useChatStore } from '@renderer/stores/chat';


// 使用 chatStore 中的 sessionList
const chatStore = useChatStore()


// 选择会话
const selectConversation = (friendId: number) => {
  chatStore.setSelectedSessionId(friendId)
  chatStore.getChatRecord()
}

// 格式化时间戳
const formatTimestamp = (timestamp: string) => {
  const date = new Date(timestamp)
  return date.toLocaleString()
}
</script>
