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
          v-for="(item, index) in menuItems"
          :key="index"
          @click="handleClick(item.action)"
          class="p-3 w-[40px] h-[40px] flex items-center justify-center bg-white text-center rounded-lg cursor-pointer hover:bg-gray-200 transform hover:-translate-y-1 transition duration-200"
        >
          <a-tooltip :title="$t(`common.${item.tooltip}`)">
            <component :is="item.icon" />
          </a-tooltip>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  IdcardOutlined,
  LogoutOutlined,
  MessageOutlined,
  RobotOutlined,
  SettingOutlined,
  UserOutlined
} from '@ant-design/icons-vue'
import { UserStore } from '../../stores/user'

const store = UserStore()

// 封装菜单项数据
const menuItems = [
  { icon: MessageOutlined, action: () => (store.selectedSidebar = 0), tooltip: 'messages' },
  { icon: UserOutlined, action: () => (store.selectedSidebar = 1), tooltip: 'contacts' },
  { icon: RobotOutlined, action: () => (store.selectedSidebar = 2), tooltip: 'ai_assistant' },
  { icon: IdcardOutlined, action: () => (store.selectedSidebar = 3), tooltip: 'profile' },
  { icon: SettingOutlined, action: () => (store.selectedSidebar = 4), tooltip: 'settings' },
  { icon: LogoutOutlined, action: () => store.Logout(), tooltip: 'logout' }
]

// 处理点击事件
const handleClick = (action: () => void) => {
  action()
}
</script>

<style scoped>
/* 你可以在这里添加一些样式 */
</style>
