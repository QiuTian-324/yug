<template>
  <div
    class="h-full w-full p-4 bg-light-base-1 dark:bg-dark-base-1 text-light-text dark:text-dark-text"
  >
    <ul class="friend-list">
      <li
        v-for="friend in userStore.friendList"
        :key="friend.user_id"
        @click="onSelect(friend.user_id)"
        class="flex items-center cursor-pointer bg-light-fill-2 dark:bg-dark-fill-2 hover:bg-light-hover-bg dark:hover:bg-dark-hover-bg p-3 rounded-md"
      >
        <div class="flex items-center">
          <img :src="friend.avatar_url" alt="avatar" class="w-10 h-10 rounded-full" />
          <span class="ml-2 font-bold">{{ friend.nickname }}</span>
        </div>
      </li>
    </ul>
  </div>
</template>

<script lang="ts" setup>
import { UserStore } from '@renderer/stores/user/user'

const userStore = UserStore()

function onSelect(friendId: number) {
  userStore.selectedSeesionId = friendId
}
</script>
