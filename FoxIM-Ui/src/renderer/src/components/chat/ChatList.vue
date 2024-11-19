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
        class="h-[30px] w-[180px] rounded-2xl"
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
    <div class="overflow-auto text-black w-full h-[calc(100%-55px)]">
      <ChatSeesionList v-if="store.selectedSidebar == 0" />
      <ChatContactPerson v-if="store.selectedSidebar == 1" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '@renderer/stores/user'
import { ref } from 'vue'
import AddFriend from './AddFriend.vue'
import ChatContactPerson from './ChatContactPerson.vue'
import ChatSeesionList from './ChatSeesionList.vue'
const store = UserStore()

// 搜索框绑定的文本
const searchText = ref('')
</script>

<style scoped>
/* 调整列表的高度和样式 */
.overflow-auto {
  height: calc(100% - 55px);
}
</style>
