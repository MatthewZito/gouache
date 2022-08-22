import { authApi } from '@/services'
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
      this.exp = Math.floor(Date.now() / 1000) + exp

      this.isAuthenticated = true
    },

    async verifySession() {
      try {
        const { ok, data } = await authApi.renew()

        if (!ok) {
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
        return
      }

      // conv `now` to seconds because exp claim is represented as seconds since UNIX epoch
      const now = Date.now() / 1000

      // diff between expiration in seconds since epoch and current time in seconds since epoch
      let timeUntilRenewal = this.exp - now
      // adjust renewal time to 15 min prior to expiration
      timeUntilRenewal -= 15 * 60

      // set timeout to renewal time in ms
      this.renewalTask = setTimeout(this.verifySession, timeUntilRenewal * 1000)
    },
  },
})
