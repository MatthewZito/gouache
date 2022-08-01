import authRoutes from './auth'
import baseRoutes from './base'

import type { RouteRecordRaw } from 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
  }
}

export default [
  {
    path: '',
    name: 'BaseLayout',
    component: async () => import('@/views/Layout.vue'),
    meta: {
      authRequired: true,
    },
    children: baseRoutes,
  },
  {
    path: '/auth',
    name: 'AuthLayout',
    component: async () => import('@/views/AuthLayout.vue'),
    children: authRoutes,
  },
  {
    path: '/:catchAll(.*)*',
    name: 'NotFound',
    component: async () => import('@/views/PageNotFound.vue'),
    meta: {
      authRequired: true,
    },
  },
] as RouteRecordRaw[]
