import { BaseError } from '.'

import { showNotification } from '@/plugins/notification'
import { logger } from '@/services/logger'

interface ErrorHandlerOptions {
  notify: boolean
  fallback?: string
}

/**
 * Error handling middleware. Each error class in /errors uses an `internal` message (for dev) and `friendly` message (for user).
 * This handler will filter any captured errors and conditionally dispatch a notification to the user.
 *
 * @param ex Captured thrown object.
 * @param options.notify Display the `friendly` message to the user?
 */
export const useErrorHandler = (
  ex: any,
  { notify = false, fallback }: ErrorHandlerOptions,
) => {
  logger.error({ ex })

  let ret =
    fallback || 'Uh oh, something went wrong. Please try refreshing the page.'

  // captures any instance of `BaseError`, including all subclasses
  if (ex instanceof BaseError) {
    const { friendly } = ex.serialize()

    ret = friendly
  }

  if (notify) {
    showNotification('error', ret)
  }

  return ret
}
