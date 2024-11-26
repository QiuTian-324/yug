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

  <div class="h-full w-full bg-white border-r">
    <!-- 搜索框和添加按钮的容器 -->
    <div class="flex py-2 border-b justify-center items-center h-[55px] w-full">
      <!-- 搜索框 -->
      <a-input
        v-model:value="searchText"
        :placeholder="$t('user.search_conversation')"
        class="h-[30px] w-[180px] rounded-2xl"
      >
        <template #suffix>
          <UserAddOutlined class="cursor-pointer" @click="store.addFriendModalVisible = true" />
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
import { ref } from 'vue'
import { UserStore } from '../../stores/user'
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
