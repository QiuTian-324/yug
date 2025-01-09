<template>
  <!-- 左侧侧边菜单 -->
  <div
    class="w-full flex items-center justify-center h-full bg-light-bg-2 dark:bg-dark-bg-2 rounded-xl"
  >
    <div class="h-full w-full">
      <div class="flex flex-1 items-center justify-center h-[11%] w-full">
        <a-avatar :size="50" shape="square">
          <IconUser />
        </a-avatar>
      </div>
      <!-- 侧边菜单项 -->
      <ul class="flex h-[85%] flex-col items-center justify-center gap-4 text-xl">
        <li
          v-for="(item, index) in menuItems"
          :key="index"
          @click="handleClick(item.action)"
          :class="
            currentSidebarItem === item.tooltip
              ? 'bg-light-accent dark:bg-dark-accent text-light-text-1 dark:text-dark-text-1'
              : 'bg-light-bg-2 dark:bg-dark-bg-2 text-light-text-2 dark:text-dark-text-2'
          "
          class="p-3 w-[40px] h-[40px] flex items-center justify-center text-center rounded-full cursor-pointer
          {{ currentSidebarItem !== item.tooltip ? 'hover:bg-light-hover-bg dark:hover:bg-dark-hover-bg' : '' }}
          transform hover:-translate-y-1 transition duration-200 relative"
        >
          <a-tooltip placement="right" :title="$t(`common.${item.tooltip}`)">
            <IconFont :type="item.icon" />
          </a-tooltip>
          <!-- 设置菜单 -->
          <div
            v-if="isSettingsMenuOpen && item.tooltip === 'settings'"
            class="absolute left-full ml-2 text-sm bg-light-bg-3 dark:bg-dark-bg-3 text-light-text-2 dark:text-dark-text-2 shadow-lg p-4 rounded w-40"
            @mouseleave="isSettingsMenuOpen = false"
          >
            <p
              v-for="(setting, index) in settingsMenuItems"
              :key="index"
              class="flex items-center cursor-pointer hover:bg-light-hover-bg dark:hover:bg-dark-hover-bg p-2 rounded"
              @click="setting.action"
            >
              <IconFont :type="setting.icon" class="mr-2" />
              {{ setting.label }}
            </p>
          </div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import router from '@renderer/router'
import { ThemeStore } from '@renderer/stores/theme/index'
import { UserStore } from '@renderer/stores/user/user'
import { ref } from 'vue'

const store = UserStore()
const themeStore = ThemeStore()
const currentSidebarItem = ref('messages')

const isSettingsMenuOpen = ref(false)

// 封装菜单项数据
const menuItems = [
  { icon: 'icon-jiaoyou', action: () => router.push('/chat'), tooltip: 'messages' },
  {
    icon: 'icon-yingyongzhongxin',
    action: () => router.push('/leyu_island'),
    tooltip: 'leyu_Island'
  },
  { icon: 'icon-jiqiren', action: () => router.push('/ai'), tooltip: 'ai_assistant' },
  { icon: 'icon-Line_suixinji', action: () => router.push('/remember'), tooltip: 'remember' },
  { icon: 'icon-xingqiu', action: () => router.push('/planet'), tooltip: 'planet' },
  {
    icon: 'icon-shezhi',
    action: () => (isSettingsMenuOpen.value = !isSettingsMenuOpen.value),
    tooltip: 'settings'
  }
]

// 设置菜单项数据
const settingsMenuItems = [
  {
    icon: 'icon-gerenxinxi',
    action: () => {
      goToProfile()
    },
    label: '个人信息'
  },
  {
    icon: themeStore.isDarkMode ? 'icon-kaiqitaiyangguangshezhi' : 'icon-yueliang',
    action: () => {
      toggleTheme()
    },
    label: themeStore.isDarkMode ? '日间模式' : '夜间模式'
  },
  {
    icon: 'icon-tuichu',
    action: () => {
      handleLogout()
    },
    label: '退出登录'
  }
]

// 处理点击事件
const handleClick = (action: () => void) => {
  action()
}

// 导航到个人信息页面
const goToProfile = () => {
  router.push('/profile')
  isSettingsMenuOpen.value = false
}

// 切换主题
const toggleTheme = () => {
  themeStore.toggleDarkMode()
}

// 处理退出登录
const handleLogout = () => {
  store.handleLogout()
  isSettingsMenuOpen.value = false
}
</script>

<style scoped></style>
