/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        light: {
          primary: '#dbf3ff', // 浅蓝色 主要背景色
          secondary: '#cfe8ff', // 更浅的蓝色 次要背景色
          accent: '#0056b3', // 深蓝色 主要按钮颜色
          base: '#ffffff', // 白色 基础背景色
          text: '#000000', // 黑色 主要文本颜色
          text_secondary: '#333333', // 深灰色 次要文本颜色
          border: '#bcd4e6', // 浅蓝灰色 边框颜色
          border_secondary: '#e2e8f0', // 更浅的灰色 次要边框颜色
          hover: '#fafafa', // 淡灰色 悬停背景色
          active: '#00509e', // 深蓝色 激活背景色
          button_primary: '#007bff', // 按钮颜色
          icon_hover: '#0056b3', // 图标悬停颜色
          message_me: '#90caf9', // 中蓝色 当前用户消息背景
          message_other: '#e3f2fd', // 浅蓝色 其他用户消息背景
        }
      },
    },
  },
  variants: {},
  plugins: [],
  darkMode: 'class',
}
