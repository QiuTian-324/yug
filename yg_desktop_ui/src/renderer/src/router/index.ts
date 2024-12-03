import { newWarningMessage } from '@renderer/pkg/messages';
import { UserStore } from '@renderer/stores/user';
import { createRouter, createWebHistory } from 'vue-router';
import ChatRoom from '../views/ChatRoom.vue';
import Login from '../views/Login.vue';
import Register from '../views/Register.vue';



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
  // {
  //   path: '/callback',
  //   name: 'callback',
  //   component: Callback,
  // },
];

const router = createRouter({
  history: createWebHistory(),
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

