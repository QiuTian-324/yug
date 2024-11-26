<template>
  <div class="w-full max-w-md mx-auto">
    <a-form :model="store.addFriendData">
      <a-form-item
        field="friendName"
        :tooltip="$t('user.add_friend_tooltip')"
        :label="$t('user.username')"
      >
        <a-input v-model="store.addFriendData.username" class="w-full" />
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '../../stores/user'
const store = UserStore()
// const form = reactive({
//   friendName: ''

// })
</script>
