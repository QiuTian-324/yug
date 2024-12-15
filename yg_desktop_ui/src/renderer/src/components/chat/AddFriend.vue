<template>
  <div class="w-full max-w-md mx-auto bg-light-base text-light-text p-4 rounded-lg">
    <a-form :model="store.addFriendData">
      <a-form-item
        field="username"
        :tooltip="$t('user.add_friend_tooltip')"
        :label="$t('user.username')"
      >
        <a-input
          v-model="store.addFriendData.username"
          class="w-full bg-light-base text-light-text border border-light-border rounded-md"
        />
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '@renderer/stores/user'
const store = UserStore()
</script>
