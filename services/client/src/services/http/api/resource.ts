import type { MutableResource, Resource } from '@/types'
import type { UUID } from '@/types/scalar'
import { HttpClient } from '../client'

const client = new HttpClient('http://localhost:5000/resource')

export const resourceApi = {
  async getResource(id: UUID) {
    return client.get<Resource>(`/${id}`)
  },

  async getResources() {
    return client.get<Resource[]>()
  },

  async createResource<R = UUID, D = MutableResource>(payload: D) {
    return client.post<R, D>(`/`, payload)
  },

  async updateResource<D = MutableResource>(resourceId: UUID, payload: D) {
    return client.patch<null, D>(`/${resourceId}`, payload)
  },
}
