import baseRoutes from './base'

import type { RouteRecordRaw } from 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    requiresAuth?: boolean
  }
}

export default [
  {
    path: '/',
    name: 'Layout',
    component: async () => import('@/views/Layout/Index.vue'),
    meta: {
      authRequired: true,
    },
    children: baseRoutes,
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
