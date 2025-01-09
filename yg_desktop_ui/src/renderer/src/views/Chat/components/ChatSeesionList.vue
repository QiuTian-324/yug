<template>
  <!-- 会话列表 -->
  <div class="h-screen bg-light-bg-2 dark:bg-dark-bg-2 p-4">
    <ul class="w-fullspace-y-2">
      <li
        v-for="conversation in chatStore.sessionList"
        :key="conversation.friend_id"
        class="flex items-center gap-2 p-3 rounded-md cursor-pointer"
        :class="[
          chatStore.selectedSessionId == conversation.friend_id
            ? 'text-light-text-white  bg-light-bg-3 dark:bg-dark-bg-3'
            : 'hover:text-light-hover-text hover:bg-light-hover-bg dark:hover:text-dark-text-2 dark:hover:bg-dark-hover-bg'
        ]"
        @click="selectConversation(conversation.friend_id)"
      >
        <!-- 头像 -->
        <div class="w-[40px] h-[40px]">
          <a-avatar :size="40" :src="conversation.avatar" />
        </div>
        <!-- 会话信息 -->
        <div class="flex-1">
          <!-- 昵称和预览消息 -->
          <div class="flex justify-between items-center">
            <div class="font-bold text-[14px]">
              {{ conversation.nickname }}
            </div>
            <div class="text-[12px]">
              {{ conversation.last_msg_at }}
            </div>
          </div>
          <div class="text-[12px] w-full truncate flex items-center">
            <span class="w-[90%]">
              {{ conversation.last_msg }}
            </span>
            <span
              v-if="conversation.unread_num > 0"
              class="w-[10%] flex items-center justify-center rounded-full bg-common-notification-chat text-common-white"
            >
              {{ conversation.unread_num }}
            </span>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import router from '@renderer/router'
import { ChatStore } from '@renderer/stores/chat/chat'
import { onMounted } from 'vue'

const chatStore = ChatStore()

onMounted(() => {
  // 判断会话列表和当前选择的会话id是否为空
  if (chatStore.sessionList.length != 0 || chatStore.selectedSessionId != '') {
    router.push({ name: 'chatDetail', params: { id: chatStore.selectedSessionId } })
  }
})

// 选择会话
const selectConversation = (friendId: string) => {
  chatStore.setSelectedSessionId(friendId)
  chatStore.setSessionInfo()
  router.push({ name: 'chatDetail', params: { id: friendId } })
}
</script>
