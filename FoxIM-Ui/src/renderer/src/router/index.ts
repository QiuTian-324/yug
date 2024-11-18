import { createRouter, createWebHashHistory } from 'vue-router';
import ChatRoom from '../layout/ChatRoom.vue'; // 假设这是聊天页面

const routes = [
  {
    path: '/',
    name: 'ChatRoom',
    component: ChatRoom,  // 默认路由加载聊天页面
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
