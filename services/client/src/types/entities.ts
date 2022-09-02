import type { UUID, UNIXTimestamp } from './scalar'

export interface Resource {
  id: UUID
  title: string
  createdAt: UNIXTimestamp
  updatedAt: UNIXTimestamp
  tags: string[]
}

export type MutableResource = Pick<Resource, 'tags' | 'title'>

export interface Report {
  id: UUID
  name: string
  caller: string
  data: string
  ts: UNIXTimestamp
}

export type MutableReport = Pick<Report, 'name' | 'caller' | 'data'>

export interface UserCredentials {
  username: string
  password: string
}

export interface CredentialsResponse {
  username: string
  exp: number
}
