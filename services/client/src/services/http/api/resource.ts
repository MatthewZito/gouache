import type { MutableResource, Resource } from '@/types'
import type { UUID } from '@/types/scalar'
import { HttpClient } from '../client'

const client = new HttpClient({
  baseUrl: `${import.meta.env.VITE_RESOURCE_API}/api/resource`,
})

export const resourceApi = {
  async getResource(id: UUID) {
    return client.get<Resource>(`/${id}`)
  },

  async getResources() {
    return client.get<Resource[]>()
  },

  async createResource<R = UUID, D = MutableResource>(payload: D) {
    return client.post<R, D>(`/`, payload, true)
  },

  async updateResource<D = MutableResource>(resourceId: UUID, payload: D) {
    return client.patch<null, D>(`/${resourceId}`, payload, true)
  },
}
