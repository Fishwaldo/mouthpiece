import { h, resolveComponent } from 'vue'
import { createRouter, createWebHashHistory } from 'vue-router'

import DefaultLayout from '@/layouts/DefaultLayout'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: DefaultLayout,
    redirect: '/dashboard',
    children: [
      {
        path: '/dashboard',
        name: 'Dashboard',
        component: () =>
          import(/* webpackChunkName: "dashboard" */ '@/views/Dashboard.vue'),
        meta: {
          auth: true
        },
      },
      {
        path: '/messages',
        name: 'Messages',
        component: () => import('@/views/messages/Messages.vue'),
        meta: {
          auth: true
        },
      },
      {
        path: '/apps',
        name: 'Applications',
        component: () => import('@/views/apps/Applications.vue'),
        meta: {
          auth: true
        },
      },
      {
        path: '/users',
        name: 'Users',
        component: () => import('@/views/users/Users.vue'),
        meta: {
          auth: true
        },
      },
      {
        path: '/transports',
        name: 'Transports',
        component: () => import('@/views/transports/Transports.vue'),
        meta: {
          auth: true
        },
      },
      {
        path: '/notifications',
        name: 'Notifications',
        component: () => import('@/views/notifications/Notifications.vue'),
        meta: {
          auth: true
        },
      },
      {
        path: '/settings',
        name: 'Settings',
        component: () => import('@/views/settings/Settings.vue'),
        meta: {
          auth: true
        },
      },        
    ],
  },
  {
    path: '/pages',
    redirect: '/pages/404',
    name: 'Pages',
    component: {
      render() {
        return h(resolveComponent('router-view'))
      },
    },
    children: [
      {
        path: '404',
        name: 'Page404',
        component: () => import('@/views/pages/Page404'),
      },
      {
        path: '500',
        name: 'Page500',
        component: () => import('@/views/pages/Page500'),
      },
      {
        path: 'login',
        name: 'Login',
        component: () => import('@/views/pages/Login'),
      },
      {
        path: 'register',
        name: 'Register',
        component: () => import('@/views/pages/Register'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHashHistory(process.env.BASE_URL),
  routes,
  scrollBehavior() {
    // always scroll to top
    return { top: 0 }
  },
})

export default router
