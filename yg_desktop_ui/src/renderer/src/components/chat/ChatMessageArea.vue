<template>
  <div class="p-4 overflow-auto w-full chat-container">
    <!-- 显示每条消息 -->
    <div v-for="msg in messages" :key="msg.id" class="mb-6">
      <!-- 当前用户的消息（右侧显示） -->
      <div v-if="msg.isMe === 'true'" class="flex items-start justify-end gap-3">
        <div class="flex flex-col items-end">
          <div class="font-semibold text-gray-700">{{ msg.sender }}</div>
          <div class="message-box bg-blue-500 text-white">
            {{ msg.text }}
          </div>
          <div class="text-xs text-gray-500 mt-1">{{ msg.timestamp }}</div>
        </div>
        <a-avatar shape="square" />
      </div>

      <!-- 其他用户的消息（左侧显示） -->
      <div v-else class="flex items-start gap-3">
        <a-avatar shape="square" />
        <div class="flex flex-col items-start">
          <div class="font-semibold text-gray-700">{{ msg.sender }}</div>
          <div class="message-box bg-gray-200 text-gray-800">
            {{ msg.text }}
          </div>
          <div class="text-xs text-gray-500 mt-1">{{ msg.timestamp }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

// 模拟的消息数据
const messages = ref([
  { id: 1, sender: 'User1', isMe: 'true', text: 'Hello, how are you?', timestamp: '12:00:00' },
  {
    id: 2,
    sender: 'User2',
    isMe: 'false',
    text: 'I am good, thanks! How about you?',
    timestamp: '12:00:30'
  },
  {
    id: 3,
    sender: 'User1',
    isMe: 'true',
    text: "I am doing well. What's up?",
    timestamp: '12:01:00'
  },
  { id: 4, sender: 'User2', isMe: 'false', text: 'Not much, just coding.', timestamp: '12:01:30' },
  {
    id: 5,
    sender: 'User1',
    isMe: 'true',
    text: 'Nice, I am working on the same project.',
    timestamp: '12:02:00'
  }
])
</script>

<style scoped>
.chat-container {
  height: calc(100% - 55px - 90px);
  overflow-y: auto;
  padding-right: 8px;
}

/* 自定义滚动条 */
.chat-container::-webkit-scrollbar {
  width: 6px; /* 调整滚动条的宽度 */
}

.chat-container::-webkit-scrollbar-track {
  background: transparent; /* 滚动条轨道背景 */
}

.chat-container::-webkit-scrollbar-thumb {
  background-color: #b0b0b0; /* 滚动条滑块的颜色 */
  border-radius: 10px; /* 滚动条滑块的圆角 */
}

.chat-container::-webkit-scrollbar-thumb:hover {
  background-color: #888; /* 滑块在悬停时变色 */
}

.message-box {
  padding: 10px;
  border-radius: 10px;
  word-wrap: break-word;
}

.message-box.bg-blue-500 {
  border-top-right-radius: 0;
}

.message-box.bg-gray-200 {
  border-top-left-radius: 0;
}
</style>
