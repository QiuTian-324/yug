<template>
  <div class="w-full max-w-md mx-auto bg-light-base-1 text-light-text p-4 rounded-lg">
    <a-form :model="store.addFriendData">
      <a-form-item
        field="username"
        :tooltip="$t('user.add_friend_tooltip')"
        :label="$t('user.username')"
      >
        <a-input
          v-model="store.addFriendData.username"
          class="w-full text-light-text border border-light-border rounded-md"
        />
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '@renderer/stores/user/user'
const store = UserStore()
</script>

<style scoped>
/* :deep(.css-dev-only-do-not-override-1p3hq3p).ant-input-affix-wrapper > input.ant-input {
  @apply bg-light-base-1 dark:bg-dark-base-1 text-light-text-black dark:text-dark-text-white;
} */
</style>
