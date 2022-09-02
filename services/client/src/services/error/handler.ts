import { showNotification } from '@/plugins/notification'
import { logger } from '@/services/logger'

import { GouacheError } from '.'
import { reportingApi } from '../http'

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
  // eslint-disable-next-line @typescript-eslint/no-unsafe-assignment
  logger.error({ ex })

  reportingApi.createReport({
    name: 'TEST',
    data: JSON.stringify(ex),
  })

  let ret =
    fallback || 'Uh oh, something went wrong. Please try refreshing the page.'

  // captures any instance of `GouacheError`, including all subclasses
  if (ex instanceof GouacheError) {
    const { friendly } = ex.serialize()

    ret = friendly
  }

  if (notify) {
    showNotification('error', ret)
  }

  return ret
}

function serializeError(ex: any) {
  if (typeof ex === 'string') {
    return ex
  }

  try {
    return JSON.stringify(ex)
  } catch (ex2) {
    return 'n/a'
  }
}
