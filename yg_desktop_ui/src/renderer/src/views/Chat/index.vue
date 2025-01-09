<template>
  <div
    class="flex w-full h-screen py-4 pr-4 bg-light-bg-1 text-light-text-1 dark:bg-dark-bg-1 dark:text-dark-text-1"
  >
    <div class="flex-grow flex">
      <div
        class="min-w-[300px] mr-2 rounded-xl overflow-hidden h-full bg-light-bg-2 dark:bg-dark-bg-2"
      >
        <ChatList />
      </div>
      <div
        class="flex-grow flex-1 h-full rounded-xl text-light-text-1 dark:text-dark-text-1 overflow-hidden bg-light-bg-2 dark:bg-dark-bg-2"
      >
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ChatStore } from '@renderer/stores/chat/chat'
import { UserStore } from '@renderer/stores/user/user'
import ChatList from '@renderer/views/Chat/components/ChatList.vue'
import { onMounted } from 'vue'

const userStore = UserStore()
const chatStore = ChatStore()

onMounted(() => {
  userStore.Init()
  chatStore.Init()
})
</script>
