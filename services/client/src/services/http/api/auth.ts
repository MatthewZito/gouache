import type { CredentialsResponse, UserCredentials } from '@/types'

import { HttpClient } from '../client'

const client = new HttpClient({
  baseUrl: `${import.meta.env.VITE_AUTH_API}/auth`,
  withCredentials: true,
})

export const authApi = {
  async login({ username, password }: UserCredentials) {
    return client.post<CredentialsResponse, UserCredentials>(
      '/login',
      {
        username,
        password,
      },
      true,
    )
  },

  async logout() {
    return client.post<null, undefined>('/logout')
  },

  async register({ username, password }: UserCredentials) {
    return client.post<CredentialsResponse, UserCredentials>('/register', {
      username,
      password,
    })
  },

  async renew() {
    return client.post<CredentialsResponse, undefined>('/renew')
  },
}
