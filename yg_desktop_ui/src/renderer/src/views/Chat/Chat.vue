<template>
  <div class="flex flex-col h-full">
    <ChatHeader />
    <ChatMessageArea />
    <a-style-provider hash-priority="high">
      <ChatMessageInput class="flex-1" />
    </a-style-provider>
  </div>
</template>

<script setup lang="ts">
import { ChatStore } from '@renderer/stores/chat/chat'
import ChatHeader from '@renderer/views/Chat/components/ChatHeader.vue'
import ChatMessageArea from '@renderer/views/Chat/components/ChatMessageArea.vue'
import ChatMessageInput from '@renderer/views/Chat/components/ChatMessageInput.vue'
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
const chatStore = ChatStore()
const route = useRoute()
// 获取路由id
onMounted(() => {
  const id = route.params.id as string
  chatStore.setSelectedSessionId(id)
  chatStore.getChatRecord()
})
</script>

<style scoped></style>
