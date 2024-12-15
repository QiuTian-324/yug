<template>
  <div class="flex w-screen h-screen bg-light-primary dark:bg-dark-primary">
    <!-- 左侧菜单和会话列表 -->
    <div class="w-20 h-screen text-light-text dark:text-dark-text flex">
      <ChatSidebar />
    </div>
    <div class="flex-grow flex pr-4 py-4">
      <div
        class="min-w-[200px] rounded-l-xl overflow-hidden w-[200px] h-full bg-light-base dark:bg-dark-base"
      >
        <ChatList v-if="store.selectedSidebar == 0 || store.selectedSidebar == 1" />
        <ChatContactPerson v-if="store.selectedSidebar == 1" />
      </div>
      <div
        class="flex-grow h-full rounded-r-xl text-light-text dark:text-dark-text overflow-hidden bg-light-base dark:bg-dark-base"
      >
        <div v-if="store.selectedSidebar == 0" class="flex flex-col h-full">
          <ChatHeader />
          <ChatMessageArea />
          <ChatMessageInput class="flex-1" />
        </div>
        <div
          v-if="store.selectedSidebar == 1 && store.selectedSeesionId == 0"
          class="flex flex-col h-full"
        >
          <Nodata />
        </div>
        <div v-if="store.selectedSidebar == 1 && store.selectedSeesionId != 0">
          <ChatUserInfo />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ChatContactPerson from '@renderer/components/chat/ChatContactPerson.vue'
import ChatHeader from '@renderer/components/chat/ChatHeader.vue'
import ChatList from '@renderer/components/chat/ChatList.vue'
import ChatMessageArea from '@renderer/components/chat/ChatMessageArea.vue'
import ChatMessageInput from '@renderer/components/chat/ChatMessageInput.vue'
import ChatSidebar from '@renderer/components/chat/ChatSidebar.vue'
import ChatUserInfo from '@renderer/components/chat/ChatUserInfo.vue'
import Nodata from '@renderer/components/public/nodata.vue'
import { useChatStore } from '@renderer/stores/chat'
import { UserStore } from '@renderer/stores/user'
import { onMounted } from 'vue'

const store = UserStore()
const chatStore = useChatStore()

onMounted(() => {
  store.Init()
  chatStore.Init()
})
</script>
