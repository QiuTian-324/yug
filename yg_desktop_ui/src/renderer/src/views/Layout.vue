<template>
  <div class="flex w-screen h-screen">
    <!-- left menu asidebar -->
    <div
      class="w-24 p-4 pr-2 h-full bg-light-bg-1 dark:bg-dark-bg-1 text-light-text-1 dark:text-dark-text-1 flex"
    >
      <Sidebar />
    </div>
    <div class="flex-grow flex-1 flex bg-light-bg-1 dark:bg-dark-bg-1">
      <router-view />
    </div>
  </div>
</template>

<script setup lang="ts">
import Sidebar from '@renderer/components/Sidebar.vue'
</script>
