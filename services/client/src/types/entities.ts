import type { UUID, UNIXTimestamp } from './scalar'

export interface Resource {
  id: UUID
  title: string
  createdAt: UNIXTimestamp
  updatedAt: UNIXTimestamp
  tags: string[]
}

export interface MutableResource {
  title: string
  tags: string[]
}
