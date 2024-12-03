<template>
  <div class="flex w-screen h-screen bg-[#dbf3ff]">
    <!-- 左侧菜单和会话列表 -->
    <div class="w-20 h-scree text-black flex">
      <ChatSidebar />
    </div>
    <div class="flex-grow flex pr-4 py-4">
      <div class="min-w-[200px] rounded-l-xl overflow-hidden w-[200px] h-full bg-white">
        <ChatList v-if="store.selectedSidebar == 0 || store.selectedSidebar == 1" />
        <ChatContactPerson v-if="store.selectedSidebar == 1" />
      </div>
      <div class="flex-grow h-full rounded-r-xl text-black overflow-hidden bg-white">
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
import { onMounted } from 'vue'
import ChatContactPerson from '../components/chat/ChatContactPerson.vue'
import ChatHeader from '../components/chat/ChatHeader.vue'
import ChatList from '../components/chat/ChatList.vue'
import ChatMessageArea from '../components/chat/ChatMessageArea.vue'
import ChatMessageInput from '../components/chat/ChatMessageInput.vue'
import ChatSidebar from '../components/chat/ChatSidebar.vue'
import ChatUserInfo from '../components/chat/ChatUserInfo.vue'
import Nodata from '../components/public/nodata.vue'
import { UserStore } from '../stores/user'
const store = UserStore()

onMounted(() => {
  store.Init()
})
</script>
