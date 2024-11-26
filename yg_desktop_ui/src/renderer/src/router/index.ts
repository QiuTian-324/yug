import { newWarningMessage } from '@renderer/pkg/messages';
import { UserStore } from '@renderer/stores/user';
import { createRouter, createWebHashHistory } from 'vue-router';
import ChatRoom from '../layout/ChatRoom.vue';
import Login from '../layout/Login.vue';
import Register from '../layout/Register.vue';



const routes = [
  {
    path: '/',
    name: 'ChatRoom',
    component: ChatRoom,
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});


router.beforeEach((to, from, next) => {
  const store = UserStore();
  const isAuthenticated = !!store.token;
  console.log('isAuthenticated', isAuthenticated);

  if (to.meta.requiresAuth && !isAuthenticated) {
    newWarningMessage('Please login first');
    next({ name: 'login' });
  } else {
    next();
  }
});

export default router;

