import { authApi, logger } from '@/services'
import { defineStore } from 'pinia'
import router from '@/router'
import type { SessionStoreState } from './types'
import type { CredentialsResponse } from '@/types'

/**
 * The session store manages app-wide user session state.
 */
export const useSessionStore = defineStore('session', {
  state: (): SessionStoreState => ({
    isAuthenticated: false,
    exp: null,
    renewalTask: null,
    username: null,
  }),

  actions: {
    setUserState({ username, exp }: CredentialsResponse) {
      this.username = username
      this.exp = exp

      this.isAuthenticated = true
    },

    async verifySession() {
      try {
        const { ok, data } = await authApi.renew()

        if (!ok) {
          logger.error(`renew failed `, { ok, data })
          throw Error('@todo')
        }

        this.setUserState(data)
        this.autoRenew()
      } catch (ex) {
        await this.logout()
      }
    },

    async logout() {
      if (this.renewalTask != null) {
        clearTimeout(this.renewalTask)
      }

      await authApi.logout()
      this.$reset()
      window.sessionStorage.clear()

      if (router.currentRoute.value.name !== 'Login') {
        router.push({ name: 'Login' })
      }
    },

    autoRenew() {
      // For whatever reason, the token does not expire.
      if (!this.exp || this.exp <= 0) {
        logger.error(`exp not set: ${this.exp}`)
        return
      }

      const timeUntilRenewal = this.exp * 0.75
      logger.info(
        `renewal task scheduled for ${
          timeUntilRenewal / 60.0
        } minutes from now`,
      )

      // set timeout to renewal time in ms
      this.renewalTask = setTimeout(this.verifySession, timeUntilRenewal * 1000)
    },
  },
})
