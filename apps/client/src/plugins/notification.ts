import type { Plugin } from 'vue'

import { Notify } from 'quasar'

type V = Parameters<Notify['registerType']>[1]
type NotificationTypes = 'error' | 'success' | 'warning'

const notificationTypes: Record<NotificationTypes, V> = {
  error: {
    icon: 'mdi-alert',
    color: 'red',
    timeout: 4000,
  },
  success: {
    icon: 'mdi-checkbox-marked-circle',
    color: 'green',
    timeout: 4000,
  },
  warning: {
    icon: 'mdi-alert',
    color: 'orange',
    timeout: 4000,
  },
}

export const showNotification = (type: NotificationTypes, message: string) => {
  Notify.create({
    type,
    message,
  })
}

export const registerNotifyPlugin: Plugin = {
  install() {
    Object.entries(notificationTypes).forEach(([key, value]) => {
      Notify.registerType(key, value)
    })
  },
}
