import { logger } from '@/services/logger'
import { useSessionStore } from '@/state'

import { predicate } from './utils'

import type { Router } from 'vue-router'

let attemptedEntryPoint: string | null = null

/**
 * Wrapper for all system navigation guards
 */
export function guards(this: Router) {
  this.beforeEach((to, from, next) => {
    const routeHas = predicate(to)

    const { isAuthenticated } = useSessionStore()

    if (isAuthenticated) {
      if (to.path === '/login') {
        next({ name: 'Dashboard' })
        return
      }
    }

    if (!routeHas('authRequired')) {
      next()
      return
    }

    logger.info('matched an `authRequired` route')

    if (!isAuthenticated) {
      attemptedEntryPoint = to.path
      next({ name: 'Login' })
      return
    }

    const redirect = attemptedEntryPoint
    attemptedEntryPoint = null

    redirect ? next({ path: redirect }) : next()
  })

  return this
}
