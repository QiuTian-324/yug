<template>
  <div class="overflow-hidden overflow-y-auto w-full p-4 h-[calc(100vh-12rem)] custom-scrollbar">
    <!-- 显示每条消息 -->
    <div v-for="msg in chatStore.chatRecordList" :key="msg.id" class="mb-6">
      <!-- 当前用户的消息（右侧显示） -->
      <div
        v-if="msg.sender_id != chatStore.selectedSessionId"
        class="flex items-start justify-end gap-3"
      >
        <div class="flex flex-col items-end">
          <div class="font-semibold text-light-text">{{ userStore.userInfo.nickname }}</div>
          <div
            class="message-box bg-light-message-bubble-me dark:bg-dark-message-bubble-me p-2 rounded-lg rounded-tr-none break-words"
          >
            <!-- 文本消息 -->
            <div v-if="msg.content_type === 0" class="text-common-white">
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
          <div class="text-xs text-light-text-3 mt-1">
            {{ msg.created_at }}
          </div>
        </div>
        <a-avatar :src="userStore.userInfo?.avatar_url" @click="openUserInfo" />
      </div>

      <!-- 其他用户的消息（左侧显示） -->
      <div v-else class="flex items-start gap-3">
        <a-avatar :src="chatStore.currentSessionInfo.avatar" @click="openUserInfo" />
        <div class="flex flex-col items-start">
          <div class="font-semibold text-light-text">
            {{ chatStore.currentSessionInfo.nickname }}
          </div>
          <div
            class="message-box bg-light-message-bubble-other dark:bg-dark-message-bubble-other p-2 rounded-lg rounded-tl-none break-words"
          >
            <!-- 文本消息 -->
            <div v-if="msg.content_type === 0" class="text-common-black dark:text-common-white">
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
          <div class="text-xs text-light-text-3 mt-1">
            {{ msg.created_at }}
          </div>
        </div>
      </div>
    </div>
  </div>

  <!-- 用户信息弹窗 -->
  <a-modal v-model:visible="isUserInfoModalVisible" title="用户信息" :footer="null">
    <div class="flex flex-col items-center justify-center p-6">
      <!-- 背景图作为容器的背景 -->
      <div
        class="relative w-full h-[150px] rounded-lg overflow-hidden flex justify-center items-end bg-cover bg-center z-0"
        :style="{
          backgroundImage: `url(${mockUserInfo.backgroundImage})`
        }"
      ></div>
      <div class="w-full h-full flex items-center justify-start gap-4 m-4">
        <div class="text-sm">
          <IconFont :type="mockUserInfo.gender" class="text-[80px]" />
        </div>
        <!-- 用户信息 -->
        <div class="flex flex-col items-start space-y-2">
          <div class="font-semibold text-xl flex w-full items-center gap-2">
            <span>
              {{ mockUserInfo.nickname }}
            </span>
            <a-dropdown>
              <template #default>
                <a-tooltip :title="getMoodIconTitle(mockUserInfo.moodIcon)">
                  <IconFont :type="mockUserInfo.moodIcon" />
                </a-tooltip>
              </template>
              <template #overlay>
                <div class="icon-menu">
                  <a-tooltip v-for="icon in moodIcons" :key="icon" :title="getMoodIconTitle(icon)">
                    <div class="icon-item">
                      <IconFont :type="icon" @click="selectMoodIcon(icon)" />
                    </div>
                  </a-tooltip>
                </div>
              </template>
            </a-dropdown>
          </div>
          <div class="text-sm">{{ mockUserInfo.address }}</div>
          <div class="text-sm">{{ mockUserInfo.signature }}</div>
        </div>
      </div>
    </div>
  </a-modal>
</template>

<script setup lang="ts">
import { ChatStore } from '@renderer/stores/chat/chat'
import { UserStore } from '@renderer/stores/user/user'
import { ref } from 'vue'

const chatStore = ChatStore()
const userStore = UserStore()
const isUserInfoModalVisible = ref(false)
const openUserInfo = () => {
  isUserInfoModalVisible.value = true
}

// 模拟用户信息数据
const mockUserInfo = ref({
  nickname: '张三',
  gender: 'icon-a-17nanxingjiaose',
  address: '北京市朝阳区',
  signature: '生活是一场旅行',
  mood: '开心',
  backgroundImage:
    'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwzNjUyOXwwfDF8c2VhcmNofDJ8fGJlYWNofGVufDB8fHx8MTY4NDA3NjQ3MA&ixlib=rb-1.2.1&q=80&w=1920',
  moodIcon: 'icon-kaixin'
})

const isMoodIconMenuVisible = ref(false)
const moodIcons = [
  'icon-kaixin',
  'icon-emoji-22',
  'icon--tired',
  'icon-bukaixin',
  'icon-075-hypnotized-1',
  'icon-jiaolv_1',
  'icon-emoji-25',
  'icon-a-dakubeishangshangxinbiaoqingxiaolian',
  'icon-kaixin',
  'icon-emoji-22',
  'icon--tired',
  'icon-bukaixin',
  'icon-075-hypnotized-1',
  'icon-jiaolv_1',
  'icon-emoji-25',
  'icon-a-dakubeishangshangxinbiaoqingxiaolian'
] // 示例图标


const selectMoodIcon = (icon: string) => {
  // 通过icon查找getMoodIconTitle中的title
  const title = getMoodIconTitle(icon)
  console.log(title)
  mockUserInfo.value.mood = title
  mockUserInfo.value.moodIcon = icon
}

const getMoodIconTitle = (icon: string) => {
  const titles: Record<string, string> = {
    'icon-kaixin': '开心',
    'icon-emoji-22': '微笑',
    'icon--tired': '疲惫',
    'icon-bukaixin': '不开心',
    'icon-075-hypnotized-1': '催眠',
    'icon-jiaolv_1': '焦虑',
    'icon-emoji-25': '惊讶',
    'icon-a-dakubeishangshangxinbiaoqingxiaolian': '大哭'
  }
  return titles[icon] || '心情'
}
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  display: none; /* 隐藏滚动条 */
}

.custom-scrollbar {
  -ms-overflow-style: none; /* IE 和 Edge */
  scrollbar-width: none; /* Firefox */
}

.icon-menu {
  background-color: white;
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 8px;
  display: flex;
  flex-wrap: wrap; /* 允许换行 */
  gap: 4px;
  z-index: 10;
  max-width: 320px; /* 控制每行最多10个图标 */
}

.icon-item {
  width: 30px; /* 每个图标的宽度 */
  height: 30px; /* 每个图标的高度 */
  display: flex;
  font-size: 20px;
  align-items: center;
  justify-content: center;
}
</style>
