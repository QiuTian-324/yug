<template>
  <div class="p-4 overflow-auto w-full chat-container">
    <!-- 显示每条消息 -->
    <div v-for="msg in chatStore.chatRecordList" :key="msg.id" class="mb-6">
      <!-- 当前用户的消息（右侧显示） -->
      <div
        v-if="msg.sender_id !== chatStore.selectedSessionId"
        class="flex items-start justify-end gap-3"
      >
      
        <div class="flex flex-col items-end">
          <div class="font-semibold text-light-text">{{ userStore.userInfo?.nickname }}</div>
          <div class="message-box bg-light-message_me text-gray-800">
            <!-- 文本消息 -->
            <div v-if="msg.content_type === 0">
              {{ msg.content }}
            </div>
            <!-- 图片消息 -->
            <div v-else-if="msg.content_type === 1">
              <img :src="msg.url" alt="图片消息" class="max-w-xs rounded" />
            </div>
            <!-- 文件消息 -->
            <div v-else-if="msg.content_type === 2">
              <a :href="msg.url" target="_blank" class="text-blue-500 underline">
                {{ msg.file_name }}
              </a>
            </div>
          </div>
          <div class="text-xs text-light-text_secondary mt-1">
            {{ msg.created_at }}
          </div>
        </div>
        <a-avatar shape="square" :src="userStore.userInfo?.avatar_url" />
      </div>

      <!-- 其他用户的消息（左侧显示） -->
      <div v-else class="flex items-start gap-3">
        <a-avatar shape="square" :src="chatStore.sessionInfo?.avatar" />
        <div class="flex flex-col items-start">
          <div class="font-semibold text-light-text">{{ chatStore.sessionInfo?.nickname }}</div>
          <div class="message-box bg-light-message_other text-gray-800">
            <!-- 文本消息 -->
            <div v-if="msg.content_type === 0">
              {{ msg.content }}
            </div>
            <!-- 图片消息 -->
            <div v-else-if="msg.content_type === 1">
              <img :src="msg.url" alt="图片消息" class="max-w-xs rounded" />
            </div>
            <!-- 文件消息 -->
            <div v-else-if="msg.content_type === 2">
              <a :href="msg.url" target="_blank" class="text-blue-500 underline">
                {{ msg.file_name }}
              </a>
            </div>
          </div>
          <div class="text-xs text-light-text_secondary mt-1">
            {{ msg.created_at }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useChatStore } from '@renderer/stores/chat'
import { UserStore } from '@renderer/stores/user'

const chatStore = useChatStore()
const userStore = UserStore()
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

.message-box.bg-light-message_me {
  border-top-right-radius: 0;
}

.message-box.bg-light-message_other {
  border-top-left-radius: 0;
}

img {
  max-width: 100%;
  border-radius: 10px;
}
</style>
