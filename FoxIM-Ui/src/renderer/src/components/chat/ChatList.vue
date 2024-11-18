<template>
  <a-modal
    v-model:visible="store.addFriendModalVisible"
    @ok="store.addFriend"
    @cancel="store.addFriendModalVisible = false"
  >
    <template #title>添加好友 </template>
    <div>
      <AddFriend />
    </div>
  </a-modal>

  <div class="h-full w-full bg-white border-r">
    <!-- 搜索框和添加按钮的容器 -->
    <div class="flex py-2 border-b justify-center items-center h-[55px] w-full">
      <!-- 搜索框 -->
      <a-input
        v-model:value="searchText"
        placeholder="搜索会话..."
        class="h-[35px] w-[180px] rounded-2xl"
      >
        <template #suffix>
          <icon-user-add
            class="cursor-pointer"
            size="20"
            @click="store.addFriendModalVisible = true"
          />
        </template>
      </a-input>
    </div>

    <!-- 会话列表 -->
    <div class="overflow-auto h-[calc(100%-55px)]">
      <ul class="w-full text-black">
        <li
          v-for="conversation in filteredConversations"
          :key="conversation.sessionId"
          @click="selectConversation(conversation.sessionId)"
          :class="[
            'flex items-center gap-3 p-2 mx-2 my-2 rounded-lg cursor-pointer',
            selectedSessionId === conversation.sessionId
              ? 'text-white bg-[rgb(var(--arcoblue-5))]'
              : 'hover:text-black hover:bg-[rgb(var(--arcoblue-1))]'
          ]"
        >
          <!-- 头像 -->
          <div class="w-[40px] h-[40px]">
            <a-avatar :src="conversation.avatar" />
          </div>
          <!-- 会话信息 -->
          <div class="flex-1 flex items-center justify-between">
            <!-- 昵称和预览消息 -->
            <div>
              <div class="font-bold">{{ conversation.nickname }}</div>
              <div class="text-[12px] w-[120px] truncate">
                {{ conversation.messagePreview }}
              </div>
              <div class="text-[10px]">{{ conversation.timestamp }}</div>
            </div>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '@renderer/stores/user'
import { computed, onMounted, ref } from 'vue'
import AddFriend from './AddFriend.vue'
const store = UserStore()

// 会话数据接口
interface Conversation {
  sessionId: number
  avatar: string
  nickname: string
  messagePreview: string
  timestamp: string
}

// 模拟会话数据
const conversations: Conversation[] = [
  {
    sessionId: 1,
    avatar: 'https://i.pravatar.cc/100?img=1',
    nickname: 'Alice',
    messagePreview: 'Hey! Are we still meeting later?',
    timestamp: '11:45 AM'
  },
  {
    sessionId: 2,
    avatar: 'https://i.pravatar.cc/100?img=2',
    nickname: 'Bob',
    messagePreview: 'I’ve sent the files to your email.',
    timestamp: '10:30 AM'
  },
  {
    sessionId: 3,
    avatar: 'https://i.pravatar.cc/100?img=3',
    nickname: 'Charlie',
    messagePreview: 'Can you review this document?',
    timestamp: 'Yesterday'
  },
  {
    sessionId: 4,
    avatar: 'https://i.pravatar.cc/100?img=4',
    nickname: 'David',
    messagePreview: 'Great job on the presentation!',
    timestamp: 'Yesterday'
  },
  {
    sessionId: 5,
    avatar: 'https://i.pravatar.cc/100?img=5',
    nickname: 'Eva',
    messagePreview: 'Let’s catch up this weekend.',
    timestamp: '2 days ago'
  }
]

// 搜索框绑定的文本
const searchText = ref('')

// 选中的会话ID
const selectedSessionId = ref<number | null>(null)

// 筛选会话列表
const filteredConversations = computed(() =>
  conversations.filter((conversation) =>
    conversation.nickname.toLowerCase().includes(searchText.value.toLowerCase())
  )
)
// 选择会话
const selectConversation = (sessionId: number) => {
  selectedSessionId.value = sessionId
}


</script>

<style scoped>
/* 调整列表的高度和样式 */
.overflow-auto {
  height: calc(100% - 55px);
}
</style>
