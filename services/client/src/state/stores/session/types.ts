export interface SessionStoreState {
  isAuthenticated: boolean
  username: string | null
  exp: number | null
  renewalTask: NodeJS.Timeout | null
}
