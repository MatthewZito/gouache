import { predicate } from './utils'

import type { Router } from 'vue-router'

import { logger } from '@/services/logger'
import { useSessionStore } from '@/state'

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
        return next({ name: 'Dashboard' })
      }
    }

    if (!routeHas('authRequired')) {
      return next()
    }

    logger.info('matched an `authRequired` route')

    if (!isAuthenticated) {
      attemptedEntryPoint = to.path
      return next({ name: 'Login' })
    }

    const redirect = attemptedEntryPoint
    attemptedEntryPoint = null

    return redirect ? next({ path: redirect }) : next()
  })

  return this
}
