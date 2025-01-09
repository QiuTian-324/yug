import { UserStore } from '@renderer/stores/user/user'
import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'

const routes = [
  {
    path: '/',
    name: 'layout',
    component: () => import('@renderer/views/Layout.vue'),
    redirect: '/chat',
    children: [
      {
        path: '/chat',
        name: 'chat',
        component: () => import('@renderer/views/Chat/index.vue'),
        meta: {
          requiresAuth: true
        },
        children: [
          {
            path: '/chat/:id',
            name: 'chatDetail',
            component: () => import('@renderer/views/Chat/Chat.vue')
          }
        ]
      },
      {
        path: '/ai',
        name: 'ai',
        component: () => import('@renderer/views/AI/index.vue'),
        meta: {
          requiresAuth: true
        }
      },
      {
        path: '/remember',
        name: 'remember',
        component: () => import('@renderer/views/Remember/index.vue')
      },
      {
        path: '/planet',
        name: 'planet',
        component: () => import('@renderer/views/Planet/index.vue')
      },
      {
        path: '/leyu_island',
        name: 'leyu_island',
        component: () => import('@renderer/views/Leyu_Island/index.vue')
      }
    ]
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: {
      requiresAuth: false
    }
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
    meta: {
      requiresAuth: false
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const store = UserStore();
  const isAuthenticated = !!store.token;
  console.log('isAuthenticated', isAuthenticated);

  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'login' });
  } else {
    next();
  }
});

export default router
