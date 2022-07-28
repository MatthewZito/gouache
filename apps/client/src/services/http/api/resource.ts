import type { Resource } from '@/types'
import { HttpClient } from '../client'

const client = new HttpClient('http://localhost:5000/resource')

export const resourceApi = {
  async getResources() {
    return client.get<Resource[]>()
  },

  async createResource(payload: Omit<Resource, 'id'>) {
    return client.post(`/`, payload)
  },

  async updateResource<D = Omit<Resource, 'id' | 'key'>>(
    resourceKey: string,
    payload: D,
  ) {
    return client.patch<null, D>(`/${resourceKey}`, payload)
  },
}
