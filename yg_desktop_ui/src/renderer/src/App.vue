<template>
  <div>
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '@renderer/stores/user/user'
import { watch } from 'vue'
const userStore = UserStore()

// 定义 WindowApi 接口
interface WindowApi {
  send: (channel: string | null, ...args: any[]) => void
  on: (channel: string, listener: (event: any, ...args: any[]) => void) => void
}

// 扩展全局 Window 接口
declare global {
  interface Window {
    api: WindowApi
  }
}

watch(
  () => userStore.token,
  (newToken) => {
    window.api.send('isLogin', newToken)
  },
  {
    immediate: true
  }
)
</script>

<style scoped>
/* 你的样式 */
</style>
