<template>
  <div class="flex w-screen h-screen bg-[rgb(var(--arcoblue-2))]">
    <!-- 左侧菜单和会话列表 -->
    <div class="w-20 h-scree text-black flex">
      <ChatSidebar />
    </div>
    <div class="flex-grow flex pr-4 py-4">
      <div class="min-w-[200px] rounded-l-xl overflow-hidden w-[200px] h-full">
        <ChatList />
      </div>
      <div class="flex-grow h-full rounded-r-xl overflow-hidden bg-white">
        <ChatHeader />
        <ChatMessageArea />
        <ChatMessageInput class="flex-1" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ChatHeader from '../components/chat/ChatHeader.vue'
import ChatList from '../components/chat/ChatList.vue'
import ChatMessageArea from '../components/chat/ChatMessageArea.vue'
import ChatMessageInput from '../components/chat/ChatMessageInput.vue'
import ChatSidebar from '../components/chat/ChatSidebar.vue'
</script>
