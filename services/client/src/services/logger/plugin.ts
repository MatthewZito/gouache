import type { App, Plugin } from 'vue'

import { logger } from '.'

import type { Logger } from './client'

function debug(this: App, logger: Logger) {
  /* Runtime Exceptions */
  window.onerror = (message, source, line, column) => {
    logger.error(
      `Uncaught Exception: ${JSON.stringify(
        message,
      )}\nInfo: ${source} - Ln${line} Col${column}`,
    )
  }

  /* Unhandled Promise Rejections */
  window.onunhandledrejection = e => {
    logger.error(`Unhandled Promise Rejection: ${e.reason}`)

    e.preventDefault()
  }

  /* Vue Errors */
  this.config.errorHandler = (err, vm, info) => {
    logger.error(`Exception: ${err}\nInfo: ${info}`)
  }

  /* Vue Render Warnings */
  this.config.warnHandler = (msg, vm, trace) => {
    logger.warn(`Warning: ${msg}\nTrace: ${trace}`)
  }
}

/* Register globally accessible logger service */
export const debugPlugin: Plugin = {
  install(app) {
    debug.call(app, logger)

    app.config.globalProperties.$logger = logger
    app.provide('$logger', logger)
  },
}
