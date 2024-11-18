<template>
  <!-- 左侧侧边菜单 -->
  <div class="w-full flex items-center justify-center h-full">
    <div class="h-full w-full">
      <div class="flex items-end justify-center h-[11%] w-full">
        <a-avatar :size="50" shape="square">
          <IconUser />
        </a-avatar>
      </div>
      <!-- 侧边菜单项 -->
      <ul class="flex h-[89%] flex-col items-center justify-center gap-4 text-xl">
        <li
          class="p-3 w-[50px] bg-white text-center rounded-lg cursor-pointer hover:bg-gray-200 transform hover:-translate-y-2 transition duration-200"
        >
          <icon-message />
        </li>
        <li
          class="p-3 w-[50px] bg-white text-center rounded-lg cursor-pointer hover:bg-gray-200 transform hover:-translate-y-1 transition duration-200"
        >
          <icon-user />
        </li>
        <li
          class="p-3 w-[50px] bg-white text-center rounded-lg cursor-pointer hover:bg-gray-200 transform hover:-translate-y-1 transition duration-200"
        >
          <icon-robot />
        </li>
        <li
          class="p-3 w-[50px] bg-white text-center rounded-lg cursor-pointer hover:bg-gray-200 transform hover:-translate-y-1 transition duration-200"
        >
          <icon-idcard />
        </li>
        <li
          class="p-3 w-[50px] bg-white text-center rounded-lg cursor-pointer hover:bg-gray-200 transform hover:-translate-y-1 transition duration-200"
        >
          <icon-settings />
        </li>
        <li
          @click="store.LoginIn"
          class="p-3 w-[50px] bg-white text-center rounded-lg cursor-pointer hover:bg-gray-200 transform hover:-translate-y-1 transition duration-200"
        >
          <icon-export />
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { UserStore } from '@renderer/stores/user'

const store = UserStore()
</script>
