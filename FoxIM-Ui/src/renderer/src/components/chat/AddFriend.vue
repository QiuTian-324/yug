<template>
  <div class="w-full max-w-md mx-auto">
    <a-form :model="store.addFriendData">
      <a-form-item field="friendName" tooltip="Please enter friend's username" label="用户名">
        <a-input
          v-model="store.addFriendData.username"
          placeholder="输入好友用户名"
          class="w-full"
        />
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '@renderer/stores/user'
const store = UserStore()
// const form = reactive({
//   friendName: ''

// })
</script>
