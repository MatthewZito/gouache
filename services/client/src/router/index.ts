import { createWebHistory, createRouter } from 'vue-router'

import routes from '@/router/routes'

import { guards } from './guards'

/* Root Router Configurations */
export default guards.call(
  createRouter({
    history: createWebHistory(),
    routes,
  }),
)
