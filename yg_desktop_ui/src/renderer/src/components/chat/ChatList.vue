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

  <div class="h-full w-full bg-light-base border-r">
    <!-- 搜索框和添加按钮的容器 -->
    <div class="flex py-2 border-b justify-center items-center h-[55px] w-full">
      <!-- 搜索框 -->
      <a-input
        :placeholder="$t('user.search_conversation')"
        v-model:value="searchText"
        class="h-[30px] w-[150px] rounded-2xl bg-light-base text-light-text border border-light-border"
      >
      </a-input>
      <!-- 添加好友按钮 -->
      <UserAddOutlined
        class="cursor-pointer text-light-text ml-2"
        @click="store.addFriendModalVisible = true"
      />
    </div>
    <!-- 会话列表 -->
    <div class="overflow-auto text-light-text w-full h-[calc(100%-55px)]">
      <ChatSeesionList v-if="store.selectedSidebar == 0" />
      <ChatContactPerson v-if="store.selectedSidebar == 1" />
    </div>
  </div>
</template>

<script setup lang="ts">
import AddFriend from '@renderer/components/chat/AddFriend.vue'
import ChatContactPerson from '@renderer/components/chat/ChatContactPerson.vue'
import ChatSeesionList from '@renderer/components/chat/ChatSeesionList.vue'
import { UserStore } from '@renderer/stores/user'
import { ref } from 'vue'
const store = UserStore()

// 搜索框绑定文本
const searchText = ref('')
</script>

<style scoped>
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
</style>
