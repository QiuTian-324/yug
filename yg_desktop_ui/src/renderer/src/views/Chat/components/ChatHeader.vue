<template>
  <div
    class="flex justify-between items-center h-[55px] border-b-[0.1px] px-4 border-light-border-1 dark:border-dark-border-1 bg-light-bg-2 dark:bg-dark-bg-2"
  >
    <div class="text-xl font-semibold">{{ chatstore.currentSessionInfo.nickname }}</div>
    <div class="text-xl space-x-4">
      <phone-outlined class="cursor-pointer" />
      <video-camera-outlined class="cursor-pointer" />
      <MoreOutlined class="cursor-pointer" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ChatStore } from '@renderer/stores/chat/chat'
const chatstore = ChatStore()
</script>
