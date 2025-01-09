<template>
  <a-modal
    v-model:open="store.addFriendModalVisible"
    @ok="store.addFriend"
    @cancel="store.addFriendModalVisible = false"
  >
    <template #title>{{ $t('user.add_friend') }}</template>
    <div>
      <AddFriend />
    </div>
    <template #footer>
      <a-button @click="store.addFriendModalVisible = false">{{ $t('common.cancel') }}</a-button>
      <a-button type="primary" @click="store.addFriend">{{ $t('common.save') }}</a-button>
    </template>
  </a-modal>

  <div class="h-full w-full flex flex-col items-center bg-light-bg-2 dark:bg-dark-bg-2">
    <!-- 搜索框和添加按钮的容器 -->
    <div class="flex p-4 gap-2 justify-center items-center h-[55px] w-full">
      <a-style-provider hash-priority="high">
        <a-input
          :placeholder="$t('user.search_conversation')"
          v-model:value="searchText"
          class="h-[35px] flex-1 rounded-2xl bg-light-bg-2 dark:bg-dark-bg-2 border-light-border-4 dark:border-dark-border-4 text-light-text-1 dark:text-dark-text-1"
        >
          <template #prefix>
            <SearchOutlined />
          </template>
          <template #suffix>
            <!-- 添加好友按钮 -->
            <UserAddOutlined
              class="cursor-pointer flex justify-center items-center text-light-text-1 ml-2 dark:text-dark-text-1"
              @click="store.addFriendModalVisible = true"
            />
          </template>
        </a-input>
      </a-style-provider>
    </div>

    <!-- chat list -->

    <div class="w-full flex-1 flex flex-col">
      <!-- 选择按钮 -->
      <div class="w-full h-[50px] p-4 flex justify-center items-center">
        <div
          class="flex justify-evenly items-center w-full h-8 bg-light-bg-3 dark:bg-dark-bg-3 rounded-md relative"
        >
          <!-- 移动的背景 -->
          <div
            class="absolute transition-all duration-300 ease-in-out bg-light-primary-6 dark:bg-dark-primary-6 text-light-text-4 dark:text-dark-text-4 rounded-md"
            :style="{ left: indicatorPosition, width: '25%' }"
          ></div>
          <div
            v-for="(type, index) in chatTypes"
            :key="type.value"
            @click="setChatListType(type.value)"
            class="w-1/4 cursor-pointer flex justify-center items-center rounded-md relative z-10"
            :class="
              currentChatListType === type.value
                ? 'text-common-white'
                : 'text-light-text-4 dark:text-dark-text-4'
            "
          >
            {{ type.label }}
          </div>
        </div>
      </div>
      <!-- 会话列表 -->
      <ChatSeesionList v-if="currentChatListType === 'all'" class="overflow-auto w-full flex-1" />
      <!-- 好友列表 -->
      <ChatContactPerson
        v-if="currentChatListType === 'personal'"
        class="overflow-auto w-full flex-1"
      />
      <!-- 群组列表 -->
      <!-- <GroupList class="overflow-auto w-full flex-1" /> -->
    </div>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '@renderer/stores/user/user'
import ChatContactPerson from '@renderer/views/Chat/components/ChatContactPerson.vue'
import ChatSeesionList from '@renderer/views/Chat/components/ChatSeesionList.vue'
import { computed, ref } from 'vue'
import AddFriend from './AddFriend.vue'

const store = UserStore()
const currentChatListType = ref('all')
// 搜索框绑定文本
const searchText = ref('')

// 定义聊天类型数组
const chatTypes = [
  { value: 'all', label: '全部' },
  { value: 'personal', label: '好友' },
  { value: 'group', label: '群组' },
  { value: 'notification', label: '通知' }
]

// 计算指示器的位置
const indicatorPosition = computed(() => {
  switch (currentChatListType.value) {
    case 'all':
      return '0%'
    case 'personal':
      return '25%'
    case 'group':
      return '50%'
    case 'notification':
      return '75%'
    default:
      return '0%'
  }
})

// 更新当前聊天列表类型
const setChatListType = (type: string) => {
  currentChatListType.value = type
}
</script>

<style scoped>
:deep(.css-dev-only-do-not-override-1p3hq3p).ant-input-affix-wrapper > input.ant-input {
  @apply bg-light-bg-2 dark:bg-dark-bg-2 text-light-text-1 dark:text-dark-text-1;
}

:deep(.ant-input::placeholder) {
  @apply text-light-text-4 dark:text-dark-text-4;
}

/* 调整列表的高度和样式 */
.overflow-auto {
  height: calc(100% - 55px);
}
.web {
  .ant-modal-content {
    background-color: transparent;
    box-shadow: none;
  }
}
.absolute {
  width: 33.33%;
  height: 100%;
  top: 0;
}
</style>
