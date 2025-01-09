import { defineStore } from "pinia"
import { ref } from "vue"

export const ThemeStore = defineStore('theme', () => {
  const isDarkMode = ref(false)
  // 切换黑暗模式
  const toggleDarkMode = () => {
    isDarkMode.value = !isDarkMode.value
    document.documentElement.classList.toggle('dark', isDarkMode.value)
  }
  return {
    isDarkMode,
    toggleDarkMode
  }
})